package utils

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

func WatchFolderChange(path string) {
	if path == "" {
		path = `D:\TrongTran\meta-projects\.POSPRO\source\meta-vi.com\pospro\golang\test\data`
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("error:", err)
		return
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
				fmt.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	<-done
}
