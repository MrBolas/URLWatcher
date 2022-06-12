package models

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test constructor function
func TestWatcherConstructor(t *testing.T) {

	mockUrl := url.URL{Scheme: "http",
		Host: "www.google.com"}

	var poolingTime time.Duration = 5 * time.Second

	mockWatcher, err := NewWatcher(mockUrl, poolingTime)
	if err != nil {
		log.Println(err)
	}

	assert.NotEmpty(t, mockWatcher.id)
	assert.Equal(t, "http://www.google.com", mockWatcher.url.String())
	assert.Equal(t, 5*time.Second, mockWatcher.poolingTime)
	assert.Equal(t, false, mockWatcher.watching)
}

// Test Watcher start function
func TestWatcherStart(t *testing.T) {

	// launch mock web server
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("last-modified", "mock_modified_date")
	}))
	defer s.Close()

	// define watcher
	mockUrl, err := url.Parse(s.URL)
	if err != nil {
		log.Println(err)
	}

	var poolingTime time.Duration = 5 * time.Second

	mockWatcher, err := NewWatcher(*mockUrl, poolingTime)
	if err != nil {
		log.Println(err)
	}

	// launch go routine
	go mockWatcher.start()

	// Time to allow go routine to start
	time.Sleep(50 * time.Millisecond)

	// check the watcher is watching
	assert.Equal(t, true, mockWatcher.watching)
}

func TestWatcherStopsWithChannelMessage(t *testing.T) {

	// launch mock web server
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("last-modified", "mock_modified_date")
	}))
	defer s.Close()

	// define watcher
	mockUrl, err := url.Parse(s.URL)
	if err != nil {
		log.Println(err)
	}

	var poolingTime time.Duration = 5 * time.Second

	mockWatcher, err := NewWatcher(*mockUrl, poolingTime)
	if err != nil {
		log.Println(err)
	}

	// launch go routine
	go mockWatcher.start()
	time.Sleep(50 * time.Millisecond)

	// check the watcher is watching
	assert.Equal(t, true, mockWatcher.watching)

	// send stop message on channel
	mockWatcher.channel <- Message{keepAlive: false}

	assert.Equal(t, false, mockWatcher.watching)
}

// Test URL updated validation
func TestWasUrlUpdated(t *testing.T) {

	// launch mock web server
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("last-modified", "mock_modified_date")
	}))
	defer s.Close()

	// mock webserver url
	mockUrl, err := url.Parse(s.URL)
	if err != nil {
		log.Println(err)
	}

	var poolingTime time.Duration = 5 * time.Second

	mockWatcher, err := NewWatcher(*mockUrl, poolingTime)
	if err != nil {
		log.Println(err)
	}

	c := http.Client{}

	updated, err := mockWatcher.wasUrlUpdated(c)
	if err != nil {
		log.Println(err)
	}

	assert.NotEmpty(t, mockWatcher.id)
	assert.Equal(t, mockUrl.String(), mockWatcher.url.String())
	assert.Equal(t, 5*time.Second, mockWatcher.poolingTime)
	assert.Equal(t, false, mockWatcher.watching)
	assert.Equal(t, "mock_modified_date", mockWatcher.lastModified)

	assert.Equal(t, true, updated)
}
