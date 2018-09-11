package flag

import "net/url"

type URL struct {
	*url.URL
}

func (u *URL) Set(s string) error {
	if ur, err := url.Parse(s); err != nil {
		return err
	} else {
		u.URL = ur
	}
	return nil
}

func (u *URL) String() string {
	return u.URL.String()
}
