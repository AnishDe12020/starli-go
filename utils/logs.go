package utils

import (
	"fmt"

	logsymbols "github.com/defaltd/log-symbols"
	"github.com/fatih/color"
)

func Success(str string) {
	fmt.Printf("%s %s\n", logsymbols.SUCCESS, color.GreenString(str))
}

func Error(str string) error {
	s := color.RedString(str)
	return fmt.Errorf("%s %s", logsymbols.ERROR, s)
}

func ErrorPrint(str string) {
	fmt.Println(Error(str))
}

func Info(str string) {
	color.Cyan(str)
}

func Warning(str string) {
	fmt.Printf("%s %s\n", logsymbols.WARNING, color.YellowString(str))
}

func Debug(str string) {
	color.Magenta(str)
}
