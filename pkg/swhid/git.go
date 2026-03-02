//go:build git
package swhid

import (
	"fmt"
	"strings"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type Repository struct {
	Repo *git.Repository
}

func NewRepository(path string) (*Repository, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	return &Repository{Repo: repo}, nil
}

func NewHashFromHash(hash plumbing.Hash) Hash {
	return hash[:]
}

func NewHashesFromHashes(hashes []plumbing.Hash) []Hash {
	result := []Hash{}
	for _, hash := range hashes {
		result = append(result, NewHashFromHash(hash))
	}
	return result
}

func (repo *Repository) Head() (string, error) {
	head, err := repo.Repo.Head()
	if err != nil {
		return "", err
	}
	return head.Strings()[1], nil
}

func (repo *Repository) NewRevisionFromHead() (*Revision, error) {
	ref, err := repo.Repo.Head()
	if err != nil {
		return nil, err
	}
	return repo.newRevisionFromObject(ref.Hash())
}

func (repo *Repository) NewRevisionFromHash(hash string) (*Revision, error) {
	return repo.newRevisionFromObject(plumbing.NewHash(hash))
}

func (repo *Repository) newRevisionFromObject(hash plumbing.Hash) (*Revision, error) {
	commit, err := repo.Repo.CommitObject(hash)
	if err != nil {
		return nil, err
	}
	rev := NewRevision()
	rev.Directory = NewHashFromHash(commit.TreeHash)
	rev.Parents = NewHashesFromHashes(commit.ParentHashes)
	rev.Author.Name = commit.Author.Name
	rev.Author.Email = commit.Author.Email
	rev.Author.Timestamp = commit.Author.When.UnixMilli()
	rev.Author.Offset = commit.Author.When.Format("-0700")
	rev.Committer.Name = commit.Committer.Name
	rev.Committer.Email = commit.Committer.Email
	rev.Committer.Timestamp = commit.Committer.When.UnixMilli()
	rev.Committer.Offset = commit.Committer.When.Format("-0700")
	rev.Message = commit.Message
	return rev, nil
}

func (repo *Repository) Tag(name string) (string, error) {
	tag, err := repo.Repo.Tag(name)
	if err != nil {
		return "", err
	}
	return tag.Strings()[1], nil
}

func (repo *Repository) NewReleaseFromTag(name string) (*Release, error) {
	ref, err := repo.Repo.Tag(name)
	if err != nil {
		return nil, err
	}
	return repo.newReleaseFromObject(ref.Hash())
}

func (repo *Repository) NewReleaseFromHash(hash string) (*Release, error) {
	return repo.newReleaseFromObject(plumbing.NewHash(hash))
}

func (repo *Repository) newReleaseFromObject(hash plumbing.Hash) (*Release, error) {
	tag, err := repo.Repo.TagObject(hash)
	if err != nil {
		return nil, err
	}
	rel := NewRelease()
	rel.Object = NewHashFromHash(tag.Target)
	rel.ObjectType = tag.TargetType.String()
	rel.Tag = tag.Name
	rel.Tagger.Name = tag.Tagger.Name
	rel.Tagger.Email = tag.Tagger.Email
	rel.Tagger.Timestamp = tag.Tagger.When.UnixMilli()
	rel.Tagger.Offset = tag.Tagger.When.Format("-0700")
	rel.Message = &tag.Message
	return rel, nil
}

func (repo *Repository) Tags() ([]string, error) {
	iter, err := repo.Repo.Tags()
	if err != nil {
		return nil, err
	}
	tags := []string{}
	err = iter.ForEach(func(ref *plumbing.Reference) error {
		tag := ref.Strings()[1]
		tags = append(tags, tag)
		return nil
	})
	return tags, err
}

func (repo *Repository) NewSnapshot() (*Snapshot, error) {
	branches := []*Branch{}
	head, err := repo.Repo.Head()
	if err != nil {
		return nil, err
	}
	branch := &Branch{Name: "HEAD"}
	branch.TargetType = "alias"
	branch.Target = []byte(head.Name().String())
	branches = append(branches, branch)
	iter, err := repo.Repo.References()
	if err != nil {
		return nil, err
	}
	err = iter.ForEach(func(ref *plumbing.Reference) error {
		name := ref.Strings()[0]
		refname := ref.Name()
		targetType := ""
		target := []byte{}
		if refname.IsBranch() {
			 _, err := repo.Repo.Branch(ref.Name().Short())
			 if err != nil {
				 return err
			 }
			targetType = "revision"
			rev, err := repo.NewRevisionFromHash(ref.Hash().String())
			 if err != nil {
				 return err
			 }
			target = rev.Swhid().Hash
		} else if refname.IsTag() {
			 _, err := repo.Repo.TagObject(ref.Hash())
			 if err == plumbing.ErrObjectNotFound {
				 return nil // unannotated tag
			 }
			 if err != nil {
				 return err
			 }
			targetType = "release"
			rel, err := repo.NewReleaseFromTag(ref.Name().Short())
			if err != nil {
				 return err
			}
			fmt.Printf("%s\n", name)
			target = rel.Swhid().Hash
		} else {
			return nil
		}
		branch := &Branch{Name: name}
		branch.TargetType = targetType
		branch.Target = target
		branches = append(branches, branch)
		return nil
	})
	if err != nil {
		return nil, err
	}
	snp := NewSnapshot(branches)
	return snp, nil
}

func (repo *Repository) Branches() ([]string, error) {
	iter, err := repo.Repo.Branches()
	if err != nil {
		return nil, err
	}
	branches := []string{}
	err = iter.ForEach(func(ref *plumbing.Reference) error {
		branch := ref.Strings()[0]
		branch = strings.Replace(branch, "refs/heads/", "", 1)
		branches = append(branches, branch)
		return nil
	})
	return branches, err
}
