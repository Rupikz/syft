package main

import (
	"context"
	"os"

	"github.com/mholt/archives"
)

func main() {
	z := archives.Zip{
		SelectiveCompression: true,
		ContinueOnError:      false,
	}

	ctx := context.Background()
	archive, err := os.Create("test.zip")
	if err != nil {
		panic(err)
	}

	archiveFiles, err := archives.FilesFromDisk(ctx, nil, map[string]string{"main.go": ""})
	if err != nil {
		panic(err)
	}

	if err = z.Archive(ctx, archive, archiveFiles); err != nil {
		panic(err)
	}
}
