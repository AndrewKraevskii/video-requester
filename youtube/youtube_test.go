package youtube

import (
	"testing"
)

func TestYoutube(t *testing.T) {
	id, err := GetIdFromText("(`asdasfafhttps://www.youtube.com/watch?v=UkgK8eUdpAo asd")
	if err != nil {
		t.Fatal(err)
	}
	video, err := FetchVideo(id)
	if err != nil {
		t.Fatal(err)
	}
	duration, err := ParseDuration(video.ContentDetails.Duration)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(duration)
}
