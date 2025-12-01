package services

import "io/ioutil"

func SaveTheme(path, name string) error {
	return ioutil.WriteFile(path, []byte(name), 0644)
}

func LoadTheme(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
