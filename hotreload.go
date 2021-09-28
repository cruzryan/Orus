package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/fatih/color"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

func watch() {
	color.Cyan("Starting Orus Watcher...")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					var wg sync.WaitGroup
					wg.Add(1)
					go func() {
						compile()
						defer wg.Done()
					}()
					wg.Wait()
					fmt.Println("--------- WAITING DONDE --------")
					restartVsim()
					run()
					examine("switch", "a")
					examine("switch", "b")
					examine("switch", "f")
					truthTable()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func watchDir(path string, fi os.FileInfo, err error) error {

	// since fsnotify can watch all the files in a directory, watchers only need
	// to be added to each nested directory
	if fi.Mode().IsDir() {
		return watcher.Add(path)
	}

	return nil
}

func update() {

}
