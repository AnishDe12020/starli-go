package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func Success(str string) {
	color.Green(str)
}

func Error(str string) error {
	s := color.RedString(str)
	return fmt.Errorf(s)
}

func Info(str string) {
	color.Blue(str)
}

func Warning(str string) {
	color.Yellow(str)
}

func Debug(str string) {
	color.Magenta(str)
}
