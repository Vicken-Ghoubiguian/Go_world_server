package main

import (
	"fmt"
	"strings"
	"net/http"
	"treatment_on_timezones_module"
)

func main() {

	http.HandleFunc("/", handlerFunction)

	http.ListenAndServe(":8080", nil)

}

func handlerFunction(w http.ResponseWriter, r *http.Request) {

	bruts_timezones_array := strings.Split(r.URL.Path[1:], "/")

	master_string := treatment_on_timezones_module.Master_function(bruts_timezones_array)

	fmt.Fprintf(w, master_string)

}
