package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (c *Controller) WelcomeScreen(w fyne.Window) fyne.CanvasObject {
	const (
		header1 = "## Greetings, brave adventurer!"
		text1   = "Whether you're stepping into the mystical realms of Dman for the first time or returning for another epic journey, we're thrilled to have you here. Prepare yourself for a world of magic, challenges, and camaraderie."
		header2 = "## Ready to Begin Your Journey?"
		text2   = "Follow the wizard to host your very own server or choose an existing server to join forces and forge alliances as you embark on quests together."
	)

	header1Label := widget.NewRichTextFromMarkdown(header1)
	// header1Label.Wrapping = fyne.TextWrapWord

	text1Label := widget.NewRichTextFromMarkdown(text1)
	text1Label.Wrapping = fyne.TextWrapWord

	header2Label := widget.NewRichTextFromMarkdown(header2)
	// header2Label.Wrapping = fyne.TextWrapWord

	text2Label := widget.NewRichTextFromMarkdown(text2)
	text2Label.Wrapping = fyne.TextWrapWord

	return container.NewBorder(
		nil,
		container.NewVBox(
			container.NewHBox(
				layout.NewSpacer(),
				widget.NewHyperlink("GitHub", parseURL("https://github.com/dimspell/gladiator")),
				widget.NewLabel("-"),
				widget.NewHyperlink("Discord", parseURL("https://discord.gg/XCNrwvdV6R")),
				widget.NewLabel("-"),
				widget.NewHyperlink("Reddit", parseURL("https://www.reddit.com/r/DispelRPG")),
				layout.NewSpacer(),
			),
			widget.NewLabelWithStyle("Version: "+c.version, fyne.TextAlignCenter, fyne.TextStyle{}),
		),
		nil,
		nil,
		container.NewPadded(
			container.NewVBox(
				widget.NewRichTextFromMarkdown("# Gladiator Dispel Multi"),

				container.NewVBox(
					header1Label,
					text1Label,
					header2Label,
					text2Label,
				),

				widget.NewLabel(""),

				container.NewHBox(
					layout.NewSpacer(),
					widget.NewButtonWithIcon("Start", theme.NavigateNextIcon(), func() {
						changePage(w, "Start", c.StartScreen(w, startOptionNone))
					}),
					layout.NewSpacer()),
			),
		),
	)
}
