// Cowsayer  -  Cowsay As A Service

package cowsaid

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func GetFortune(f string) string {
	// fmt.Println("f: ", f)
	cmd := "fortune " + f
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return fmt.Sprintf("Failed to execute command: %s", cmd)
	}
	log.Printf("%v", out)

	mu.Lock()
	count++
	mu.Unlock()

	return string(out)
}

func CowFortune(f string) string {
	var cmd string
	if f != "" {
		// cmd = "fortune | cowsay " + fmt.Sprintf("%q", f)
		cmd = "cowsay " + fmt.Sprintf("%q", f)
	} else {
		cmd = "fortune | cowsay"
	}
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return fmt.Sprintf("Failed to execute command: %s", cmd)
	}

	mu.Lock()
	count++
	mu.Unlock()

	return string(out)
}

//func uploadFile(f string) string {
//	fmt.Println("file: ", f)
//	log.Println("Upload attempt")
//	return f
//}

func ViewErrLines(errLines map[int]error) []string {
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
func Info(w http.ResponseWriter, r *http.Request) {
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
func Counter(w http.ResponseWriter, r *http.Request) {
	//mu.Lock()
	//count++
	//mu.Unlock()
	fmt.Fprintf(w, "Count %d\n", count)
	//log.Printf("Count %d\n", count)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
	//log.Println("pong")
}

func ReadInFile(i string) []string {
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
