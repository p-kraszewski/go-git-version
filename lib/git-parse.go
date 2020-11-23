package lib

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func ParseGit(d string) (*Info, error) {
	opts := git.PlainOpenOptions{
		DetectDotGit:          true,
		EnableDotGitCommonDir: true,
	}
	r, err := git.PlainOpenWithOptions(d, &opts)
	if err != nil {
		return nil, err
	}

	tagrefs, err := r.Tags()
	// CheckIfError(err)
	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		fmt.Printf("TAGREF: %#+v\n", t)
		return nil
	})
	// CheckIfError(err)

	// Print each annotated tag object (lightweight tags are not included)
	//	Info("for t in $(git show-ref --tag); do if [ \"$(git cat-file -t $t)\" = \"tag\" ]; then git cat-file -p $t ; fi; done")

	tags, err := r.TagObjects()
	// CheckIfError(err)
	err = tags.ForEach(func(t *object.Tag) error {
		fmt.Printf("TAG: %#+v\n", t)
		return nil
	})
	// CheckIfError(err)

	return nil, nil
}
