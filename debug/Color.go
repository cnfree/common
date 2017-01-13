package debug

import "io"

type ConsoleColor int

const (
	COLOR_RED    ConsoleColor = 31
	COLOR_GREEN  ConsoleColor = 32
	COLOR_YELLOW ConsoleColor = 33
	COLOR_BLUE   ConsoleColor = 34
	COLOR_GRAY   ConsoleColor = 37
)

type debugColor struct {
	Red    ConsoleColor
	Green  ConsoleColor
	Yellow ConsoleColor
	Blue   ConsoleColor
	Gray   ConsoleColor
}

var DebugColor = debugColor{
	COLOR_RED, COLOR_GREEN, COLOR_YELLOW, COLOR_BLUE, COLOR_GRAY,
}

func ColorPrint(w io.Writer, s string, c ConsoleColor) {
	w.Write([]byte{0x1b, '[', byte('0' + c/10), byte('0' + c%10), 'm'})
	w.Write([]byte(s))
	w.Write([]byte("\x1b[0m"))
}
