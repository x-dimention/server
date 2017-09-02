package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	str := getCurrentDirectory()
	strP := getParentDirectory(str)
	strPP := getParentDirectory(strP)
	client := strPP + "/client/dist"
	var f []byte
	var err error
	fmt.Print(r.URL.String())
	if a, _ := regexp.MatchString("static", r.URL.String()); a {
		css, _ := regexp.MatchString("css", r.URL.String())
		js, _ := regexp.MatchString("js", r.URL.String())
		if css {
			w.Header().Set("Content-Type", "text/css")
		}
		if js {
			w.Header().Set("Content-Type", "application/x-javascript")
		}

		f, err = ioutil.ReadFile(client + r.URL.String())
		if err != nil {
			fmt.Print(err)
		}
	} else {
		f, err = ioutil.ReadFile(client + "/index.html")
		if err != nil {
			fmt.Print(err)
		}
	}

	w.Write(f)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
