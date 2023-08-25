package main

import (
	"image"
	"image/color"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

// func checkError(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func colorIt(rstart int, rend int, cstart int, cend int, img *image.RGBA, clr *color.RGBA, wg *sync.WaitGroup, fi *canvas.Image) {
	defer wg.Done()
	for i := rstart; i <= rend; i++ {
		for j := cstart; j <= cend; j++ {
			img.Set(i, j, clr)
			fi.Refresh()
			time.Sleep(time.Nanosecond)
		}
	}
}

func main() {

	width := 600
	height := 600

	img := image.NewRGBA(image.Rect(1, 1, width, height))
	// Create a new image of type RGBA which takes a parameter of rect which creates a rect with specified minimum and maximum coordinates

	clr := &color.RGBA{255, 255, 255, 255}
	for i := 1; i <= 600; i++ {
		for j := 1; j <= 600; j++ {
			img.Set(i, j, clr)

		}
	}

	myGui := app.New() // Initialising the fyne class
	window := myGui.NewWindow("Thread Visualizer")

	fyneImage := canvas.NewImageFromImage(img)

	window.SetContent(fyneImage)

	window.Resize(fyne.NewSize(600, 600))

	red := &color.RGBA{255, 0, 0, 255}
	blue := &color.RGBA{0, 0, 255, 255}
	green := &color.RGBA{0, 255, 0, 255}
	randColor := &color.RGBA{128, 124, 234, 255}

	wg := &sync.WaitGroup{}
	wg.Add(4)
	go colorIt(1, 300, 1, 300, img, red, wg, fyneImage)

	go colorIt(301, 600, 1, 300, img, blue, wg, fyneImage)

	go colorIt(1, 300, 301, 600, img, green, wg, fyneImage)

	go colorIt(301, 600, 301, 600, img, randColor, wg, fyneImage)

	window.ShowAndRun()

	wg.Wait()

}
