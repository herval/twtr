package util

import (
	"fmt"
	"time"
)

func SplitText(text string, maxLineSize int) []string {
	res := []string{}

	i := 0
	for len(text) > 0 {
		l := text[0:Min(maxLineSize, len(text))]
		res = append(res, l)
		i += maxLineSize
		text = text[Min(maxLineSize, len(text)):]
	}

	return res
}

func TimeAgo(ts time.Time) string {
	dur := time.Since(ts)
	if dur.Minutes() < 60 {
		return fmt.Sprintf("%dm", int(dur.Minutes()))
	}
	if dur.Hours() < 24 {
		return fmt.Sprintf("%dh", int(dur.Hours()))
	}
	return fmt.Sprintf("%dd", int(dur.Hours()/24))
}
