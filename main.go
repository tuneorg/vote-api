package main

import (
	"fmt"
	"log"
	"net/http"
	config "vote-api/src"
	vote "vote-api/src/vote"
)

func main() {
	conf, err := config.Init("config.yaml")
	if err != nil {
		panic(fmt.Sprintf("couldn't initialize config: %v", err))
	}

	http.HandleFunc("/topgg", vote.VoteHandler(conf))

	addr := fmt.Sprintf("%s:%d", conf.ADDRESS, conf.PORT)

	fmt.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
