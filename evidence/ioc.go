package evidence

import (
	"time"
)

type IOC struct {
	Type   string
	Value  string
	Date   time.Time
	Source string
}

func NewIOC(t, value string, discovered time.Time, source string) IOC {
	return IOC{
		Type:   t,
		Value:  value,
		Date:   discovered,
		Source: source,
	}
}
