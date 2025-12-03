package components

type Keymap struct {
	Toggle	string
	Confirm string
	Left 	string
	Right   string
	Up 	 	string
	Down	string
	Back 	string
	Quit	string
}

func DefaultKeymap() Keymap {
	return Keymap{
		Toggle: "/",
		Confirm: "enter",
		Left: "h",
		Right: "l",
		Up: "k",
		Down: "j",
		Back: "b",
		Quit: "q",
	}
}

