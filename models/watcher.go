package models

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gofrs/uuid"
)

type Message struct {
	keepAlive bool
}

type Watcher struct {
	id           uuid.UUID
	url          url.URL
	lastModified string
	poolingTime  time.Duration
	watching     bool
	channel      chan Message
}

func NewWatcher(url url.URL, poolingTime time.Duration) (Watcher, error) {
	newUuid, err := uuid.NewV4()
	if err != nil {
		return Watcher{}, err
	}

	return Watcher{
		id:           newUuid,
		url:          url,
		lastModified: "",
		poolingTime:  poolingTime,
		watching:     false,
		channel:      make(chan Message),
	}, nil
}

// function which will run has a go Routine
func (w *Watcher) start() {

	c := http.Client{}
	w.watching = true
	defer func() {
		if r := recover(); r != nil {
			w.watching = false
			fmt.Println("Watcher", w.id, "panicked:", r)
			return
		}
	}()

	// Set initial state of lastModified
	_, err := w.wasUrlUpdated(c)
	if err != nil {
		log.Panicln("Watcher failed with error:", err, " for url:", w.url.String())
	}

	for {
		updated, err := w.wasUrlUpdated(c)
		if err != nil {
			log.Panicln("Watcher failed with error:", err, " for url:", w.url.String())
		}

		// if last modified was updated log url
		if updated {
			log.Println(w.url.String(), "was updated")
		}

		// listens for kill signal on watcher channel
		select {
		case msg := <-w.channel:
			if !msg.keepAlive {
				w.watching = false
				log.Println("Watcher", w.id, "was terminated")
				return
			}
		default:
		}

		time.Sleep(w.poolingTime)
	}

}

// function assert if url asset was updated
func (w *Watcher) wasUrlUpdated(c http.Client) (bool, error) {

	// send Head request
	resp, err := c.Head(w.url.String())
	if err != nil {
		return false, err
	}

	// update last modified
	lastModifiedFromResp := resp.Header.Get("last-modified")

	if lastModifiedFromResp != w.lastModified {
		w.lastModified = lastModifiedFromResp
		return true, nil
	}

	return false, nil
}
