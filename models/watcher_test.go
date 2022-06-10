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

func TestWasUrlUpdated(t *testing.T) {

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("last-modified", "mock_modified_date")
	}))
	defer s.Close()

	mockUrl := url.URL{Scheme: "http",
		Host: "localhost"}

	var poolingTime time.Duration = 5 * time.Second

	mockWatcher, err := NewWatcher(mockUrl, poolingTime)
	if err != nil {
		log.Println(err)
	}

	c := s.Client()

	updated, err := mockWatcher.wasUrlUpdated(*c)
	if err != nil {
		log.Println(err)
	}

	assert.NotEmpty(t, mockWatcher.id)
	assert.Equal(t, "http://localhost", mockWatcher.url.String())
	assert.Equal(t, 5*time.Second, mockWatcher.poolingTime)
	assert.Equal(t, false, mockWatcher.watching)
	assert.Equal(t, "mock_modified_date", mockWatcher.lastModified)

	assert.Equal(t, true, updated)

}
