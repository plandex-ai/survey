//go:build ignore

package main

import "github.com/plandex-ai/survey/v2"

func main() {
	color := ""
	prompt := &survey.Select{
		Message: "Choose a color:",
		Options: []string{
			"a",
			"b",
			"c",
			"d",
			"e",
			"f",
			"g",
			"h",
			"i",
			"j",
		},
	}
	survey.AskOne(prompt, &color)
}
