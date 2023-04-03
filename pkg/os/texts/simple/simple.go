package simple

type Simple struct {
	text string
}

func New(text string) Simple {
	return Simple{text}
}

func (s Simple) String() string {
	return s.text
}

func (s Simple) Empty() bool {
	return len(s.text) == 0
}
