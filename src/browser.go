package main

import (
	"errors"
	"os/exec"
	"runtime"
)

type OsAdapter struct {
	command string
	args    []string
}

func (osAdapter *OsAdapter) OpenUrl(url string) error {
	args := osAdapter.args
	args = append(args, url)

	return exec.Command(osAdapter.command, args...).Start()
}

type Browser struct {
	osAdapter *OsAdapter
}

func newBrowser() (*Browser, error) {
	switch runtime.GOOS {
	case "linux":
		return &Browser{&OsAdapter{command: "xdg-open"}}, nil
	case "windows":
		return &Browser{&OsAdapter{command: "rundll32", args: []string{"url.dll,FileProtocolHandler"}}}, nil
	default:
		return nil, errors.New("Unsupported platform")
	}
}

func (browser *Browser) OpenUrl(url string) error {
	return browser.osAdapter.OpenUrl(url)
}
