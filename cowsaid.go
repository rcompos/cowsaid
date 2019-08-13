// Cowsayer  -  Cowsay As A Service

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	cowsay "github.com/Code-Hex/Neo-cowsay"
)

const (
	timeFormat = "2006-01-02 15:04:05"
	// dateFormat is the date stamp
	dateFormat = "Mon Jan _2 2006"
	// logFileDateFormat for log files
	logFileDateFormat = "2006-01-02-150405"
	// cowsayBalloonWidth
	cowsayBalloonWidth = 80
	fortuneAltDir     = "/usr/share/fortunes-alt"
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
		phrase := getFortune("")
		say, err := cowsay.Say(
			cowsay.Phrase(phrase),
			cowsay.Type("default"),
		)
		if err != nil {
			log.Println(err)
		}
		w.Write([]byte(say))
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
		phrase := getFortune(dirAlt + dirPath)
		if phrase == "" {
			phrase = "404 Phrase Not Found"
		}
		say, err := cowsay.Say(
			cowsay.Phrase(phrase),
			cowsay.Type("default"),
		)
		if err != nil {
			log.Println(err)
		}
		w.Write([]byte(say))
		w.Write([]byte("\n"))
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
	http.HandleFunc("/api/v1/info", info)
	http.HandleFunc("/api/v1/info/", info)
	http.HandleFunc("/api/v1/ping", ping)
	http.HandleFunc("/api/v1/ping/", ping)
	http.HandleFunc("/api/v1/count", counter)
	http.HandleFunc("/api/v1/count/", counter)
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

func getFortune(f string) string {
	cmd := exec.Command("fortune", f)
	//var stdout, stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	//cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)
	}
	//outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	//fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	//outStr := string(stdout.Bytes())
	tmpStr := string(stdout.Bytes())
	outStr := fmt.Sprintf(strings.TrimSpace(tmpStr))
	//fmt.Printf("out:\n%s\n", outStr)
	fmt.Printf("out:\n%s", outStr)

	mu.Lock()
	count++
	mu.Unlock()

	return outStr
}

//func uploadFile(f string) string {
//	fmt.Println("file: ", f)
//	log.Println("Upload attempt")
//	return f
//}

func viewErrLines(errLines map[int]error) []string {
	var flatBadLine []string
	for i, j := range errLines {
		k := strconv.Itoa(i)
		l := j.Error()
		badString := k + "\t " + l + "\n"
		fmt.Println(badString)
		flatBadLine = append(flatBadLine, badString)
	}
	return flatBadLine
}

// info handler displays http header
func info(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// counter displays the page count
func counter(w http.ResponseWriter, r *http.Request) {
	//mu.Lock()
	//count++
	//mu.Unlock()
	fmt.Fprintf(w, "Count %d\n", count)
	//log.Printf("Count %d\n", count)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
	//log.Println("pong")
}

func readInFile(i string) []string {
	// Read line-by-line
	var lines []string
	file, err := os.Open(i)
	if err != nil {
		log.Println(err)
		return lines
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
