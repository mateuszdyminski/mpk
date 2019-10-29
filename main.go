package main

import (
	"net/http"
	"time"

	"encoding/json"
	"fmt"

	"bytes"
	"io/ioutil"
	"net/url"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Handle routes
	r.HandleFunc("/api/positions", getPositions)
	r.Handle("/{path:.*}", http.FileServer(http.Dir("statics")))

	http.Handle("/", loggingHandler{r})

	// Listen on hostname:port
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalf("Error: %s", err)
	}
}

type loggingHandler struct {
	http.Handler
}

func (h loggingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// path := req.URL.Path
	t := time.Now()
	h.Handler.ServeHTTP(w, req)
	log.Printf("%s [%s] \"%s %s %s\" \"%s\" \"%s\" \"Took: %s\"", req.RemoteAddr,
		t.Format("02/Jan/2006:15:04:05 -0700"), req.Method, req.RequestURI, req.Proto, req.Referer(), req.UserAgent(), time.Since(t))
}

func getPositions(w http.ResponseWriter, req *http.Request) {
	v := vehicles{}
	if err := json.NewDecoder(req.Body).Decode(&v); err != nil {
		fmt.Fprintf(w, "Can't decode json. Err: %v", err)
		return
	}

	log.Infof("Query for positions: %+v", v)

	form := url.Values{}
	for _, no := range v.Trains {
		form.Add("busList[train][]", no)
	}

	for _, no := range v.Trams {
		form.Add("busList[tram][]", no)
	}

	for _, no := range v.Buses {
		form.Add("busList[bus][]", no)
	}
	client := &http.Client{}
	r, _ := http.NewRequest("POST", "http://mpk.wroc.pl/position.php", bytes.NewBufferString(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	resp, err := client.Do(r)
	if err != nil {
		fmt.Fprintf(w, "Can't send request. Err: %v", err)
		return
	}

	defer resp.Body.Close()
	buf, _ := ioutil.ReadAll(resp.Body)
	w.Write(buf)
}

type vehicles struct {
	Trams  []string `json:"tram,omitempty"`
	Trains []string `json:"train,omitempty"`
	Buses  []string `json:"bus,omitempty"`
}
