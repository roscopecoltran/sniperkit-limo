package output

import (
	"reflect"
	"testing"
	"time"

	"github.com/roscopecoltran/sniperkit-limo/model"
	"github.com/stretchr/testify/assert"
)

var text Text

func TestTextDoesRegisterItself(t *testing.T) {
	assert.Equal(t, "*output.Text", reflect.TypeOf(ForName("text")).String())
}

func ExampleText_Inline() {
	text.Inline("This is inline")
	// Output: This is inline
}

func ExampleText_Info() {
	text.Info("This is info")
	// Output: This is info
}

func ExampleText_Tick() {
	text.Tick()
	// Output: .
}

func ExampleText_StarLine() {
	fullName := "roscopecoltran/sniperkit-limo"
	language := "Go"
	star := &model.Star{
		FullName:   &fullName,
		Stargazers: 1000000,
		Language:   &language,
	}
	text.StarLine(star)
	// Output: roscopecoltran/sniperkit-limo *:1000000 Go
}

func ExampleText_Star() {
	fullName := "roscopecoltran/sniperkit-limo"
	language := "Go"
	description := "A CLI for managing starred Git repositories"
	homepage := "https://github.com/roscopecoltran/sniperkit-limo"
	url := "https://github.com/roscopecoltran/sniperkit-limo.git"
	star := &model.Star{
		FullName:    &fullName,
		Stargazers:  1000000,
		Language:    &language,
		Description: &description,
		Homepage:    &homepage,
		URL:         &url,
		StarredAt:   time.Date(2016, time.June, 21, 14, 56, 5, 0, time.UTC),
		Tags: []model.Tag{
			{
				Name: "cli",
			},
			{
				Name: "git",
			},
		},
	}
	text.Star(star)
	// Output:
	// roscopecoltran/sniperkit-limo *:1000000 Go https://github.com/roscopecoltran/sniperkit-limo.git
	// cli, git
	// A CLI for managing starred Git repositories
	// Home page: https://github.com/roscopecoltran/sniperkit-limo
	// Starred on Tue Jun 21 14:56:05 UTC 2016
}
