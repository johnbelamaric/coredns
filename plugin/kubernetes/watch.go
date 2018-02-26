package kubernetes

import (
	"fmt"

	"github.com/coredns/coredns/plugin/pkg/watch"
)

// SetWatchChan implements watch.Watchable
func (k *Kubernetes) SetWatchChan(c watch.Chan) {
	k.watchChan = c
}

// Watch is called when a watch is started for a name.
func (k *Kubernetes) Watch(qname string) error {
	if k.watchChan == nil {
		return fmt.Errorf("cannot start watch because the channel has not been set")
	}
	k.watched[qname] = true
	return nil
}

// StopWatching is called when no more watches remain for a name
func (k *Kubernetes) StopWatching(qname string) {
	delete(k.watched, qname)
}
