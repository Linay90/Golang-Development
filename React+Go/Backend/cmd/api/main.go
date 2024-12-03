package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"enviornment"`

	Version string `json:"veion"`
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "server port to listen")
	flag.StringVar(&cfg.env, "env", "development", "application environment")
	flag.Parse()

	fmt.Println("running")
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		currentStaus := AppStatus{
			Status:      "Available",
			Environment: cfg.env,
			Version:     version,
		}
		js, err := json.MarshalIndent(currentStaus, "", "\t")
		if err != nil {
			log.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)
	if err != nil {
		log.Println(err)
	}

}
