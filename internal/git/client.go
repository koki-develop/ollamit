package git

import (
	"fmt"
	"os/exec"
	"strings"
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

type Diff struct {
	Files   []string
	Content string
}

func (c *Client) DiffStaged() (*Diff, error) {
	excludes := []string{}
	for _, f := range excludeFiles {
		excludes = append(excludes, fmt.Sprintf(":(exclude)%s", f))
	}

	d := &Diff{}
	{
		args := []string{"diff", "--staged", "--name-only", "--diff-algorithm=histogram"}
		out, err := exec.Command("git", args...).Output()
		if err != nil {
			return nil, err
		}
		sout := strings.TrimSpace(string(out))
		if sout == "" {
			return nil, nil
		}
		d.Files = append(d.Files, strings.Split(sout, "\n")...)
	}

	{
		args := []string{"diff", "--staged", "--diff-algorithm=histogram"}
		out, err := exec.Command("git", append(args, excludes...)...).Output()
		if err != nil {
			return nil, err
		}
		d.Content = string(out)
	}

	return d, nil
}

func (c *Client) Commit(msg string) error {
	if err := exec.Command("git", "commit", "-m", msg).Run(); err != nil {
		return err
	}
	return nil
}
