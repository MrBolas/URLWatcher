package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	cO "github.com/MrBolas/URLWatcher/console"
	"github.com/MrBolas/URLWatcher/models"
)

func main() {

	// read urls from file

	// Create Manager
	var wM models.WatcherManager

	// Build a cmd line prompt waiting for inputs
	// Available commands:
	// load: load file from path
	// add: url to add
	// list: lists watchers
	// kill: kills watcher
	// start: starts watchers
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("----------URL Watcher Shell-----------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		// Load File
		if strings.HasPrefix(text, "load") {
			args, err := cO.SanitizeInputs(text)
			if err != nil {
				fmt.Println(err)
				continue
			}

			urls, err := cO.ReadUrlsFromFile(args[0])
			if err != nil {
				fmt.Println(err)
				continue
			}

			wM.AddWatchers(urls)
			continue
		}

		// Add Url
		if strings.HasPrefix(text, "add") {
			args, err := cO.SanitizeInputs(text)
			if err != nil {
				fmt.Println(err)
				continue
			}

			urls := cO.ConvertToUrl(args)

			wM.AddWatchers(urls)
			continue
		}

		// list watchers
		if strings.HasPrefix(text, "ls") {
			wM.PrintStatus()
			continue
		}

		// Start watchers
		if strings.HasPrefix(text, "start") {
			wM.Start()
			continue
		}

		// kill watcher
		if strings.HasPrefix(text, "kill") {
			args, err := cO.SanitizeInputs(text)
			if err != nil {
				fmt.Println(err)
				continue
			}

			wM.KillWatcher(args[0])
			continue
		}

		fmt.Println("Unknown command")
	}

}
