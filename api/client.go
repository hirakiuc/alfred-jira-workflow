package api

import (
	"errors"
	"os"
	"strings"

	"github.com/andygrunwald/go-jira"
)

type Client struct {
	jira   *jira.Client
	config *Config
}

type Config struct {
	APIToken string
	BaseURL  string
	Email    string
}

func loadConfig() (*Config, error) {
	token := os.Getenv("JIRA_API_TOKEN")
	if token == "" {
		return nil, errors.New("JIRA_API_TOKEN required")
	}

	baseURL := os.Getenv("JIRA_BASE_URL")
	if baseURL == "" {
		return nil, errors.New("JIRA_BASE_URL required")
	}

	email := os.Getenv("JIRA_EMAIL")
	if email == "" {
		return nil, errors.New("JIRA_EMAIL required")
	}

	return &Config{
		APIToken: token,
		BaseURL:  baseURL,
		Email:    email,
	}, nil
}

func NewClient() (*Client, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	tp := jira.BasicAuthTransport{
		Username: config.Email,
		Password: config.APIToken,
	}

	client, err := jira.NewClient(tp.Client(), config.BaseURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		jira:   client,
		config: config,
	}, nil
}

func (client *Client) BaseURL() string {
	baseURL := client.config.BaseURL

	if strings.HasPrefix(baseURL, "/") {
		return baseURL
	}

	return baseURL + "/"
}
