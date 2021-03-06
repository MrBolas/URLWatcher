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

const poolingTime = 1 * time.Second

// Create new Manager to handle watchers
func NewManager() WatcherManager {
	return WatcherManager{}
}

// Starts Watchers
func (m *WatcherManager) Start() {
	for _, watcher := range m.watchers {
		if !watcher.watching {
			go watcher.start()
		}
	}
}

func (m *WatcherManager) StartById(ids []string) {
	for _, watcher := range m.watchers {
		for _, id := range ids {
			if watcher.id == uuid.FromStringOrNil(id) && !watcher.watching {
				go watcher.start()
			}
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

func (m *WatcherManager) StopWatcher(id string) {

	for _, watcher := range m.watchers {
		if watcher.id == uuid.FromStringOrNil(id) {
			// use channel to stop it
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
