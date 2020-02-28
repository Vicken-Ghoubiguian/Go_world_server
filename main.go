package main

import (
	"log"
	"strings"
	"html/template"
	"net/http"
	"treatment_on_timezones_module"
	"signal_handlers_module"
)

type Page struct {

	Title string
	Main_section_title string
	Tz_array []string

}

func main() {

	signal_handlers_module.Setup_ctrl_c_handler()

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))

	http.HandleFunc("/", handlerFunction)

	http.ListenAndServe(":8080", nil)

}

func handlerFunction(w http.ResponseWriter, r *http.Request) {

	bruts_timezones_array := strings.Split(r.URL.Path[1:], "/")

	master_array := treatment_on_timezones_module.Master_function(bruts_timezones_array)

	p := Page{"Go world server", "Bienvenue sur Go world server, le serveur mondial qui déchire !!!!", master_array}

	t := template.New("New tmpl")

	t = template.Must(t.ParseFiles("tmpl/maintmpl.tmpl", "tmpl/bodytmpl.tmpl"))

	err := t.ExecuteTemplate(w, "maintmpl", p)

	if err != nil {
        	log.Fatalf("Template execution: %s", err)
    	}
}
