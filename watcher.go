package fsnotify

import (
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
)

type evtCb struct {
	fsnotify.Op
	cb func(fsnotify.Event)
}

func Watcher(path string, evt evtCb) (watcher *fsnotify.Watcher) {
	var err error
	// Create new watcher.
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	//defer watcher.Close()
	// Start listening for events.
	go func() {
		is1st := true
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(evt.Op) {
					if is1st { //prevent double event
						is1st = false
						evt.cb(event)
					}
					go func() {
						time.Sleep(time.Second)
						is1st = true
					}()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println(err)
			}
		}
	}()
	// Add a path.
	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	//<-make(chan struct{})
	return
}
