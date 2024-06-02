package main

import "net/http"

func handlerERR(w http.ResponseWriter, r *http.Request) {
	respondWithERROR(w, 400, "Something went wrong")
}
