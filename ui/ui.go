// Copyright 2016 Albert Nigmatzianov. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func Ask(s string) string {
	fmt.Print(s + " ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		Term("Couldn't read the input", err)
	}
	return strings.TrimSpace(input)
}

func Error(message string, err error) {
	var out string
	if err == nil {
		out = message
	} else if message == "" {
		out = err.Error()
	} else {
		out = fmt.Sprintf("%v: %v", message, err)
	}
	Say(color.RedString(out))
}

func Quit() {
	os.Exit(0)
}

func Newline() {
	Say("")
}

func Say(s string) {
	fmt.Println(s)
}

func Success(s string) {
	Say(color.GreenString(s))
}

func Term(message string, err error) {
	Error(message, err)
	os.Exit(1)
}

func Warning(message string) {
	Say(color.YellowString(message))
}
