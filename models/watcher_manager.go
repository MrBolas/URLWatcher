package models

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gofrs/uuid"
)

type WatcherManager struct {
	watchers []*Watcher
}

const poolingTime = 5 * time.Second

// Create new Manager to handle watchers
func NewManager() WatcherManager {
	return WatcherManager{}
}

// Starts Watchers
func (m *WatcherManager) Start() {
	for _, watcher := range m.watchers {
		if !watcher.watching {
			go watcher.start()
			watcher.watching = true
		}
	}
}

// Add watcher to manager and start it right away
func (m *WatcherManager) AddWatchers(urls []url.URL) {
	for _, url := range urls {
		newWatcher, err := NewWatcher(url, poolingTime)
		if err != nil {
			log.Println("Failed to create watcher for url:", url, "with error:", err)
		} else {
			m.watchers = append(m.watchers, &newWatcher)
		}
	}
}

func (m *WatcherManager) KillWatcher(id string) {

	for _, watcher := range m.watchers {
		if watcher.id == uuid.FromStringOrNil(id) {
			// use channel to kill it
			watcher.channel <- Message{keepAlive: false}
		}
	}
}

func (m *WatcherManager) PrintStatus() {
	fmt.Println("-------------------------------------------------- Watchers --------------------------------------------------")
	fmt.Println("ID					status		url			last modified")
	for _, w := range m.watchers {
		if w.watching {
			fmt.Println(w.id.String(), " Watching ", w.url.String(), w.lastModified)
		} else {
			fmt.Println(w.id.String(), " Not Watching ", w.url.String(), w.lastModified)
		}
	}
}
