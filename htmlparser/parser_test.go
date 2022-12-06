package htmlparser

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var parseDataProvider = []struct {
	inputPath string
	output    []Link
}{
	{
		inputPath: "./test_cases/ex0.html",
		output: []Link{
			{
				Href: "/dog",
				Text: "Something in a spanText not in a spanBold text!",
			},
		},
	},
	{
		inputPath: "./test_cases/ex1.html",
		output: []Link{
			{
				Href: "/other-page",
				Text: "A link to another page",
			},
		},
	},
	{
		inputPath: "./test_cases/ex2.html",
		output: []Link{
			{
				Href: "https://www.twitter.com/joncalhoun",
				Text: "Check me out on twitter",
			},
			{
				Href: "https://github.com/gophercises",
				Text: "Gophercises is onGithub!",
			},
		},
	},
	{
		inputPath: "./test_cases/ex3.html",
		output: []Link{
			{
				Href: "https://twitter.com/marcusolsson",
				Text: "@marcusolsson",
			},
			{
				Href: "/lost",
				Text: "Lost? Need help?",
			},
			{
				Href: "#",
				Text: "Login",
			},
		},
	},
	{
		inputPath: "./test_cases/ex4.html",
		output: []Link{
			{
				Href: "/dog-cat",
				Text: "dog cat",
			},
		},
	},
}

func TestReader(t *testing.T) {
	for _, provider := range parseDataProvider {
		t.Run(provider.inputPath, func(t *testing.T) {
			// given
			htmlInput, err := os.Open(provider.inputPath)
			if err != nil {
				t.Fatalf("Failed to load example html: %s", provider.inputPath)
			}
			// when
			actual, _ := Parse(htmlInput)
			// then
			assert.ElementsMatch(t, actual, provider.output)
		})
	}
}
