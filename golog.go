package golog

import (
	"github.com/gookit/color"
)

func Success(message string) {
	color.Green.Printf("SUCCESS: %s\n", message)
}

func Error(message string) {
	color.Red.Printf("ERROR: %s\n", message)
}

func Info(message string) {
	color.Cyan.Printf("INFO: %s\n", message)
}
