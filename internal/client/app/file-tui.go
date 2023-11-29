package app

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"sort"

	"github.com/KryukovO/goph-keeper/api/serverpb"
	"github.com/KryukovO/goph-keeper/internal/entities"
	"github.com/KryukovO/goph-keeper/pkg/utils"
	"github.com/rivo/tview"
)

func (a *App) setupBinaryDataMenu() {
	a.flex.Clear()
	a.list.Clear()

	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RequestTimeout)
	defer cancel()

	resp, err := a.client.FileNamesList(ctx, nil)
	if err != nil {
		a.logCh <- err.Error()

		a.setupMainMenu()
	}

	fileNames := resp.GetFileNames()
	sort.Strings(fileNames)

	for index, file := range fileNames {
		a.list.AddItem(file, "", rune(49+index), nil)
	}

	a.flex.SetDirection(tview.FlexRow).
		AddItem(a.list, 0, 4, true).
		AddItem(tview.NewButton("Добавить").SetSelectedFunc(
			func() {
				a.setupBinaryDataForm()

				a.pages.SwitchToPage(formPage)
			},
		), 0, 1, false).
		AddItem(tview.NewButton("Скачать").SetSelectedFunc(
			func() {
				fileName := fileNames[a.list.GetCurrentItem()]

				a.logCh <- fmt.Sprintf("Начало скачивания файла %s в папку %s", fileName, a.cfg.FileStorage)

				go func() {
					ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RequestTimeout)
					defer cancel()

					stream, err := a.client.BinaryData(ctx, &serverpb.BinaryDataRequest{
						FileName: fileName,
					})
					if err != nil {
						a.logCh <- fmt.Sprintf("Ошибка скачивания файла %s: %s", fileName, err.Error())

						return
					}

					resp, err := stream.Recv()
					if err != nil {
						a.logCh <- fmt.Sprintf("Ошибка скачивания файла %s: %s", fileName, err.Error())

						return
					}

					file := entities.File{
						FileName: resp.GetData().GetFileName(),
					}

					fileData := bytes.Buffer{}

					for {
						resp, err := stream.Recv()
						if errors.Is(err, io.EOF) {
							file.Data = fileData

							break
						}

						if err != nil {
							a.logCh <- fmt.Sprintf("Ошибка скачивания файла %s: %s", file.FileName, err.Error())

							return
						}

						chunk := resp.GetData().GetChunk()

						_, err = fileData.Write(chunk)
						if err != nil {
							a.logCh <- fmt.Sprintf("Ошибка скачивания файла %s: %s", file.FileName, err.Error())

							return
						}
					}

					err = utils.SaveFile(a.cfg.FileStorage, file.FileName, file.Data)
					if err != nil {
						a.logCh <- fmt.Sprintf("Ошибка скачивания файла %s: %s", file.FileName, err.Error())

						return
					}

					a.logCh <- fmt.Sprintf("Файл %s успешно скачан в папку %s", file.FileName, a.cfg.FileStorage)
				}()
			},
		), 0, 1, false).
		AddItem(tview.NewButton("Удалить").SetSelectedFunc(
			func() {
				fileName := fileNames[a.list.GetCurrentItem()]

				ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RequestTimeout)
				defer cancel()

				_, err := a.client.DeleteBinaryData(ctx, &serverpb.DeleteBinaryDataRequest{
					FileName: fileName,
				})
				if err != nil {
					a.logCh <- err.Error()

					return
				}

				a.logCh <- "Файл удален"

				a.setupBinaryDataMenu()
			},
		), 0, 1, false).
		AddItem(tview.NewButton("Назад").SetSelectedFunc(
			func() {
				a.setupMainMenu()
			},
		), 0, 1, false)
}

func (a *App) setupBinaryDataForm() {
	var filePath string

	a.flex.Clear()
	a.form.Clear(true)

	a.flex.SetDirection(tview.FlexRow).
		AddItem(a.form, 0, 1, true)

	a.form.AddInputField("Путь до файла", "", 20, nil, func(path string) {
		filePath = path
	})

	a.form.AddButton("Сохранить", func() {
		a.logCh <- fmt.Sprintf("Начата загрузка файла %s", filePath)

		a.setupBinaryDataMenu()

		a.pages.SwitchToPage(menuPage)

		go func() {
			file, err := utils.GetFile(filePath)
			if err != nil {
				a.logCh <- err.Error()

				return
			}

			fileName := filepath.Base(filePath)

			ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RequestTimeout)
			defer cancel()

			stream, err := a.client.AddBinaryData(ctx)
			if err != nil {
				a.logCh <- fmt.Sprintf("Ошибка загрузки файла %s: %s", filePath, err.Error())

				return
			}

			err = stream.Send(&serverpb.AddBinaryDataRequest{
				Data: &serverpb.BinaryData{
					Data: &serverpb.BinaryData_FileName{FileName: fileName},
				},
			})
			if err != nil {
				a.logCh <- fmt.Sprintf("Ошибка загрузки файла %s: %s", filePath, err.Error())

				return
			}

			reader := bufio.NewReader(file)
			buffer := make([]byte, 1024)

			for {
				n, err := reader.Read(buffer)
				if err == io.EOF {
					break
				}

				if err != nil {
					a.logCh <- fmt.Sprintf("Ошибка загрузки файла %s: %s", filePath, err.Error())

					return
				}

				err = stream.Send(&serverpb.AddBinaryDataRequest{
					Data: &serverpb.BinaryData{
						Data: &serverpb.BinaryData_Chunk{Chunk: buffer[:n]},
					},
				})
				if err != nil {
					a.logCh <- fmt.Sprintf("Ошибка загрузки файла %s: %s", filePath, err.Error())

					return
				}
			}

			_, err = stream.CloseAndRecv()
			if err != nil {
				a.logCh <- fmt.Sprintf("Ошибка загрузки файла %s: %s", filePath, err.Error())

				return
			}

			a.logCh <- fmt.Sprintf("Файл %s успешно загружен на сервер", filePath)
		}()
	})

	a.form.AddButton("Назад", func() {
		a.setupMainMenu()

		a.pages.SwitchToPage(menuPage)
	})
}
