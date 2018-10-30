package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/howeyc/fsnotify"
	"gopkg.in/gcfg.v1"
)

// AppConfig type
type AppConfig struct {
	StatHat struct {
		APIKey string
	}
	Flags struct {
		HasStatHat bool
	}
	GeoIP struct {
		Directory string
	}
}

// Config var
var Config = new(AppConfig)

func configWatcher(fileName string) {

	configReader(fileName)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := watcher.Watch(*flagconfig); err != nil {
		fmt.Println(err)
		return
	}

	for {
		select {
		case ev := <-watcher.Event:
			if ev.Name == fileName {
				if ev.IsCreate() || ev.IsModify() || ev.IsRename() {
					time.Sleep(200 * time.Millisecond)
					configReader(fileName)
				}
			}
		case err := <-watcher.Error:
			log.Println("fsnotify error:", err)
		}
	}

}

var lastReadConfig time.Time

func configReader(fileName string) error {

	stat, err := os.Stat(fileName)
	if err != nil {
		log.Printf("Failed to find config file: %s\n", err)
		return err
	}

	if !stat.ModTime().After(lastReadConfig) {
		return err
	}

	lastReadConfig = time.Now()

	log.Printf("Loading config: %s\n", fileName)

	cfg := new(AppConfig)

	err = gcfg.ReadFileInto(cfg, fileName)
	if err != nil {
		log.Printf("Failed to parse config data: %s\n", err)
		return err
	}

	cfg.Flags.HasStatHat = len(cfg.StatHat.APIKey) > 0

	// log.Println("STATHAT APIKEY:", cfg.StatHat.ApiKey)
	// log.Println("STATHAT FLAG  :", cfg.Flags.HasStatHat)

	Config = cfg

	return nil
}
