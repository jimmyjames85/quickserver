package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jimmyjames85/eflag"
)

type eflags struct {
	Port    int    `flag:"port,p" desc:"port to serve on"`
	Message string `flag:"msg,m" desc:"message to return"`
}

func main() {
	s := &eflags{Port: 33333, Message: "quickserver is up"}
	eflag.StructVar(s)
	flag.Usage = eflag.POSIXStyle
	flag.Parse()
	qs := new(s.Message)
	http.HandleFunc("/", qs.HandleMessage)
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	return
}

func (qs *quickserver) HandleMessage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", qs.Message)
}

type quickserver struct {
	Message string
}

func new(msg string) *quickserver {
	return &quickserver{Message: msg}
}
