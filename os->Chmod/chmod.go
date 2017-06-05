package main

import (
	"log"
	"os"
)

func main() {
	/*
		accoding to unix spec 0644 is:
		+ (owning) User: read & write
		+ Group: read
		+ Other: read

		in your terminal this might read like this if you execute "ls -la"
		-rw-r--r--
	*/
	if err := os.Chmod("example", 0644); err != nil {
		log.Fatal(err)
	}
}
