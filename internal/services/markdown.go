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

func RenderMarkdown(md string) (string, error) {
	r, _ := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(80),
	)

	out, err := r.Render(md)
	if err != nil {
		return "", err
	}

	return out, nil
}