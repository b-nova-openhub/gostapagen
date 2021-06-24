package repo

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"os"
)

func Clone(toMemory bool, path string) error {
	if toMemory {
		return cloneToMemory()
	}
	return cloneToFilesystem(path)
}

func cloneToMemory() error {
	fmt.Println("cloning to memory...")
	_, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: "raffaelschneider",
			Password: "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",
		},
		URL:      "https://github.com/b-nova-openhub/jams-vanilla-content",
		Progress: os.Stdout,
	})
	return err
}

func cloneToFilesystem(path string) error {
	fmt.Println("cloning to filesystem...")
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: "raffaelschneider",
			Password: "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",
		},
		URL:      "https://github.com/b-nova-openhub/jams-vanilla-content",
		Progress: os.Stdout,
	})
	return err
}
