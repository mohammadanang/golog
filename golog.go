package golog

import (
	"github.com/gookit/color"
)

func Success(message string) {
	color.Success.Tips(message)
}

func Error(message string) {
	color.Danger.Tips(message)
}

func Info(message string) {
	color.Primary.Print("INFO:", message, "\n")
}
