package template_displayer_module

import (
	"html/template"
	"net/http"
	"logs_module"
	"colors_in_terminal_module"
	"log"
)

type Page struct {

        Title string
        Main_section_title string
        Tz_array []string

}

func Display_template_function(defined_page Page, bruts_timezones_string string, w http.ResponseWriter) {

	t := template.New("New tmpl")

	t = template.Must(t.ParseFiles("tmpl/maintmpl.tmpl", "tmpl/bodytmpl.tmpl"))

	err := t.ExecuteTemplate(w, "maintmpl", defined_page)

	if err != nil {

		log.Fatalf("Template execution: %s", err)

		logs_module.Writing_log_in_terminal_function(colors_in_terminal_module.Red, "An error occured, consult logs file to know more...")

		logs_module.Display_purple_line()

	} else {

		if len(defined_page.Tz_array) > 0 {

			logs_module.Writing_log_in_terminal_function(colors_in_terminal_module.Green, "Display timezones [" + bruts_timezones_string + "]...")

			logs_module.Display_purple_line()

		} else {

			logs_module.Writing_log_in_terminal_function(colors_in_terminal_module.Green, "Display welcome page...")

			logs_module.Display_purple_line()

		}
	}
}
