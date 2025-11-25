package services

import "github.com/charmbracelet/glamour"

func GenerateMarkdown(md string) (string, error) {
	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
	)
	if err != nil {
		return "", err
	}
	
	out, err := r.Render(md)
	if err != nil {
		return "", err
	}

	return out, nil
}