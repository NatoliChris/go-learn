package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00",
	)

	p(t1)

	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(),
	)
}
