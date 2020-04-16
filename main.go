package main

import (
	"strings"
	"template_displayer_module"
	"net/http"
	"logs_module"
	"colors_in_terminal_module"
	"treatment_on_timezones_module"
	"signal_handlers_module"
)

func main() {

	signal_handlers_module.Setup_ctrl_c_handler()

	signal_handlers_module.Setup_ctrl_z_handler()

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))

	logs_module.Writing_log_in_terminal_function(colors_in_terminal_module.Green, "Connectez-vous à Go world server à l'adresse suivante: http://localhost:8080/...")

	logs_module.Writing_log_in_log_file_function("Connectez-vous à Go world server à l'adresse suivante: http://localhost:8080/...")

	logs_module.Display_purple_line()

	logs_module.Writing_log_in_log_file_function("--------------------")

	http.HandleFunc("/", handlerFunction)

	http.ListenAndServe(":8080", nil)

}

func handlerFunction(w http.ResponseWriter, r *http.Request) {

	bruts_timezones_array := strings.Split(r.URL.Path[1:], "/")

	welcome_indicator := false

	if r.URL.Path == "/" || r.URL.Path == "" {

		welcome_indicator = true

	}

	if !welcome_indicator {

		master_array := treatment_on_timezones_module.Master_function(bruts_timezones_array)

		bruts_timezones_string := treatment_on_timezones_module.Display_as_string_function(bruts_timezones_array)

		template_displayer_module.Display_template_function(template_displayer_module.Page{"Go world server", "Bienvenue sur Go world server, le serveur mondial qui déchire !!!!", master_array, welcome_indicator}, bruts_timezones_string, w)

	} else {

		template_displayer_module.Display_template_function(template_displayer_module.Page{"Go world server", "Bienvenue sur Go world server, le serveur mondial qui déchire !!!!", []string{}, welcome_indicator}, "Yessi !!!!", w)

	}
}
