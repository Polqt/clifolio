package components

type Keymap struct {
	Up 	 	string
	Down	string
	Open	string
	Back 	string
	Quit	string
}

func DefaultKeymap() Keymap {
	return Keymap{
		Up: "k",
		Down: "j",
		Open: "enter",
		Back: "b",
		Quit: "q",
	}
}

