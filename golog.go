package golog

import (
	"github.com/gookit/color"
)

func Success(message string) {
	if message == "" {
		message = ""
	}
	color.Success.Tips(message)
}

func Error(message string) {
	if message == "" {
		message = ""
	}
	color.Danger.Tips(message)
}

func Info(message string) {
	if message == "" {
		message = ""
	}
	color.Info.Tips(message)
	color.Primary.Tips(message)
}
