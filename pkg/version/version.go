package version

import (
	"fmt"
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
