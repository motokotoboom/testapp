package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	location := r.URL.Query().Get("location")
	if location != "" {
		w.Header().Add("location", location)
	}

	answerStatus := r.URL.Query().Get("status")
	if answerStatus != "" {
		code, err := strconv.Atoi(answerStatus)
		if err != nil {
			code = 200
		}

		if code >= 100 && code <= 999 {
			w.WriteHeader(code)
		}
	}

	if r.ContentLength > 10000 {
		reqDump, err := httputil.DumpRequest(r, false)
		if err != nil {
			fmt.Fprintf(w, "Could not dump request")
			return
		}

		w.Write(reqDump)
		return
	}

	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Fprintf(w, "Could not dump request")
		return
	}

	w.Write(reqDump)
}
