package lib

import (
	"time"
)

type Info struct {
	Branch string
	Commit string
	Tag    string
	Clean  bool

	CommitDate time.Time
	BuildDate  time.Time
}
