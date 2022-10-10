# Project content
Desktop application that displays and logs ip address and location using the fyne.io library in the Go programming language.

# Install
```
go get fyne.io/fyne/v2
```

# Versions used
Mainly used software versions,
```
Go version: 1.18.4
fyne.io version: 2.2.3
```
Other software and resources used are written on the `go.mod` page.

# API Usage
```go
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
```
