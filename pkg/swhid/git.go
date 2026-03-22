//go:build git

package swhid

import (
	"io"
	"io/fs"
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

func (repo *Repository) NewContentFromBlob(hash string) (*Content, error) {
	blob, err := repo.Repo.BlobObject(plumbing.NewHash(hash))
	if err != nil {
		return nil, err
	}
	reader, err := blob.Reader()
	if err != nil {
		return nil, err
	}
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return NewContent(bytes), nil
}

func (repo *Repository) NewDirectoryFromTree(hash string) (*Directory, error) {
	tree, err := repo.Repo.TreeObject(plumbing.NewHash(hash))
	if err != nil {
		return nil, err
	}
	entries := []*Entry{}
	for _, entry := range tree.Entries {
		if entry.Mode.IsFile() {
			content, err := repo.NewContentFromBlob(entry.Hash.String())
			if err != nil {
				return nil, err
			}
			mode, err := entry.Mode.ToOSFileMode()
			if err != nil {
				return nil, err
			}
			entries = append(entries, &Entry{
				name:   entry.Name,
				mode:   fs.FileMode(mode),
				target: []byte(content.Swhid().Hash),
			})
		} else {
			directory, err := repo.NewContentFromBlob(entry.Hash.String())
			if err != nil {
				return nil, err
			}
			mode, err := entry.Mode.ToOSFileMode()
			if err != nil {
				return nil, err
			}
			entries = append(entries, &Entry{
				name:   entry.Name,
				mode:   mode,
				target: []byte(directory.Swhid().Hash),
			})
		}
	}
	return NewDirectory(entries), nil
}

func (repo *Repository) Head() (string, error) {
	head, err := repo.Repo.Head()
	if err != nil {
		return "", err
	}
	return head.Strings()[1], nil
}

func (repo *Repository) Commit(name string) (string, error) {
	ref, err := repo.Repo.Tag(name)
	if err != nil {
		return "", err
	}
	tag, err := repo.Repo.TagObject(ref.Hash())
	if err != nil {
		return "", err
	}
	return tag.Target.String(), nil
}

func (repo *Repository) Tree(hash string) (string, error) {
	commit, err := repo.Repo.CommitObject(plumbing.NewHash(hash))
	if err != nil {
		return "", err
	}
	return commit.TreeHash.String(), nil
}

func (repo *Repository) NewDirectoryFromHash(hash string) (*Directory, error) {
	tree, err := repo.Tree(hash)
	if err != nil {
		return nil, err
	}
	return repo.NewDirectoryFromTree(tree)
}

func (repo *Repository) NewRevisionFromTag(name string) (*Revision, error) {
	ref, err := repo.Repo.Tag(name)
	if err != nil {
		return nil, err
	}
	tag, err := repo.Repo.TagObject(ref.Hash())
	if err != nil {
		return nil, err
	}
	return repo.NewRevisionFromHash(tag.Target.String())
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
