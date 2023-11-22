package main

import (
	"ehut/cmd/web"
	"flag"
	"log"
)

func main() {
	webport := flag.String("wport", ":8999", "def port 8999")
	webs := flag.Bool("web", true, "start web")
	flag.Parse()

	// de modificat in viitor
	if *webs {
		web.Srv(*webport)
	} else {
		log.Println("sv not start")
	}

}
