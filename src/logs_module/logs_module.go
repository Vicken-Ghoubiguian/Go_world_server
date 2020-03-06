package logs_module

import (
	"fmt"
	"colors_in_terminal_module"
)

func Writing_log_in_terminal_function(logs_color string, logs_message string) {

	fmt.Println(logs_color + logs_message + colors_in_terminal_module.Reset)

}
