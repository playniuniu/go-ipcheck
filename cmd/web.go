package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"ipcheck/ip2region"
	"log"
	"net/http"
)

var port int
var db string
var region *ip2region.Ip2Region

func index(w http.ResponseWriter, req *http.Request) {
	ipData := getIPDataAll(req)
	log.Printf("INFO: req / from: %s", ipData["ip"])
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	resp, _ := json.Marshal(ipData)
	_, err := w.Write(resp)
	if err != nil {
		panic(err)
	}
}

func search(w http.ResponseWriter, req *http.Request) {
	var respData map[string]string
	ipAddr, ok := req.URL.Query()["ip"]
	if !ok || len(ipAddr[0]) < 1 {
		respData = map[string]string{
			"error":  "missing ip params",
			"status": "failed",
		}
	} else {
		respData = getIPData(ipAddr[0])
	}
	log.Printf("INFO: req /search from: %s", getIP(req))
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	resp, _ := json.Marshal(respData)
	_, err := w.Write(resp)
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.IntVar(&port, "port", 8000, "server port")
	flag.StringVar(&db, "db", "./assets/ip2region.db", "ip database")
	flag.Parse()

	var err error
	region, err = ip2region.New(db)
	if err != nil {
		log.Fatal("FATAL: read ip2region database error")
	}
	defer region.Close()
	serverPort := fmt.Sprintf(":%d", port)
	http.HandleFunc("/", index)
	http.HandleFunc("/search", search)
	log.Printf("INFO: listen on %s\n", serverPort)
	log.Fatal(http.ListenAndServe(serverPort, nil))
}
