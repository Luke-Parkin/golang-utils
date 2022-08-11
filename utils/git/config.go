package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/ARM-software/golang-utils/utils/config"
)

// GitActionConfig describes how a clone or checkout should be performed.
type GitActionConfig struct {
	// The (possibly remote) repository URL to clone from.
	URL string `mapstructure:"url"`
	// Auth credentials, if required, to use with the remote repository.
	Auth transport.AuthMethod
	// Limit fetching to the specified number of commits.
	Depth int `mapstructure:"depth"`
	// Hash is the hash of the commit to be checked out. If used, HEAD will be
	// in detached mode. If Create is not used, Branch and Hash are mutually
	// exclusive.
	Hash string `mapstructure:"hash"`
	// Branch to be checked out, if Branch and Hash are empty is set to `master`.
	Branch string `mapstructure:"branch"`
	// RecurseSubmodules after the clone is created, initialize all submodules
	// within, using their default settings. This option is ignored if the
	// cloned repository does not have a worktree.
	RecurseSubmodules bool `mapstructure:"recursive_submodules"`
	// Tags describe how the tags will be fetched from the remote repository,
	// by default is AllTags.
	Tags git.TagMode `mapstructure:"tags"`
	// CreateBranch a new branch named Branch and start it at Hash.
	CreateBranch bool `mapstructure:"create_branch"`
	// No checkout of HEAD after clone if true.
	NoCheckout bool
}

func (c *GitActionConfig) GetURL() string {
	return c.URL
}

func (c *GitActionConfig) GetAuth() transport.AuthMethod {
	return c.Auth
}

func (c *GitActionConfig) GetDepth() int {
	return c.Depth
}

func (c *GitActionConfig) GetHash() string {
	return c.Hash
}

func (c *GitActionConfig) GetBranch() string {
	return c.Branch
}

func (c *GitActionConfig) GetRecursiveSubModules() bool {
	return c.RecurseSubmodules
}

func (c *GitActionConfig) GetTags() git.TagMode {
	return c.Tags
}

func (c *GitActionConfig) GetCreate() bool {
	return c.CreateBranch
}

func (c *GitActionConfig) GetNoCheckout() bool {
	return c.NoCheckout
}

func (c *GitActionConfig) Validate() error {
	validation.ErrorTag = "mapstructure"

	// Validate Embedded Structs
	err := config.ValidateEmbedded(c)
	if err != nil {
		return err
	}
	return validation.ValidateStruct(c,
		validation.Field(&c.URL, validation.Required),
	)
}

func NewGitActionConfig(url string) GitActionConfig {
	return GitActionConfig{
		URL: url,
	}
}
