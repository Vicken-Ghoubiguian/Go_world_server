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
	Welcome_indicator bool

}

func Display_template_function(defined_page Page, bruts_timezones_string string, w http.ResponseWriter) {

	t := template.New("New tmpl")

	if defined_page.Welcome_indicator {

		t = template.Must(t.ParseFiles("tmpl/maintmpl.tmpl", "tmpl/welcomebodytmpl.tmpl", "tmpl/footertmpl.tmpl"))

	} else {

		t = template.Must(t.ParseFiles("tmpl/maintmpl.tmpl", "tmpl/bodytmpl.tmpl", "tmpl/footertmpl.tmpl"))

	}

	err := t.ExecuteTemplate(w, "maintmpl", defined_page)

	if err != nil {

		log.Fatalf("Template execution: %s", err)

		logs_module.Writing_log_in_terminal_function(colors_in_terminal_module.Red, "An error occured, consult logs file to know more...")

		logs_module.Writing_log_in_log_file_function("An error occured, consult logs file to know more...")

		logs_module.Display_purple_line()

		 logs_module.Writing_log_in_log_file_function("--------------------")

	} else {

		if defined_page.Welcome_indicator {

			logs_module.Writing_log_in_terminal_function(colors_in_terminal_module.Green, "Display welcome page...")

			 logs_module.Writing_log_in_log_file_function("Display welcome page...")

                         logs_module.Display_purple_line()

			 logs_module.Writing_log_in_log_file_function("--------------------")

		} else {

			logs_module.Writing_log_in_terminal_function(colors_in_terminal_module.Green, "Display timezones [" + bruts_timezones_string + "]...")

			logs_module.Writing_log_in_log_file_function("Display timezones [" + bruts_timezones_string + "]...")

                        logs_module.Display_purple_line()

			logs_module.Writing_log_in_log_file_function("--------------------")

		}
	}
}
