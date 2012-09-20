package dungeon

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"appengine"
	"appengine/memcache"
	"appengine/datastore"
)

var world *World

func serveError(c appengine.Context, w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "Internal Server Error")
	c.Errorf("%v", err)
}

func handle(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	tutZone := NewZone("Tutorial")
	saveNewZone(c, w, tutZone)

	item, err := memcache.Get(c, r.URL.Path)
	if err != nil && err != memcache.ErrCacheMiss {
		serveError(c, w, err)
		return
	}
	n := 0
	if err == nil {
		n, err = strconv.Atoi(string(item.Value))
		if err != nil {
			serveError(c, w, err)
			return
		}
	}
	n++
	item = &memcache.Item{
		Key:   r.URL.Path,
		Value: []byte(strconv.Itoa(n)),
	}
	err = memcache.Set(c, item)
	if err != nil {
		serveError(c, w, err)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "%q has been visited %d times", r.URL.Path, n)
}

func init() {
	http.HandleFunc("/", handle)

	setupWorld()
}

type exportZone struct {
	Name string
}

func saveNewZone(c appengine.Context, w http.ResponseWriter, zone *Zone) *datastore.Key {
	exportZone := &exportZone{zone.name}
    key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "zone", nil), exportZone)
    if err != nil {
    	serveError(c, w, err)
        return nil
    }

    return key
}

func setupWorld() {
	world = NewWorld("Mordor")
}
