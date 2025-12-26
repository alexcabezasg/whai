package ui

type UI interface {
	Print(str string, format Format)
	Println(str string, format Format)
	EmptyLine()
	RunWithSpinner(spinnerInfo string, fn func() error)
}

func NewUI() UI {
	return DefaultUI{}
}
