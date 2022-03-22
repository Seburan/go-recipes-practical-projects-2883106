package main

import (
	"fmt"
	"time"
)

// tsConvert convert time stamp in "YYYY-MM-DDTHH:MM" format from one time zone to another
func tsConvert(ts, from, to string) (string, error) {

	const tsFormat = "2006-01-02T15:04"

	fromLoc, err := time.LoadLocation(from);
	if err != nil {
		return "", fmt.Errorf("error loading location %q: %w\n", from, err);
	}
	// load time location
	toLoc, err := time.LoadLocation(to);
	if err != nil {
		return "", fmt.Errorf("error loading location %q: %w\n", to, err);
	}

	// parse time given as argument "ts" from timezone specified as "from"
	fromTime, err := time.ParseInLocation(tsFormat, ts, fromLoc);
	if err != nil {
		return "", fmt.Errorf("error parsing time %q: %w\n", ts, err);
	}

	// convert time to timezone specified in argument "to"
	toTime := fromTime.In(toLoc);

	// convert back to original formatted string
	toTimeFormatted := toTime.Format(tsFormat);

	return toTimeFormatted, nil // FIXME
}

func main() {
	ts := "2021-03-08T19:12"
	out, err := tsConvert(ts, "Europe/Paris", "Asia/Tokyo");
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	fmt.Println(out) // 2021-03-09T05:12
}
