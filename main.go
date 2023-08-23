package main

import "net/http"

func main(){

	// Adicionar Handlers

	http.ListenAndServe(":8000", nil)
}