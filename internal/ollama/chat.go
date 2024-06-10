package ollama

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatInput struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatOutputChunk struct {
	Message *ChatMessage `json:"message"`
}

type ChatOutput struct {
	scanner *bufio.Scanner
	body    io.ReadCloser
}

func (c *ChatOutput) Close() error {
	return c.body.Close()
}

func (c *ChatOutput) Scan() bool {
	return c.scanner.Scan()
}

func (c *ChatOutput) Chunk() (*ChatOutputChunk, error) {
	var chunk ChatOutputChunk
	if err := json.Unmarshal(c.scanner.Bytes(), &chunk); err != nil {
		return nil, err
	}
	return &chunk, nil
}

func (c *Client) Chat(ipt *ChatInput) (*ChatOutput, error) {
	e := "http://localhost:11434/api/chat"

	p, err := json.Marshal(ipt)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, e, bytes.NewReader(p))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	hc := new(http.Client)
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		buf := new(bytes.Buffer)
		if _, err := io.Copy(buf, resp.Body); err != nil {
			return nil, err
		}
		return nil, errors.New(buf.String())
	}

	return &ChatOutput{
		scanner: bufio.NewScanner(resp.Body),
		body:    resp.Body,
	}, nil
}
