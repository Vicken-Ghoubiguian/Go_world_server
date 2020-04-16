package signal_handlers_module

import (
	"logs_module"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"colors_in_terminal_module"
)

//Function to handle Ctrl+c signal
func Setup_ctrl_c_handler() {

	c := make(chan os.Signal, 2)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {

		<-c

		logs_module.Writing_log_in_terminal_function(colors_in_terminal_module.Cyan, "Goodbye, we will miss you (" + strconv.Itoa(os.Getpid()) + ")...\n")

		logs_module.Writing_log_in_log_file_function("Goodbye, we will miss you (" + strconv.Itoa(os.Getpid()) + ")...\n")

		os.Exit(0)
	}()
}

//Function to handle Ctrl+Z signal
func Setup_ctrl_z_handler() {

	z := make(chan os.Signal, 20)

	signal.Notify(z, os.Interrupt, syscall.SIGTSTP)

	go func() {

		<-z

		logs_module.Writing_log_in_terminal_function(colors_in_terminal_module.Cyan, "Pressed Ctrl+z, suspended process " + strconv.Itoa(os.Getpid()) + "...\n")

		logs_module.Writing_log_in_log_file_function("Pressed Ctrl+z, suspended process " + strconv.Itoa(os.Getpid()) + "...\n")

	}()
}
