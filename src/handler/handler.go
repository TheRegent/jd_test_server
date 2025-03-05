package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func SetupRoutes() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/load/cpu", loadCPUHandler)
	http.HandleFunc("/load/memory", loadMemoryHandler)
	http.HandleFunc("/crash", crashHandler)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Health check request received")
	fmt.Fprint(w, "OK")
}

func loadCPUHandler(w http.ResponseWriter, r *http.Request) {
	var (
		duration  time.Duration
		intensity int
		magnifier = 100
	)

	rawDur := r.URL.Query().Get("duration")
	rawInt, err := strconv.Atoi(r.URL.Query().Get("intensity"))
	if err != nil {
		log.Println("Error parsing intensity:", err)
	}

	if rawDur == "" || strings.HasPrefix(rawDur, "-") {
		duration = time.Second * 10
	} else {
		duration, _ = time.ParseDuration(rawDur)
	}

	if rawInt == 0 {
		intensity = 40
	} else {
		intensity = int(rawInt)
	}

	log.Printf("Received load request with duration %s and intensity %d\n", duration, intensity)

	start := time.Now()
	for time.Since(start) < duration {
		fibonacci(intensity * magnifier)
	}

	fmt.Fprintf(w, "Finished load test with duration %s and intensity %d\n", duration, intensity)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func loadMemoryHandler(w http.ResponseWriter, r *http.Request) {
	var (
		duration  time.Duration
		memSize   int
		magnifier = 10
	)

	rawDur := r.URL.Query().Get("duration")
	if rawDur == "" || strings.HasPrefix(rawDur, "-") {
		duration = 10 * time.Second
	} else {
		duration, _ = time.ParseDuration(rawDur)
	}

	rawSize := r.URL.Query().Get("intensity")
	sizeInt, err := strconv.Atoi(rawSize)
	if err != nil {
		log.Println("Error parsing size:", err)
	}
	if sizeInt == 0 {
		memSize = 10
	} else {
		memSize = sizeInt
	}

	log.Printf("Received memory load request with duration %s and intensity %d\n", duration, memSize)

	start := time.Now()
	var allocated [][]byte
	for time.Since(start) < duration {
		block := make([]byte, memSize*1024*1024*magnifier)
		for i := range block {
			block[i] = 1
		}
		allocated = append(allocated, block)
	}

	fmt.Fprintf(w, "Finished memory load test with duration %s and size %d MB\n", duration, memSize)
}

func crashHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Crash handler triggered: Application will crash now")

	os.Exit(1)
}
