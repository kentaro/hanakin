package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var y, m, p int

func init() {
	flag.IntVar(&y, "year", 2014, "The year you want to know about")
	flag.IntVar(&m, "month", 9, "The month you want to know about")
	flag.IntVar(&p, "payday", 10, "The payday of your company")
}

func main() {
	flag.Parse()
	now := time.Now()

	if y == 0 {
		y = now.Year()
	}

	if m != 0 {
		if m < 1 || m > 12 {
			log.Fatalf("month must be between 1 and 12: %s", m)
		}
	} else {
		m = int(now.Month())
	}

	calendar := NewCalendar(time.Month(m), y, p)
	fmt.Println(calendar.String())
}

type Calendar struct {
	year  int
	month *month
}

func NewCalendar(m time.Month, y, p int) (c *Calendar) {
	c = &Calendar{year: y, month: NewMonth(m, y, p)}
	return
}

func (c *Calendar) String() (s string) {
	s = fmt.Sprintf("%d %s\n", c.year, c.month.number.String())
	s += "====================\n"
	s += c.month.String()
	return
}

type month struct {
	year   int
	number time.Month
	days   []time.Time
	payday int
}

func NewMonth(n time.Month, y, p int) (m *month) {
	m = &month{year: y, number: n, payday: p}
	t := time.Date(y, n, 1, 0, 0, 0, 0, time.Local)

	for {
		m.days = append(m.days, t)
		t = time.Date(y, n, t.Day()+1, 0, 0, 0, 0, time.Local)

		if t.Month().String() != n.String() {
			break
		}
	}

	return
}

func (m *month) String() (s string) {
	for _, d := range m.days {
		if d.Weekday().String() == "Friday" && d.Day() == m.payday {
			s += "FizzBuzz\n"
		} else if d.Weekday().String() == "Friday" {
			s += "Fizz\n"
		} else if d.Day() == m.payday {
			s += "Buzz\n"
		} else {
			s += fmt.Sprintf("%d\n", d.Day())
		}
	}

	return
}
