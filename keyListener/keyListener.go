package keyboardShortcuts

import (
	"errors"
	"fmt"

	"github.com/MarinX/keylogger"
	"github.com/sirupsen/logrus"
)

type shortcutFunc func() error

type Shortcut struct {
	KeyboardDevice string
	ShortcutMap    map[[]string]shortcutFunc
	Keyboard       string
	KeysPressed    []string
}

var specialKeys = map[string]string{
	"L_CTRL":  "",
	"R_CTRL":  "",
	"L_SHIFT": "",
	"R_SHIFT": "",
	"L_ALT":   "",
	"R_ALT":   "",
}

func (sc *Shortcut) RegisterShortcut(sequence []string, method shortcutFunc) error {
	if _, ok := sc.ShortcutMap[sequence]; ok {
		logrus.Warning("Sequence %v already registered")
		return errors.New("AlreadyExists")
	}
	sc.ShortcutMap[sequence] = method
	return nil
}

func (sc *Shortcut) CaptureKeys() error {

	// find keyboard device, does not require a root permission
	sc.Keyboard = keylogger.FindKeyboardDevice()

	// check if we found a path to keyboard
	if len(sc.Keyboard) <= 0 {
		logrus.Error("No keyboard found...you will need to provide manual input path")
		return errors.New("KeyboardDeviceNotFound")
	}

	logrus.Infof("Found a keyboard at", sc.Keyboard)
	// init keylogger with keyboard
	k, err := keylogger.New(keyboard)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	defer k.Close()

	events := k.Read()

	// range of events
	for e := range events {
		switch e.Type {
		// EvKey is used to describe state changes of keyboards, buttons, or other key-like devices.
		// check the input_event.go for more events
		case keylogger.EvKey:

			fmt.Sprintf(spew.Sdumpf(e))
			// if the state of key is pressed
			if e.KeyPress() {
				logrus.Println("[event] press key ", e.KeyString())
				sc.KeysPressed = append(e.KeyString(), sc.KeysPressed)
			}

			// if the state of key is released
			if e.KeyRelease() {
				logrus.Println("[event] release key ", e.KeyString())
				if _, ok := specialKeys[e.KeyString()]; ok {
					keyString := fmt.Sprintf("%s_RELEASE", e.KeyString())
					sc.KeysPressed = append(keyString, sc.KeysPressed)
				}
			}

			break
		}
	}
}

func (sc Shortcut) checkShortcut() bool {

}
