package timer

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
)

var timerCmd = &cobra.Command{
	Use:   "timer",
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
		go notiy(notifyChannel, 2, close)
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

func Execute() {
	if err := timerCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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

func notiy(n chan (bool), numberToNotify int, close chan (os.Signal)) {
	select {
	case <-close:
		for i := 0; i < numberToNotify; i++ {
			n <- false
		}

	}
}

func exit(close chan (os.Signal)) {
	signal.Notify(close, os.Interrupt)
	return
}
