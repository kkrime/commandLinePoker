// package console is used to provide functions to write coloured text to the user for a better user exerpeicen
package console

import (
	"github.com/fatih/color"
)

var Error func(a ...interface{})
var BringWhite func(a ...interface{})
var Yellow func(a ...interface{})
var Green func(a ...interface{})
var Blue func(a ...interface{})

// init sets all the functions for writing to the console using colored output
// NOTE: all the code inside init is deterministic
func init() {
	Error = color.New(color.Bold, color.FgRed).PrintlnFunc()
	BringWhite = color.New(color.Bold, color.FgHiWhite).PrintlnFunc()
	Yellow = color.New(color.Bold, color.FgYellow).PrintlnFunc()
	Green = color.New(color.Bold, color.FgGreen).PrintlnFunc()
	Blue = color.New(color.Bold, color.FgBlue).PrintlnFunc()
}
