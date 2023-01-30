package version

import (
	"fmt"
	"regexp"
	"time"
)

var (
	version = "dev"
	date    = ""
	commit  = ""
)

func Get() string {
	parsedDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		panic(err)
	}
	formattedDate := parsedDate.Format(time.RFC1123)
	return fmt.Sprintf("%s\n= Git commit: \t%s\n= Built: \t%s", version, commit, formattedDate)
}

func GetTag() string {
	tag := "master"
	m, err := regexp.MatchString(`\d\.\d\.\d$`, version)
	if err != nil {
		panic(err)
	}
	if m {
		tag = fmt.Sprintf("v%s", version)
	}
	return tag
}
