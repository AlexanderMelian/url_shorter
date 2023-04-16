package models

import "time"

type Session struct {
	Username string
	Expires  time.Time
}

func (s Session) isExpired() bool {
	return s.Expires.Before(time.Now().UTC())
}
