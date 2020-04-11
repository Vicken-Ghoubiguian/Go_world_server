package logs_module

import (
	"fmt"
	"log"
	"os"
	"colors_in_terminal_module"
)

func Writing_log_in_terminal_function(logs_color string, logs_message string) {

	fmt.Println(logs_color + logs_message + colors_in_terminal_module.Reset)

}

func Writing_log_in_log_file_function(logs_message string) {

	f, err := os.OpenFile("log.txt", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	iLog := log.New(f, logs_message, log.LstdFlags)
	iLog.Println(logs_message + "\n")

}

func Display_purple_line() {

	fmt.Println(colors_in_terminal_module.Purple + "--------------------" + colors_in_terminal_module.Reset)

}
