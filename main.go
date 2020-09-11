package main

import (
	"fmt"
	"time"

	capture "github.com/preet-maiya/clipboard-manager/clipboardCapture"
	_ "github.com/preet-maiya/clipboard-manager/keyListener"
)

func main() {
	listeningInterval := 300 * time.Millisecond
	clipboard := capture.NewClipboard(listeningInterval)
	clipboard.StartListener()
	for i := 0; i < 10; i++ {
		fmt.Println("Clipboard: ", clipboard.GetContent())
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Pausing listener")
	clipboard.PauseListener()

	for i := 0; i < 10; i++ {
		fmt.Println("Clipboard: ", clipboard.GetContent())
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Resuming listener")
	clipboard.ResumeListener()

	for i := 0; i < 10; i++ {
		fmt.Println("Clipboard: ", clipboard.GetContent())
		time.Sleep(1 * time.Second)
	}
}
