package dnsserver

import (
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/miekg/dns"

	"github.com/coredns/coredns/pb"
	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/pkg/watch"
)

// watcher contains all the data needed to manage watches
type watcher struct {
	changes watch.Chan
	counter int64
	watches map[string]watchlist
	plugins []watch.Watchable
	mutex   sync.Mutex
}

type watchlist map[int64]pb.DnsService_WatchServer

func newWatcher(zones map[string]*Config) *watcher {
	w := &watcher{changes: make(watch.Chan), watches: make(map[string]watchlist)}

	for _, config := range zones {
		plugins := config.Handlers()
		for _, p := range plugins {
			if x, ok := p.(watch.Watchable); ok {
				x.SetWatchChan(w.changes)
				w.plugins = append(w.plugins, x)
			}
		}
	}

	//TODO: maybe a stop channel, work properly with reloads?
	go w.processWatches()
	return w
}

func (w *watcher) nextID() int64 {
	w.mutex.Lock()

	w.counter++
	id := w.counter

	w.mutex.Unlock()
	return id
}

// watch is used to monitor the results of a given query.
func (w *watcher) watch(stream pb.DnsService_WatchServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		create := in.GetCreateRequest()
		if create != nil {
			msg := new(dns.Msg)
			err := msg.Unpack(create.Query.Msg)
			if err != nil {
				log.Printf("[WARNING] Could not decode watch request: %s\n", err)
				stream.Send(&pb.WatchResponse{Err: "could not decode request"})
				continue
			}
			id := w.nextID()
			if err := stream.Send(&pb.WatchResponse{WatchId: id, Created: true}); err != nil {
				// if we fail to notify client of watch creation, don't create the watch
				continue
			}

			qname := msg.Question[0].Name
			w.mutex.Lock()
			if _, ok := w.watches[qname]; !ok {
				w.watches[qname] = make(watchlist)
			}
			w.watches[qname][id] = stream
			w.mutex.Unlock()

			for _, p := range w.plugins {
				err := p.Watch(qname)
				if err != nil {
					log.Printf("[WARNING] Failed to start watch for %s in plugin %s: %s\n", qname, p.Name(), err)
					stream.Send(&pb.WatchResponse{Err: fmt.Sprintf("failed to start watch for %s in plugin %s", qname, p.Name())})
				}
			}
			continue
		}

		cancel := in.GetCancelRequest()
		if cancel != nil {
			w.mutex.Lock()
			for qname, wl := range w.watches {
				ws, ok := wl[cancel.WatchId]
				if !ok {
					continue
				}

				// only allow cancels from the client that started it
				// TODO: test what happens if a stream tries to cancel a watchID that it doesn't own
				if ws != stream {
					continue
				}

				delete(wl, cancel.WatchId)

				// if there are no more watches for this qname, we should tell the plugins
				if len(wl) == 0 {
					for _, p := range w.plugins {
						p.StopWatching(qname)
					}
					delete(w.watches, qname)
				}

				// let the client know we canceled the watch
				stream.Send(&pb.WatchResponse{WatchId: cancel.WatchId, Canceled: true})
			}
			w.mutex.Unlock()
			continue
		}
	}
}

func (w *watcher) processWatches() {
	for {
		select {
		case changed := <-w.changes:
			w.mutex.Lock()
			for qname, wl := range w.watches {
				if plugin.Zones(changed).Matches(qname) == "" {
					continue
				}
				for id, stream := range wl {
					wr := pb.WatchResponse{WatchId: id, Qname: qname}
					err := stream.Send(&wr)
					if err != nil {
						log.Printf("Error sending to watch %d: %s\n", id, err)
					}
				}
			}
			w.mutex.Unlock()
		}
	}
}
