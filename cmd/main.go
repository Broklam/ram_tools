package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("RAM Tracker")
	myWindow.Resize(fyne.NewSize(300, 100))

	// Create a label widget to display RAM usage
	ramLabel := widget.NewLabel("")
	myWindow.SetContent(container.NewVBox(
		widget.NewLabel("System-wide RAM Usage:"),
		ramLabel,
	))

	// Periodically update the RAM usage label
	go func() {
		for {
			ramUsage, err := getRAMUsage()
			if err != nil {
				ramLabel.SetText("Error getting RAM usage")
			} else {
				ramLabel.SetText(fmt.Sprintf("%.2f MB", ramUsage))
			}
			time.Sleep(5 * time.Second)
		}
	}()

	myWindow.ShowAndRun()
}

// getRAMUsage returns the current system-wide RAM usage in MB.
func getRAMUsage() (float64, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return float64(memory.Used) / (1024 * 1024), nil
}
