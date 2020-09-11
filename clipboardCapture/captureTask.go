package capture

import (
	"log"
	"time"

	"github.com/atotto/clipboard"
)

type Clipboard struct {
	Content           string
	UpdatedOn         time.Time
	ListeningInterval time.Duration
	ListeningState    bool
}

func (clip *Clipboard) UpdateClipBoardContent() {
	clipContent, err := clipboard.ReadAll()
	if err != nil {
		log.Fatalf("Error reading clipboard content: %v", err)
	}
	clip.Content = clipContent
	clip.UpdatedOn = time.Now()
}

func (clip *Clipboard) StartListener() {
	go func(interval time.Duration) {
		for clip.ListeningState {
			clip.UpdateClipBoardContent()
			time.Sleep(interval)
		}
	}(clip.ListeningInterval)
}

func (clip Clipboard) GetContent() string {
	return clip.Content
}

func (clip *Clipboard) PauseListener() {
	clip.ListeningState = false
}

func (clip *Clipboard) ResumeListener() {
	clip.ListeningState = true
	clip.StartListener()
}

func NewClipboard(interval time.Duration) Clipboard {
	return Clipboard{
		ListeningInterval: interval,
		ListeningState:    true,
	}
}
