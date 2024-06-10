package git

import (
	"fmt"
	"os/exec"
)

var excludeFiles = []string{
	"*.lock",
	"*.lockb",
	"go.sum",
}

type Client struct{}

func New() *Client {
	return &Client{}
}

func (c *Client) DiffStaged() (string, error) {
	args := []string{
		"diff", "--cached", "--diff-algorithm=histogram",
	}
	for _, f := range excludeFiles {
		args = append(args, fmt.Sprintf(":(exclude)%s", f))
	}

	out, err := exec.Command("git", args...).Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func (c *Client) Commit(msg string) error {
	if err := exec.Command("git", "commit", "-m", msg).Run(); err != nil {
		return err
	}
	return nil
}
