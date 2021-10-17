package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/api/sample", getHelloWorld)
	http.ListenAndServe(":" + port, nil)
}

func getHelloWorld(w http.ResponseWriter, _r *http.Request) {
	ping := map[string]string{"message": "Hello World!!"}
	res, _ := json.Marshal(ping)
	w.Write(res)
}
