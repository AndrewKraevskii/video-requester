package youtube

import (
	"regexp"
)

type IdNotFound struct{}

func (m *IdNotFound) Error() string {
	return "no youtube url in text"
}

func GetIdFromText(text string) (string, error) {
	reg := regexp.MustCompile(`((?:https?:)?\/\/)?((?:www|m)\.)?((?:youtube\.com|youtu.be))(\/(?:[\w\-]+\?v=|embed\/|v\/)?)([\w\-]+)(\S+)?`)

	result := reg.FindStringSubmatch(text)
	if len(result) <= 5 {
		return "", &IdNotFound{}
	}
	value := result[5]
	return value, nil
}
