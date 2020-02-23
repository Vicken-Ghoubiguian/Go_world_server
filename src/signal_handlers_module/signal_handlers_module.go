package signal_handlers_module

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//Function to handle Ctrl+c signal
func Setup_ctrl_c_handler() {

	c := make(chan os.Signal, 2)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {

		<-c

		fmt.Println("\rGoodbye, we will miss you...\r")

		os.Exit(0)
	}()
}
