package tui

import (
	"testing"
)

func TestMouseMode(t *testing.T) {
	term := NewAnsiTerminal(80, 24)

	// Initially disabled
	if term.IsMouseModeEnabled() {
		t.Error("Mouse mode should be disabled initially")
	}

	// Enable mode 1000 (Set Mode)
	term.Write([]byte("\x1b[?1000h"))
	if !term.IsMouseModeEnabled() {
		t.Error("Mouse mode should be enabled after CSI ? 1000 h")
	}

	// Disable mode 1000 (Reset Mode)
	term.Write([]byte("\x1b[?1000l"))
	if term.IsMouseModeEnabled() {
		t.Error("Mouse mode should be disabled after CSI ? 1000 l")
	}

	// Enable mode 1002
	term.Write([]byte("\x1b[?1002h"))
	if !term.IsMouseModeEnabled() {
		t.Error("Mouse mode should be enabled after CSI ? 1002 h")
	}
}
