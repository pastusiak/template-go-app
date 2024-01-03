package desktop

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Example struct{}

func (e *Example) Window() fyne.Window {
	w := fyne.CurrentApp().NewWindow("Example window")

	w.Resize(fyne.Size{Width: 800, Height: 600})
	w.CenterOnScreen()
	w.SetContent(widget.NewLabel("Hello World!"))
	w.SetContent(widget.NewButton("Click Me!", click))

	w.SetOnClosed(fyne.CurrentApp().Quit)
	return w
}

func click() {
	w := fyne.CurrentApp().NewWindow("Modal window")

	w.Resize(fyne.Size{Width: 400, Height: 300})
	w.CenterOnScreen()

	b := widget.NewButton("Close Application", close)
	b.Resize(fyne.Size{Width: 100, Height: 30})

	w.SetContent(b)

	w.Show()
}

func close() {
	fyne.CurrentApp().Quit()
}
