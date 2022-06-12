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

func TestManagerConstructor(t *testing.T) {

	wM := NewManager()

	assert.IsType(t, WatcherManager{}, wM)
	assert.Equal(t, []*Watcher(nil), wM.watchers)
}

// Test that Watchers can be added to the manager through the .AddWatchers function
func TestManagerAddWatcher(t *testing.T) {
	wM := NewManager()

	mockUrl, err := url.Parse("http://mock.url")
	if err != nil {
		log.Println(err)
	}
	urls := []url.URL{*mockUrl, *mockUrl}

	wM.AddWatchers(urls)

	assert.Equal(t, 2, len(wM.watchers))
	assert.Equal(t, wM.watchers[0].poolingTime, 1*time.Second)
}

// Test that Starting the manager will launch all the Watchers
func TestManagerStart(t *testing.T) {

	// launch mock web server
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("last-modified", "mock_modified_date")
	}))
	defer s.Close()

	wM := NewManager()

	mockUrl, err := url.Parse(s.URL)
	if err != nil {
		log.Println(err)
	}

	// Several URLS
	urls := []url.URL{*mockUrl, *mockUrl}

	wM.AddWatchers(urls)

	wM.Start()

	time.Sleep(50 * time.Millisecond)

	// validate watcher on watchers[0] started
	assert.Equal(t, true, wM.watchers[0].watching)
	assert.Equal(t, true, wM.watchers[1].watching)

}

// Test that the manager can start just one watcher by ID
func TestManagerStartById(t *testing.T) {

	// launch mock web server
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("last-modified", "mock_modified_date")
	}))
	defer s.Close()

	wM := NewManager()

	mockUrl, err := url.Parse(s.URL)
	if err != nil {
		log.Println(err)
	}
	urls := []url.URL{*mockUrl}

	wM.AddWatchers(urls)

	wM.StartById([]string{wM.watchers[0].id.String()})

	time.Sleep(50 * time.Millisecond)

	// validate watcher on watchers[0] started
	assert.Equal(t, true, wM.watchers[0].watching)

}

// Test that the manager can stop a Watcher by ID
func TestManagerStopWatcher(t *testing.T) {

	// launch mock web server
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("last-modified", "mock_modified_date")
	}))
	defer s.Close()

	wM := NewManager()

	mockUrl, err := url.Parse(s.URL)
	if err != nil {
		log.Println(err)
	}
	urls := []url.URL{*mockUrl}

	wM.AddWatchers(urls)

	wM.StartById([]string{wM.watchers[0].id.String()})

	time.Sleep(50 * time.Millisecond)

	// validate watcher on watchers[0] started
	assert.Equal(t, true, wM.watchers[0].watching)

	wM.StopWatcher(wM.watchers[0].id.String())

	// validate watcher on watchers[0] started
	assert.Equal(t, true, !wM.watchers[0].watching)

}
