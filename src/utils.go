package main

import "regexp"

const (
	numeralPattern = "^[ \t]*([a-zA-Z]+)[ \t]+is[ \t]+([IVXLCDM]+)[ \t]*$"
	creditsPattern = "^[ \t]*([a-zA-Z ]+)[ \t]+([a-zA-Z]+)[ \t]+is[ \t]+([0-9]+)[ \t]+[cC][rR][eE][dD][iI][tT][sS][ \t]*$"
	howMuchPattern = "^[ \t]*how[ \t]+much[ \t]+is[ \t]+([a-zA-Z ]+)[ \t]*\\?$"
	howManyPattern = "^[ \t]*how[ \t]+many[ \t]+[cC][rR][eE][dD][iI][tT][sS][ \t]+is[ \t]+([a-zA-Z ]+)[ \t]+([a-zA-Z]+)[ \t]*\\?$"
)

var (
	NumeralRegex = regexp.MustCompile(numeralPattern)
	CreditsRegex = regexp.MustCompile(creditsPattern)
	HowMuchRegex = regexp.MustCompile(howMuchPattern)
	HowManyRegex = regexp.MustCompile(howManyPattern)
)
