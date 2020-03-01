package formats

import (
	"regexp"
)

func DateLayout(input string) string {
	dashedYMD := regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2})$`)
	dashedDMY := regexp.MustCompile(`^(\d{2})-(\d{2})-(\d{4})$`)
	slashedYMD := regexp.MustCompile(`^(\d{4})\/(\d{2})\/(\d{2})$`)
	slashedDMY := regexp.MustCompile(`^(\d{2})\/(\d{2})\/(\d{4})$`)

	layout := ""

	if dashedYMD.MatchString(input) {
		layout = "2006-01-02"
	} else if dashedDMY.MatchString(input) {
		layout = "02-01-2006"
	} else if slashedYMD.MatchString(input) {
		layout = "2006/01/02"
	} else if slashedDMY.MatchString(input) {
		layout = "02/01/2006"
	} else {
		layout = "2006-01-02"
	}

	return layout
}
