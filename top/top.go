package top

import (
	"errors"
	"fmt"

	"mau/github"
	"mau/net"
)

func GithubTop(options Options) ([]github.User, error) {
	var token = options.Token
	if token == "" {
		return []github.User{}, errors.New("Missing GITHUB token")
	}

	query := "repos:>1 type:user"
	for _, location := range options.Locations {
		query = fmt.Sprintf("%s location:%s", query, location)
	}

	var client = github.NewGithubClient(net.TokenAuth(token))
	users, err := client.SearchUsers(github.UserSearchQuery{Q: query, Sort: "followers", Order: "desc", MaxUsers: options.ConsiderNum})
	if err != nil {
		return []github.User{}, err
	}
	return users, nil
}

type Options struct {
	Token       string
	Locations   []string
	Amount      int
	ConsiderNum int
}
