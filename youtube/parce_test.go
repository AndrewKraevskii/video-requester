package youtube

import (
	"testing"
	"time"
)

func TestParceIso(t *testing.T) {
	duration, err := parceIso("P2DT1S")
	expect := 2*time.Hour*24+time.Second
	if err != nil {
		t.Error(err)
	}
	if duration != expect {
		t.Error("got ", duration, " expected ", expect)
	}
	duration, err = parceIso("P0DT4H5M43S")
	expect = 4*time.Hour+5*time.Minute+43*time.Second
	if err != nil {
		t.Error(err)
	}
	if duration != expect{
		t.Error("got ", duration, " expected ", expect)
	}
}
