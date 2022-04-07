package golog

import "github.com/gookit/color"

func Success(message string) {
	color.Success.Print(message, "\n")

	color.Success.Prompt(message, "\n")

	color.Success.Tips(message, "\n")
}
