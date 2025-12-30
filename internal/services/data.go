package services

type ProfileData struct {
	Name 		string
	Title		string
	Bio 		string
	Location	string
	Website		string
	Email		string
	GitHub 		string
	LinkedIn 	string
}

func GetProfileData() ProfileData {
	return ProfileData{
		Name: "Janpol Hidalgo",
		Title: "Software Engineer",
		Bio: "Full-stack developer passionate about building elegant solutions",
		Location: "Sagay City, Negros Occidental, Philippines",
		Website: "https://yojepoy.vercel.app/",
		Email: "poyhidalgo@gmail.com",
		GitHub: "github.com/Polqt",
		LinkedIn: "https://www.linkedin.com/in/janpol-hidalgo-64174a241/",
	}
}

type ExperienceItem struct {
	Position 	string
	Company 	string
	StartDate 	string
	EndDate 	string
	Description	string
	Tags		[]string
}

func GetExperiences() []ExperienceItem {
	return []ExperienceItem{
		{
			Position: 		"Mobile Developer",
			Company:  		"K92 Paints",
			StartDate: 		"December 2024",
			EndDate:   		"April 2025",
			Description: 	"Developed and maintained a mobile application using Flutter",
			Tags: 			[]string{"Flutter", "Mobile", "Dart"},		
		},
	}
}