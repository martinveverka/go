# Windows Explorer Registry Utils

Package `explorer` provides access to the Windows Explorer registry.

## Installation

```shell
go get -u github.com/martinveverka/go/windows/registry/explorer
```

## Quickstart

Here is a simple example of getting a current user Pictures folder.

```go
import (
	"github.com/martinveverka/go/windows/registry/explorer"
)

path, err := explorer.GetUserShellFolder(explorer.MyPictures)
if err != nil {
	panic(err)
}

fmt.Printf("My pictures folder is %q\n", path)
```
