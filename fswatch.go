// fswatch - Watch for changes in file system.
//
// @see https://fsnotify.org/
// @see https://medium.com/@skdomino/watch-this-file-watching-in-go-5b5a247cf71f
//
// Copyright 2019 William Yeh <william.pjyeh@gmail.com>. All Rights Reserved.
//
//

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"container/list"
	"github.com/docopt/docopt-go"
	"github.com/fsnotify/fsnotify"
)

var DIR_TO_WATCH = list.New()

var watcher *fsnotify.Watcher

const USAGE string = `fswatch - Watch for changes in file system.

Usage:
  fswatch <dir>...
  `

func main() {
	process_cmdline()

	show_dir_to_watch()
	create_watcher()
}

// This func parses and validates cmdline args
func process_cmdline() map[string]interface{} {

	args, _ := docopt.ParseDoc(USAGE)
	//fmt.Println(args)

	// collect dir(s) to watch
	for _, dir := range args["<dir>"].([]string) {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Printf("Error: no such file or directory: %s", dir)
			os.Exit(1)
		}
		DIR_TO_WATCH.PushBack(dir)
	}

	return args
}

func show_dir_to_watch() {
	fmt.Println("Dir(s) to watch...")
	for e := DIR_TO_WATCH.Front(); e != nil; e = e.Next() {
		fmt.Println(" - " + e.Value.(string))
	}

}

func create_watcher() {
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	for e := DIR_TO_WATCH.Front(); e != nil; e = e.Next() {
		if err := filepath.Walk(e.Value.(string), watchDir); err != nil {
			fmt.Println("ERROR", err)
		}
	}

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
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	<-done
}


// watchDir gets run as a walk func, searching for directories to add watchers to
func watchDir(path string, file_info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
		return err
	}

	if file_info.IsDir() {
		//fmt.Printf("--- %s\n", file_info.Name())
		return watcher.Add(path)
	}

	return nil
}