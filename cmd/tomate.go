package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
)

var tomateCmd = &cobra.Command{
	Use:   "tomate",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                  love by spf13 and friends in Go.
                  Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {

		var minutes time.Duration = 25
		var seconds time.Duration = 0
		fmt.Println("Start")
		close := make(chan os.Signal)
		notifyChannel := make(chan bool)
		go timer(minutes, seconds, notifyChannel)
		go exit(close)
		go notify(notifyChannel, 2, close)
		select {
		case <-time.After(minutes*time.Minute + seconds*time.Second):
			close <- os.Interrupt
			fmt.Println("Stop")
		case <-notifyChannel:
			fmt.Println("main notify")
			return
		}
	},
}

func timer(minutes time.Duration, seconds time.Duration, notifyChannel chan (bool)) {
	fmt.Println(minutes*time.Minute + seconds*time.Second)
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			seconds = seconds - 1
			fmt.Println(minutes*time.Minute + seconds*time.Second)
		case <-notifyChannel:
			fmt.Println("timer notify")
			return
		}
	}
}

func notify(n chan (bool), numberToNotify int, close chan (os.Signal)) {
	<-close
	for i := 0; i < numberToNotify; i++ {
		n <- false
	}

}

func exit(close chan (os.Signal)) {
	signal.Notify(close, os.Interrupt)
}

func init() {
	rootCmd.AddCommand(tomateCmd)
}
