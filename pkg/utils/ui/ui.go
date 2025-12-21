package ui

type UI interface {
	Print(str string, format Format)
	Println(str string, format Format)
	EmptyLine()
}

func NewUI() UI {
	return DefaultUI{}
}
