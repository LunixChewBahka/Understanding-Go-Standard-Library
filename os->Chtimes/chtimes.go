package main

import (
	"log"
	"os"
	"time"
)

func main() {
	mtime := time.Date(2016, time.February, 1, 3, 4, 5, 0, time.UTC)
	atime := time.Date(2017, time.March, 2, 4, 5, 6, 0, time.UTC)
	if err := os.Chtimes("example", atime, mtime); err != nil {
		log.Fatal(err)
	}
}
