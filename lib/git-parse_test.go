package lib

import (
	"os"
	"testing"
)

func TestParseGit(t *testing.T) {
	d, _ := os.Getwd()
	t.Logf("Cur dir %s", d)
	v, e := ParseGit(d)
	if e != nil {
		t.Error(e)
	}
	t.Logf("%#+v", v)

}
