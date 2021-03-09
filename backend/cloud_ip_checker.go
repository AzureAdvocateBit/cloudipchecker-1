package main

import (
	"log"
	"net/http"
	"os"

	"github.com/deanobalino/cloud_ip_checker/apiservicetags"
	"github.com/deanobalino/cloud_ip_checker/webservicetags"
)

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	log.Println("Listening on port: ", listenAddr)
	http.HandleFunc("/api/servicetags/manual", webservicetags.GetServiceTags)
	http.HandleFunc("/api/servicetags/api", apiservicetags.GetServiceTags)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
