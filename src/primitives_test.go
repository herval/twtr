package main

import (
	"testing"
	"time"
)

func TestSplitText(t *testing.T) {
	lines := splitText("multi line text", 3)
	if len(lines) != 5 {
		t.Fail()
	}
	if lines[0] != "mul" {
		t.Fail()
	}
}

func TestTimeAgo(t *testing.T) {
	if timeAgo(time.Now().Add(-10 * time.Minute).Local())!= "10m" {
		t.Fail()
	}

	if timeAgo(time.Now().Add(-1 * time.Hour).Local())!= "1h" {
		t.Fail()
	}
}
