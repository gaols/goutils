package goutils

import (
	"github.com/tcnksm/go-input"
	"os"
	"fmt"
)

// Confirm force input string must be in positiveChoices or negativeChoices,
// if false choice is selected, negative action will be invoked.
// This function can be useful that if you want user must make a choice before doing anything further.
func Confirm(query string, positiveChoices, negativeChoices []string, negativeAction func()) {
	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	choices := append(positiveChoices, negativeChoices...)
	result, err := ui.Ask(query, &input.Options{
		Loop: true,
		Required: true,
		ValidateFunc: func(input string) error {
			if !ContainsStr(choices, input) {
				return fmt.Errorf("you must input %s", choices)
			}
			return nil
		},
	})

	if err != nil {
		RedText(err.Error())
		os.Exit(1)
	}

	if ContainsStr(negativeChoices, result) {
		negativeAction()
		os.Exit(0)
	}
}
