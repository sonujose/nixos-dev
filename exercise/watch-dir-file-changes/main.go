package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

func main() {

	// Directory to monitor
	dir := "./monitor"

	// Create the directory if it doesn't exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
	}

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatalf("Failed to watch directory: %v", err)
	}

	defer watcher.Close()

	fileEvents := make(chan fsnotify.Event)
	errs := make(chan error)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fileEvents <- event
			case err := <-watcher.Errors:
				errs <- err

			}
		}
	}()

	// Goroutine to process events
	go func() {
		for {
			select {
			case event := <-fileEvents:
				handleEvent(event)
			case err := <-errs:
				log.Printf("Error: %v\n", err)
			}
		}
	}()

	// Start watching the directory
	err = watcher.Add(dir)
	if err != nil {
		log.Fatalf("Failed to watch directory: %v", err)
	}

}

func handleEvent(event fsnotify.Event) {
	switch {
	case event.Op&fsnotify.Create == fsnotify.Create:
		fmt.Printf("File created: %s\n", event.Name)
	case event.Op&fsnotify.Write == fsnotify.Write:
		fmt.Printf("File modified: %s\n", event.Name)
	case event.Op&fsnotify.Remove == fsnotify.Remove:
		fmt.Printf("File deleted: %s\n", event.Name)
	case event.Op&fsnotify.Rename == fsnotify.Rename:
		fmt.Printf("File renamed: %s\n", event.Name)
	case event.Op&fsnotify.Chmod == fsnotify.Chmod:
		fmt.Printf("File permissions changed: %s\n", event.Name)
	}
}
