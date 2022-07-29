package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-github/github"
	"github.com/gorilla/feeds"
)

var re = `https://github.com/(?P<owner>[a-zA-Z_\.-]+)/(?P<repo>[a-zA-Z_\.-]+)/issues/(?P<issue>\d+)`

type UrlParts struct {
	url   string
	owner string
	repo  string
	issue int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "%s: missing url parameter\n", os.Args[0])
		fmt.Printf("%s <GITHUB_ISSUE_URL>", os.Args[0])
		os.Exit(1)
	}
	urlParts, err := newUrlParts(os.Args[1])
	if err != nil {
		panic(err)
	}
	comments, err := fetchComments(*urlParts)
	if err != nil {
		panic(err)
	}
	feed := newFeed(*urlParts)
	for _, comment := range comments {
		item := newItem(comment)
		if item != nil {
			feed.Items = append(feed.Items, item)
		}
	}
	atom, err := feed.ToRss()
	if err != nil {
		panic(err)
	}
	fmt.Println(atom)
}

func newUrlParts(url string) (*UrlParts, error) {
	var r = regexp.MustCompile(re)
	match := r.FindStringSubmatch(url)
	issue, err := strconv.Atoi(match[3])
	if err != nil {
		return nil, err
	}
	return &UrlParts{
		url:   url,
		owner: match[1],
		repo:  match[2],
		issue: issue,
	}, nil
}

func fetchComments(u UrlParts) ([]*github.IssueComment, error) {
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

func newFeed(u UrlParts) *feeds.Feed {
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
		Description: strings.ReplaceAll(repo.GetBody(), "\n", "<br />"),
		Created:     *repo.CreatedAt,
	}
}
