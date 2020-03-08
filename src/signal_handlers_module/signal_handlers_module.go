package signal_handlers_module

import (
	"fmt"
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

		fmt.Println(colors_in_terminal_module.Cyan + "\rGoodbye, we will miss you (" + strconv.Itoa(os.Getpid()) + ")...\n" + colors_in_terminal_module.Reset)

		os.Exit(0)
	}()
}

//Function to handle Ctrl+Z signal
func Setup_ctrl_z_handler() {

	z := make(chan os.Signal, 20)

	signal.Notify(z, os.Interrupt, syscall.SIGTSTP)

	go func() {

		<-z

		fmt.Println(colors_in_terminal_module.Cyan + "\rPressed Ctrl+z, suspended process " + strconv.Itoa(os.Getpid()) + "...\n" + colors_in_terminal_module.Reset)

	}()
}
