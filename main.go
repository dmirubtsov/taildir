package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/papertrail/go-tail/follower"
	"github.com/radovskyb/watcher"
)

var followers = make(map[string]*follower.Follower)

func main() {
	var wg sync.WaitGroup
	if len(os.Args) < 2 {
		log.Fatalln("Please specify directories as arguments")
	}
	for i := 1; i < len(os.Args); i++ {
		log.Println("Watching directory " + os.Args[i])
		wg.Add(1)
		go watchDir(os.Args[i])
	}
	wg.Wait()
}

func watchDir(dirPath string) *watcher.Watcher {
	w := watcher.New()

	go func() {
		for {
			select {
			case event := <-w.Event:
				switch event.Op {
				case watcher.Create:
					go createFollower(event.Path)
				case watcher.Write:
					go createFollower(event.Path)
				case watcher.Remove:
					go closeFollower(event.Path)
				}
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.AddRecursive(dirPath); err != nil {
		log.Fatalln(err)
	}

	for path := range w.WatchedFiles() {
		go createFollower(path)
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}

	return w
}

func createFollower(filePath string) {
	if followers[filePath] != nil {
		return
	}

	f, err := follower.New(filePath, follower.Config{
		Whence: io.SeekEnd,
		Offset: 0,
		Reopen: false,
	})

	if f.Err() != nil {
		log.Fatal(err)
	}

	followers[filePath] = f

	for line := range f.Lines() {
		fmt.Println(line.String())
	}

	if f.Err() != nil {
		return
	}
}

func closeFollower(filePath string) {
	for k := range followers {
		if k == filePath {
			delete(followers, k)
			return
		}
	}
}
