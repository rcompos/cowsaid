package main

// Cowsaid  -  Cowsay As A Service

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/rcompos/cowsaid/pkg/cowsaid"
)

const (
	// timeFormat = "2006-01-02 15:04:05"
	// // dateFormat is the date stamp
	// dateFormat = "Mon Jan _2 2006"
	// // logFileDateFormat for log files
	logFileDateFormat = "2006-01-02-150405"
	// cowsayBalloonWidth
	cowsayBalloonWidth = 80
	fortuneAltDir      = "/usr/share/fortunes-alt"
	altDefault         = "/alt"
)

var mu sync.Mutex
var count int

func main() {

	//var errFile string
	//flag.StringVar(&errFile, "e", "./src/error.txt", "Errors file")
	//flag.Parse()
	var dirAlt string
	flag.StringVar(&dirAlt, "f", fortuneAltDir, "Alternate fortune directory")
	flag.Parse()

	// Create log file
	logFileDateFormat := "2006-01-02-150405"
	logStamp := time.Now().Format(logFileDateFormat)
	//logfile := "./src/cowsayer-log-" + string(logStamp) + ".out"
	logfile := "/tmp/cowsayer-log-" + string(logStamp) + ".out"

	logf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer logf.Close()
	log.SetOutput(logf) //log.Println("Test log message")

	cowSaid := func(w http.ResponseWriter, r *http.Request) {
		msgs, ok := r.URL.Query()["msg"]
		var output string
		if !ok || len(msgs) == 0 { // generate fortune
			output = cowsaid.CowFortune("")
		} else { // use msg
			// Query()["key"] will return an array of items,
			// we only want the single item.
			msg := msgs[0]
			// log.Println("Url Param 'key' is: " + string(key))
			output = cowsaid.CowFortune(msg)
		}

		w.Write([]byte(output))
		w.Write([]byte("\n"))
	}

	cowAlt := func(w http.ResponseWriter, r *http.Request) {
		// URL Path /s
		// Example: http://localhost:80/?s=8ball
		//name := r.URL.Query().Get("s")
		//fmt.Printf("name: %v\n", name)
		urlPath := html.EscapeString(r.URL.Path)
		dirPath := strings.TrimPrefix(urlPath, "/s")
		if dirPath == "" {
			dirPath = altDefault
		}

		phrase := cowsaid.GetFortune(dirAlt + dirPath)
		phrase = strings.TrimSuffix(phrase, "\n")
		if phrase == "" {
			phrase = "404 Phrase Not Found"
		}
		output := cowsaid.CowFortune(phrase)
		w.Write([]byte(output))
		// w.Write([]byte("\n"))
	}

	//uploader := func(w http.ResponseWriter, r *http.Request) {
	//	//w.Write([]byte("UPLOAD TESt"))
	//	file, err := os.Create("./src/uploaded")
	//	if err != nil {
	//		panic(err)
	//	}
	//	n, err := io.Copy(file, r.Body)
	//	if err != nil {
	//		panic(err)
	//	}
	//	w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\n", n)))
	//}

	//viewErr := func(w http.ResponseWriter, r *http.Request) {
	//	v := viewErrLines(errLines)
	//	s := fmt.Sprintf("%v", v)
	//	w.Write([]byte(s))
	//}

	apiV1 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/api/v1\n"))
		w.Write([]byte("/api/v1/cowsay\n"))
		w.Write([]byte("/api/v1/info\n"))
		w.Write([]byte("/api/v1/ping\n"))
		w.Write([]byte("/api/v1/bad\n"))
		w.Write([]byte("/api/v1/err\n"))
		w.Write([]byte("/api/v1/count\n"))
		w.Write([]byte("/s/8ball\n"))
		//w.Write([]byte("/api/v1/raw/\n"))
		//w.Write([]byte("/api/v1/new/\n"))
	}

	http.HandleFunc("/", cowSaid)
	http.HandleFunc("/api", apiV1)
	http.HandleFunc("/api/", apiV1)
	http.HandleFunc("/api/v1", apiV1)
	http.HandleFunc("/api/v1/", apiV1)
	http.HandleFunc("/api/v1/cowsay", cowSaid)
	http.HandleFunc("/api/v1/cowsay/", cowSaid)
	http.HandleFunc("/api/v1/info", cowsaid.Info)
	http.HandleFunc("/api/v1/info/", cowsaid.Info)
	http.HandleFunc("/api/v1/ping", cowsaid.Ping)
	http.HandleFunc("/api/v1/ping/", cowsaid.Ping)
	http.HandleFunc("/api/v1/count", cowsaid.Counter)
	http.HandleFunc("/api/v1/count/", cowsaid.Counter)
	//http.HandleFunc("/api/v1/err", viewErr)
	//http.HandleFunc("/api/v1/err/", viewErr)
	//http.HandleFunc("/api/v1/upload", uploader)
	//http.HandleFunc("/api/v1/upload/", uploader)
	http.HandleFunc("/s", cowAlt)
	http.HandleFunc("/s/", cowAlt)
	//http.Handle("/src/", http.StripPrefix("/src/", fs))

	// File server for upload and download
	//maxUploadSize := 2 * 1024 // 2 MB
	////httpfs := http.FileServer(http.Dir(uploadPath))
	//uploadPath := "./src"
	//http.Handle("/files/", http.StripPrefix("/files", httpfs))
	//log.Print("API endpoints /api/v1/upload/ for uploading and /files/ for downloading.")

	//httpCowsayer := "localhost:8080"
	httpCowsayer := ":80"
	listenMsg := "Listening on " + httpCowsayer + " ..."
	fmt.Println(listenMsg)
	//log.Println(listenMsg)
	log.Fatal(http.ListenAndServe(httpCowsayer, nil))

} // end main
