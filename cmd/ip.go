package main

import (
	"net/http"
	"strings"
)

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return strings.Split(forwarded, ":")[0]
	}
	return strings.Split(r.RemoteAddr, ":")[0]
}

func getIPData(ipAddr string) map[string]string {
	ipData, err := region.MemorySearch(ipAddr)
	respData := map[string]string{}
	if err == nil {
		respData["country"] = ipData.Country
		respData["province"] = ipData.Province
		respData["city"] = ipData.City
		respData["isp"] = ipData.ISP
		respData["ip"] = ipAddr
		respData["status"] = "success"
	} else {
		respData["error"] = err.Error()
		respData["status"] = "failed"
	}
	return respData
}

func getIPDataAll(r *http.Request) map[string]string {
	ipAddr := getIP(r)
	respData := getIPData(ipAddr)
	return respData
}
