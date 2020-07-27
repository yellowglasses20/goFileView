package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/rivo/tview"
)

func main() {
	os.Exit(run())
}

func run() int {

	app := tview.NewApplication()

	serviceList := tview.NewList()
	serviceList.SetTitle("- file list -").SetBorder(true)

	bucketList := tview.NewTextView()
	bucketList.SetTitle("- Text -").SetBorder(true)

	flex := tview.NewFlex().
		AddItem(serviceList, 20, 0, true).
		AddItem(bucketList, 0, 1, false)

	drawServiceList(serviceList, bucketList, app)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	return 0
}

func bucketFunc(bucketList *tview.TextView, f string, app *tview.Application) {
	bucketList.Clear()
	text, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	bucketList.SetTitle(f)
	bucketList.SetText(string(text))
	// app.SetFocus(bucketList)
}

func dirwalk(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	return files
}

func drawServiceList(serviceList *tview.List, bucketList *tview.TextView, app *tview.Application) {
	fileInfo := dirwalk("./")
	createServiceList(serviceList, bucketList, fileInfo, app, dirAbs("."))

}

func createServiceList(serviceList *tview.List, bucketList *tview.TextView, fileInfo []os.FileInfo, app *tview.Application, dir string) {
	serviceList.Clear()
	upDir(serviceList, bucketList, app, dir)
	selectServiceList(serviceList, bucketList, fileInfo, app, dir)
}

func dirAbs(path string) string {
	p, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func upDir(serviceList *tview.List, bucketList *tview.TextView, app *tview.Application, dir string) {
	filepath := filepath.Join(dir, "..")
	filepathAbs := dirAbs(filepath)
	serviceList.AddItem("../", "up", 0, func() {
		fileInfo := dirwalk(filepathAbs)
		createServiceList(serviceList, bucketList, fileInfo, app, filepathAbs)
	})
}

func selectServiceList(serviceList *tview.List, bucketList *tview.TextView, fileInfo []os.FileInfo, app *tview.Application, dir string) {
	for _, file := range fileInfo {
		serviceList.AddItem(file.Name(), file.Mode().String(), 0, func() {
			dirName, _ := serviceList.GetItemText(serviceList.GetCurrentItem())
			nextFilePath := filepath.Join(dir, dirName)
			fs, err := fileStat(nextFilePath)
			if err != nil {
				panic(err)
			}

			selectDirOrFile(fs, nextFilePath, serviceList, bucketList, app)
		})
	}
}

func fileStat(filePath string) (os.FileInfo, error) {
	f, _ := os.Open(filePath)
	defer f.Close()
	return f.Stat()
}

func selectDirOrFile(fs os.FileInfo, nextFilePath string, serviceList *tview.List, bucketList *tview.TextView, app *tview.Application) {
	if fs.IsDir() {
		f := dirwalk(nextFilePath)
		createServiceList(serviceList, bucketList, f, app, nextFilePath)
		return
	}
	bucketFunc(bucketList, nextFilePath, app)
}
