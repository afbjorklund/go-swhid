//go:build git

package swhid

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func TestNewHashFromHash(t *testing.T) {
	h := Hash(make([]byte, 20))
	assert.Equal(t, h, NewHashFromHash(plumbing.ZeroHash))
	h, err := NewHash(NewObject("blob", []byte{}).Bytes())
	assert.Nil(t, err)
	assert.Equal(t, h, NewHashFromHash(plumbing.ComputeHash(plumbing.BlobObject, []byte{})))
}

func TestNewHashesFromHashes(t *testing.T) {
	h := Hash(make([]byte, 20))
	assert.Equal(t, []Hash{h}, NewHashesFromHashes([]plumbing.Hash{plumbing.ZeroHash}))
}

func gitInit(path string) (string, error) {
	_, err := git.PlainInit(path, false)
	return path, err
}

func TestGitRepository(t *testing.T) {
	path, err := gitInit(t.TempDir())
	assert.Nil(t, err)
	_, err = NewRepository(path)
	assert.Nil(t, err)
	_, err = NewRepository("nosuchdir")
	assert.Error(t, err)
}

func TestGitTree(t *testing.T) {
	path, err := gitInit(t.TempDir())
	assert.Nil(t, err)
	repo, err := NewRepository(path)
	assert.Nil(t, err)
	worktree, err := repo.Repo.Worktree()
	assert.Nil(t, err)
	err = os.WriteFile(path+"/file", []byte{}, 0o644)
	assert.Nil(t, err)
	blob, err := worktree.Add("file")
	assert.Nil(t, err)
	author := object.Signature{Name: "me", Email: "me@example.com"}
	opts := git.CommitOptions{Author: &author}
	hash, err := worktree.Commit("message", &opts)
	assert.Nil(t, err)
	head, err := repo.Head()
	assert.Nil(t, err)
	assert.Equal(t, hash.String(), head)
	dir, err := repo.NewDirectoryFromHash(hash.String())
	assert.Nil(t, err)
	cnt, err := repo.NewContentFromBlob(blob.String())
	assert.Nil(t, err)
	assert.Equal(t, "swh:1:dir:df2b8fc99e1c1d4dbc0a854d9f72157f1d6ea078", dir.Swhid().String())
	assert.Equal(t, "swh:1:cnt:e69de29bb2d1d6434b8b29ae775ad8c2e48c5391", cnt.Swhid().String())
}

func TestGitCommit(t *testing.T) {
	path, err := gitInit(t.TempDir())
	assert.Nil(t, err)
	repo, err := NewRepository(path)
	assert.Nil(t, err)
	_, err = repo.Head() // no head yeat
	assert.Error(t, err)
	worktree, err := repo.Repo.Worktree()
	assert.Nil(t, err)
	author := object.Signature{Name: "me", Email: "me@example.com"}
	opts := git.CommitOptions{Author: &author}
	hash, err := worktree.Commit("message", &opts)
	assert.Nil(t, err)
	head, err := repo.Head()
	assert.Nil(t, err)
	assert.Equal(t, hash.String(), head)
	rev, err := repo.NewRevisionFromHash(hash.String())
	assert.Nil(t, err)
	dir, err := repo.NewDirectoryFromHash(hash.String())
	assert.Nil(t, err)
	assert.Equal(t, "swh:1:rev:85377830ad661d517a1a23006f22907a37aa81be", rev.Swhid().String())
	assert.Equal(t, "swh:1:dir:4b825dc642cb6eb9a060e54bf8d69288fbee4904", dir.Swhid().String())
}

func TestGitTag(t *testing.T) {
	path, err := gitInit(t.TempDir())
	assert.Nil(t, err)
	repo, err := NewRepository(path)
	assert.Nil(t, err)
	_, err = repo.Tag("foo") // no tag yet
	assert.Error(t, err)
	worktree, err := repo.Repo.Worktree()
	assert.Nil(t, err)
	author := object.Signature{Name: "me", Email: "me@example.com"}
	opts := git.CommitOptions{Author: &author}
	hash, err := worktree.Commit("message", &opts)
	assert.Nil(t, err)
	tagger := object.Signature{Name: "me", Email: "me@example.com"}
	tagOpts := git.CreateTagOptions{Tagger: &tagger, Message: "message"}
	ref, err := repo.Repo.CreateTag("foo", hash, &tagOpts)
	assert.Nil(t, err)
	tag, err := repo.Tag("foo")
	assert.Nil(t, err)
	assert.Equal(t, ref.Hash().String(), tag)
	commit, err := repo.Commit("foo")
	assert.Nil(t, err)
	assert.Equal(t, hash.String(), commit)
	rel, err := repo.NewReleaseFromTag("foo")
	assert.Nil(t, err)
	rev, err := repo.NewRevisionFromTag("foo")
	assert.Nil(t, err)
	rel, err = repo.NewReleaseFromHash(ref.Hash().String())
	assert.Nil(t, err)
	assert.Equal(t, "swh:1:rel:79a98c5c621a2c52421cbaf40cbe05d358e041ac", rel.Swhid().String())
	assert.Equal(t, "swh:1:rev:85377830ad661d517a1a23006f22907a37aa81be", rev.Swhid().String())
}

func TestGitBranches(t *testing.T) {
	path, err := gitInit(t.TempDir())
	assert.Nil(t, err)
	repo, err := NewRepository(path)
	assert.Nil(t, err)
	err = os.Chmod(path+"/.git/refs/heads", 0o000)
	assert.Nil(t, err)
	_, err = repo.Branches() // access denied
	assert.Error(t, err)
	err = os.Chmod(path+"/.git/refs/heads", 0o775)
	assert.Nil(t, err)
	branchConfig := config.Branch{Name: "master", Merge: "refs/heads/master"}
	err = repo.Repo.CreateBranch(&branchConfig)
	worktree, err := repo.Repo.Worktree()
	assert.Nil(t, err)
	author := object.Signature{Name: "me", Email: "me@example.com"}
	opts := git.CommitOptions{Author: &author}
	hash, err := worktree.Commit("message", &opts)
	assert.Nil(t, err)
	branches, err := repo.Branches()
	assert.Nil(t, err)
	assert.Equal(t, []string{"master"}, branches)
	snp, err := repo.NewSnapshot()
	assert.Nil(t, err)
	assert.Equal(t, "swh:1:snp:e6db8e69084f013b6e40d250c9b25293a008c176", snp.Swhid().String())
	tagger := object.Signature{Name: "me", Email: "me@example.com"}
	tagOpts := git.CreateTagOptions{Tagger: &tagger, Message: "message"}
	_, err = repo.Repo.CreateTag("foo", hash, &tagOpts)
	assert.Nil(t, err)
	_, err = repo.Repo.CreateTag("bar", hash, nil)
	assert.Nil(t, err)
	snp, err = repo.NewSnapshot()
	assert.Equal(t, "swh:1:snp:e15a28657cf4b22c8e952611ab003437bf64fa67", snp.Swhid().String())
	assert.Nil(t, err)
}

func TestGitTags(t *testing.T) {
	path, err := gitInit(t.TempDir())
	assert.Nil(t, err)
	repo, err := NewRepository(path)
	assert.Nil(t, err)
	err = os.Chmod(path+"/.git/refs/tags", 0o000)
	assert.Nil(t, err)
	_, err = repo.Tags() // access denied
	assert.Error(t, err)
	err = os.Chmod(path+"/.git/refs/tags", 0o775)
	assert.Nil(t, err)
	ref1, err := repo.Repo.CreateTag("foo", [20]byte{}, nil)
	assert.Nil(t, err)
	hash1 := ref1.Hash().String()
	ref2, err := repo.Repo.CreateTag("bar", [20]byte{}, nil)
	assert.Nil(t, err)
	hash2 := ref2.Hash().String()
	tags, err := repo.Tags()
	assert.Nil(t, err)
	assert.Equal(t, []string{hash1, hash2}, tags)
}
