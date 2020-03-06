package template_displayer_module

import (
	"html/template"
	"net/http"
	"log"
)

type Page struct {

        Title string
        Main_section_title string
        Tz_array []string

}

func Display_template_function(defined_page Page, w http.ResponseWriter) {

	t := template.New("New tmpl")

	t = template.Must(t.ParseFiles("tmpl/maintmpl.tmpl", "tmpl/bodytmpl.tmpl"))

	err := t.ExecuteTemplate(w, "maintmpl", defined_page)

	if err != nil {

		log.Fatalf("Template execution: %s", err)
	}
}
