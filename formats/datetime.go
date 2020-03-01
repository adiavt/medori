package formats

import (
"regexp"
)

func DatetimeLayout(input string) string {
	dashedYMDHis := regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2}):(\d{2})$`)
	dashedDMYHis := regexp.MustCompile(`^(\d{2})-(\d{2})-(\d{4}) (\d{2}):(\d{2}):(\d{2})$`)
	slashedYMDHis := regexp.MustCompile(`^(\d{4})\/(\d{2})\/(\d{2}) (\d{2}):(\d{2}):(\d{2})$`)
	slashedDMYHis := regexp.MustCompile(`^(\d{2})\/(\d{2})\/(\d{4}) (\d{2}):(\d{2}):(\d{2})$`)

	layout := ""

	if dashedYMDHis.MatchString(input) {
		layout = "2006-01-02 15:04:05"
	} else if dashedDMYHis.MatchString(input) {
		layout = "02-01-2006 15:04:05"
	} else if slashedYMDHis.MatchString(input) {
		layout = "2006/01/02 15:04:05"
	} else if slashedDMYHis.MatchString(input) {
		layout = "02/01/2006 15:04:05"
	} else {
		layout = "2006-01-02 15:04:05"
	}

	return layout
}

