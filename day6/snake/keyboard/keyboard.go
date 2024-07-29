package keyboard

import (
	"errors"
	"os"

	"github.com/eiannone/keyboard"
)

type keyboardEvent = keyboard.Key

const (
	KeyArrowDown  keyboardEvent = keyboard.KeyArrowDown
	KeyArrowUp    keyboardEvent = keyboard.KeyArrowUp
	KeyArrowLeft  keyboardEvent = keyboard.KeyArrowLeft
	KeyArrowRight keyboardEvent = keyboard.KeyArrowRight
)

type KeyboardHandler struct {
	callbacks map[keyboardEvent]func()
}

func New() *KeyboardHandler {
	return &KeyboardHandler{
		callbacks: make(map[keyboardEvent]func()),
	}
}

func (k *KeyboardHandler) AddEventListener(event keyboardEvent, callback func()) error {
	if callback == nil {
		return errors.New("callback function can't be nil")
	}
	if _, exists := k.callbacks[event]; exists {
		return errors.New("event already exists")
	}
	k.callbacks[event] = callback

	return nil
}

func (k *KeyboardHandler) Listen() error {
	for {
		_, key, err := keyboard.GetSingleKey()
		if err != nil {
			return err
		}

		if key == keyboard.KeyCtrlC {
			k.Close()
			os.Exit(0)
		}

		callback := k.callbacks[key]
		if callback == nil {
			continue
		}
		callback()
	}
}

func (k *KeyboardHandler) Close() {
	keyboard.Close()
}
