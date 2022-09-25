package main

import (
	"fmt"
	"net/http"
)

func (app *application) CreateAndSendInvoice(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello, %s", "world")
	if err != nil {
		return
	}
}
