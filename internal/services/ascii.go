package services

import "io/ioutil"

func LoadASCII(path string) (string, error) {
	b, err := ioutil.ReadFile("assets/ascii.txt")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
