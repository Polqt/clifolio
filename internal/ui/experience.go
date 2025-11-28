package ui


type Experience struct {
	Title       string
	Company     string
	Years 		string
	Description string
}
type experienceModel struct {
	cursor int
	experiences []Experience
}

