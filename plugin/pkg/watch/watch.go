package watch

// Chan is used to inform the server of a change
type Chan chan []string

// Watchable is the interface watchable plugins should implement
type Watchable interface {
	Name() string
	SetWatchChan(Chan)
	Watch(qname string) error
	StopWatching(qname string)
}
