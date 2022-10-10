package main

// Required libraries
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// The part required to create the application and keep its size constant
	app := app.New()
	window := app.NewWindow("What is my IP?")
	window.Resize(fyne.NewSize(275, 250))
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.SetIcon(theme.ComputerIcon())

	// The place where the necessary variables and structure for recording the data are created
	type IP struct {
		Query   string
		Country string
		City    string
	}
	var Ip []IP
	// The place where the file to write the data is opened
	log, _ := ioutil.ReadFile("ipLogData.txt")
	json.Unmarshal(log, &Ip)

	// Where the widgets in the application are created
	labelTitle := widget.NewLabel("What is my IP?")
	labelIP := widget.NewLabel("Your IP is ...")
	label_Value := widget.NewLabel("...")
	label_City := widget.NewLabel("...")
	label_Country := widget.NewLabel("...")
	btn := widget.NewButton("Run", func() {
		// Actions that will take place when the button is pressed
		obje1 := &IP{
			Country: myContry(),
			City:    myCity(),
			Query:   myIP(),
		}
		Ip = append(Ip, *obje1)
		buffer, _ := json.MarshalIndent(Ip, "", " ")
		os.WriteFile("ipLogData.txt", buffer, 0644)
		label_Value.Text = myIP()
		label_Value.Refresh()
		label_City.Text = myCity()
		label_City.Refresh()
		label_Country.Text = myContry()
		label_Country.Refresh()
	})

	// Placement of widgets in the application
	window.SetContent(
		container.NewVBox(
			labelTitle,
			labelIP,
			label_Value,
			label_City,
			label_Country,
			btn,
		),
	)

	window.ShowAndRun()
}

// Pulling given with API
func myIP() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}
	var ip IP
	json.Unmarshal(body, &ip)
	return ip.Query
}

func myCity() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}
	var ip IP
	json.Unmarshal(body, &ip)
	return ip.City
}

func myContry() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}
	var ip IP
	json.Unmarshal(body, &ip)
	return ip.Country
}

type IP struct {
	Query   string
	Country string
	City    string
}
