package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC,
	)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))

	p(now.Unix())
	p(now.UnixMilli())
	p(now.UnixNano())

	p(time.Unix(now.Unix(), 0))
	p(time.Unix(0, 0))

	p("-- time formating and parsing --")
	p(now.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-09-12T11:55:44+03:00",
	)
	p(t1, e)

	/*
		Format and Parse use example-based layouts. Usually you’ll use a constant from time for these layouts, but you can also supply custom layouts. Layouts must use the reference time Mon Jan 2 15:04:05 MST 2006 to show the pattern with which to format/parse a given time/string. The example time must be exactly as shown: the year 2006, 15 for the hour, Monday for the day of the week, etc.
	*/
	p(now.Format("3:04AM"))
	p(now.Format("Mon Jan 2 15:04:05 2006"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 42 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02d\n", now.Year(), now.Month(), now.Day())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e1 := time.Parse(ansic, "8:41")
	p(e1)
}
