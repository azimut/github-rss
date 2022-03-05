package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/google/go-github/github"
	"github.com/gorilla/feeds"
)

var re = `https://github.com/(?P<owner>[a-zA-Z_\.-]+)/(?P<repo>[a-zA-Z_\.-]+)/issues/(?P<issue>\d+)`

type urlParts struct {
	url   string
	owner string
	repo  string
	issue int
}

func main() {
	parsedUrl, err := parseUrl(os.Args[1])
	if err != nil {
		panic(err)
	}
	comments, err := fetchComments(*parsedUrl)
	if err != nil {
		panic(err)
	}
	feed := newFeed(*parsedUrl)
	for _, comment := range comments {
		item := newItem(comment)
		if item != nil {
			feed.Items = append(feed.Items, item)
		}
	}
	atom, err := feed.ToAtom()
	if err != nil {
		panic(err)
	}
	fmt.Println(atom)
}

func parseUrl(url string) (*urlParts, error) {
	var r = regexp.MustCompile(re)
	match := r.FindStringSubmatch(url)
	issue, err := strconv.Atoi(match[3])
	if err != nil {
		return nil, err
	}
	return &urlParts{
		url:   url,
		owner: match[1],
		repo:  match[2],
		issue: issue,
	}, nil
}

func fetchComments(u urlParts) ([]*github.IssueComment, error) {
	client := github.NewClient(nil)
	comments, _, err := client.Issues.ListComments(
		context.Background(),
		u.owner,
		u.repo,
		u.issue,
		nil)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func newFeed(u urlParts) *feeds.Feed {
	return &feeds.Feed{
		Title:   fmt.Sprintf("Issue activity %s/%s/%d", u.owner, u.repo, u.issue),
		Link:    &feeds.Link{Href: u.url},
		Created: time.Now(),
	}
}

func newItem(repo *github.IssueComment) *feeds.Item {
	return &feeds.Item{
		Title:       *repo.User.Login + " added a comment",
		Link:        &feeds.Link{Href: *repo.HTMLURL},
		Description: *repo.Body,
		Created:     *repo.CreatedAt,
	}
}
