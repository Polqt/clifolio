package state

type Screen int 

const (
	ScreenIntro Screen = iota
	ScreenMenu
	ScreenProjects
	ScreenProjectDetail
	ScreenSkills
	ScreenExperience
	ScreenContact
	ScreenTheme
	ScreenAbout
)

func (s Screen) String() string {
	switch s {
	case ScreenIntro:
		return "Intro"
	case ScreenMenu:
		return "Menu"
	case ScreenProjects:
		return "Projects"
	case ScreenProjectDetail:
		return "Project Detail"
	case ScreenSkills:
		return "Skills"
	case ScreenExperience:
		return "Experience"
	case ScreenContact:
		return "Contact"
	case ScreenTheme:
		return "Theme"
	case ScreenAbout:
		return "About"
	default:
		return "Unknown"
	}
}