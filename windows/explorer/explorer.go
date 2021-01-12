// +build windows

package explorer

import (
	"github.com/pkg/errors"
	"golang.org/x/sys/windows/registry"
)

const (
	// User Shell Folders registry keys
	AppData      = `AppData`
	AppDataLocal = `Local AppData`
	Captures     = `{EDC0FE71-98D8-4F4A-B920-C8DC133CB165}`
	Desktop      = `Desktop`
	Documents    = `Personal`
	Downloads    = `{374DE290-123F-4565-9164-39C4925E467B}`
	MyMusic      = `My Music`
	MyPictures   = `My Pictures`
	MyVideo      = `My Video`
	Recent       = `Recent`
	SendTo       = `SendTo`
	StartMenu    = `Start Menu`
	Startup      = `Startup`
)

func GetUserShellFolder(key string) (string, error) {
	k, err := registry.OpenKey(
		registry.CURRENT_USER,
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\User Shell Folders`,
		registry.QUERY_VALUE,
	)
	if err != nil {
		return "", errors.Wrap(err, "cannot read registry value")
	}
	defer k.Close()

	path, _, err := k.GetStringValue(key)
	if err != nil {
		return "", errors.Wrap(err, "folder not found")
	}

	// expand %USERPROFILE%
	path, err = registry.ExpandString(path)
	if err != nil {
		return "", errors.Wrap(err, "cannot expand folder")
	}

	return path, nil
}
