package ui

import "github.com/rivo/tview"

var LoginViaTokenLoginModalButton = "Login via token"
var LoginViaEmailPasswordLoginModalButton = "Login via email and password"

func NewLoginModal(onLoginModalDone func(buttonIndex int, buttonLabel string)) *tview.Modal {
	loginModal := tview.NewModal().
		SetText("Choose a login method:").
		AddButtons([]string{LoginViaTokenLoginModalButton, LoginViaEmailPasswordLoginModalButton}).
		SetDoneFunc(onLoginModalDone)
	loginModal.
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1)

	return loginModal
}