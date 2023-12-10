package app

import (
	"fmt"
	"time"

	"github.com/KryukovO/goph-keeper/api/serverpb"
	"github.com/KryukovO/goph-keeper/internal/client/config"
	"github.com/KryukovO/goph-keeper/internal/entities"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	menuPage = "Menu page"
	formPage = "Form page"
)

type App struct {
	client serverpb.KeeperClient
	user   entities.User
	token  string

	logCh chan string

	cfg *config.Config

	app   *tview.Application
	pages *tview.Pages
	log   *tview.TextView
	flex  *tview.Flex
	list  *tview.List
	form  *tview.Form
}

func NewApp(cfg *config.Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (a *App) Run() error {
	a.logCh = make(chan string, 1)

	defer close(a.logCh)

	conn, err := grpc.Dial(
		a.cfg.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(a.unaryInterceptor),
		grpc.WithStreamInterceptor(a.streamInterceptor),
	)
	if err != nil {
		return err
	}

	a.logCh <- "Клиент запущен"

	defer conn.Close()

	a.client = serverpb.NewKeeperClient(conn)

	return a.initInterface()
}

func (a *App) initInterface() error {
	a.app = tview.NewApplication()
	a.flex = tview.NewFlex()
	a.form = tview.NewForm()
	a.list = tview.NewList().ShowSecondaryText(false)

	a.pages = tview.NewPages().
		AddPage(menuPage, a.flex, true, true).
		AddPage(formPage, a.form, true, false)
	a.pages.SetBorder(true)

	a.log = tview.NewTextView().SetTextColor(tcell.ColorRed)
	a.log.SetBorder(true).SetTitle("Log").SetTitleAlign(tview.AlignLeft)

	go func() {
		for msg := range a.logCh {
			t := time.Now()
			text := fmt.Sprintf(
				"%s\n[%s] %s",
				a.log.GetText(false),
				t.Format("2006-01-02 15:04:05"),
				msg,
			)
			a.log.SetText(text)
		}
	}()

	a.setupWelcomeMenu()

	mainFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(a.pages, 0, 3, true).
		AddItem(a.log, 0, 1, false)

	return a.app.SetRoot(mainFlex, true).EnableMouse(true).Run()
}
