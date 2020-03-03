package main

import (
	"log"
	"fmt"
	"strings"
	"template_displayer_module"
	"net/http"
	"treatment_on_timezones_module"
	"signal_handlers_module"
)

func main() {

	signal_handlers_module.Setup_ctrl_c_handler()

	signal_handlers_module.Setup_ctrl_z_handler()

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))

	http.HandleFunc("/", handlerFunction)

	http.ListenAndServe(":8080", nil)

}

func handlerFunction(w http.ResponseWriter, r *http.Request) {

	bruts_timezones_array := strings.Split(r.URL.Path[1:], "/")

	welcome_indicator := false

	if r.URL.Path != "/" {

		welcome_indicator = true

	}

	if welcome_indicator {

		master_array := treatment_on_timezones_module.Master_function(bruts_timezones_array)

		err := template_displayer_module.Display_template_function(template_displayer_module.Page{"Go world server", "Bienvenue sur Go world server, le serveur mondial qui déchire !!!!", master_array}, w)

		if err != nil {

        		log.Fatalf("Template execution: %s", err)
    		}

	} else {

		fmt.Fprintf(w, "Aucune timezone à afficher !!!!")

	}
}
