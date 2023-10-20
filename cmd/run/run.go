package run

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ayn2op/discordo/config"
	"github.com/ayn2op/discordo/ui"
	"github.com/rivo/tview"
)

var (
	discordState *State

	app      = tview.NewApplication()
	mainFlex *MainFlex
)

func Run(token string) error {
	path := config.DefaultPath()
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}

	err = config.Load(path)
	if err != nil {
		return err
	}

	path = config.DefaultLogPath()
	err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	log.SetOutput(f)
	log.SetFlags(log.LstdFlags | log.Llongfile)

	if token == "" {
		lf := ui.NewLoginForm()

		go func() {
			mainFlex = newMainFlex()
			if err := <-lf.Error; err != nil {
				app.Stop()
				log.Fatal(err)
			}

			if err := openState(<-lf.Token); err != nil {
				app.Stop()
				log.Fatal(err)
			}

			app.QueueUpdateDraw(func() {
				app.SetRoot(mainFlex, true)
			})
		}()

		app.SetRoot(lf, true)
	} else {
		mainFlex = newMainFlex()
		if err := openState(token); err != nil {
			app.Stop()
		}

		app.SetRoot(mainFlex, true)
	}

	app.EnableMouse(config.Current.Mouse)
	return app.Run()
}