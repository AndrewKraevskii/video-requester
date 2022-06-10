package save

import (
	"encoding/json"
	"os"
)

// TODO: load tokens from save file
// TODO: if no saving get new

type User struct {
	Username     string `json:"username"`
	Access_token string `json:"access_token"`
	Id           string `json:"userid"`
}

const defoultFileName = ".save.json"

func Save(user *User, filename ...string) error {
	// 0666 - write & read
	name := defoultFileName
	if len(filename) != 0 {
		name = filename[0]
	}
	file, err := os.Create(name)
	if err != nil {
		return err
	}

	j, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		return err
	}
	_, err = file.Write(j)
	if err != nil {
		return err
	}
	err = file.Close()
	return err
}

func Load(filename ...string) (*User, error) {
	name := defoultFileName
	if len(filename) != 0 {
		name = filename[0]
	}

	b, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	user := &User{}
	err = json.Unmarshal(b, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
