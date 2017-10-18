package flag

import (
	"time"
)

// DateFlag is a flag for date in 'yyyy-mm-dd' format
type DateFlag struct {
	time.Time
}

// Set implements flag.Value interface
func (d *DateFlag) Set(s string) error {
	tp, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	d.Time = tp
	return nil
}

// String implements flag.Value interface
func (t *DateFlag) String() string {
	return t.Time.Format("2006-01-02")
}
