package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

func WatchFileChange(filename string) {
	// Specify the path to the file to monitor.
	// filename := "D:\\TrongTran\\Job\\Reactjs\\golang\\test\\data\\watch_file_change_test.txt"
	//
	// Get the initial modification time of the file.
	info, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	modTime := info.ModTime()
	fmt.Printf("Start watching the changes of file [%v]...\n", filename)
	for {
		// Check if the file has been modified.
		info, err = os.Stat(filename)
		if err != nil {
			log.Fatal(err)
		}
		if info.ModTime() != modTime {
			modTime = info.ModTime()
			modTimeStr := info.ModTime().Format("2006-01-02 15:04:05.99")
			fmt.Printf("[%v]File [%v] has been modified!\n", modTimeStr, filename)
		}
		// Wait for a short period before checking again.
		time.Sleep(1 * time.Second)
	}
}
