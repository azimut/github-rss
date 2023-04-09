package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-github/github"
	"github.com/gorilla/feeds"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "%s: missing organization parameter\n", os.Args[0])
		fmt.Printf("%s <GITHUB_ORGANIZATION>", os.Args[0])
		os.Exit(1)
	}
	org := os.Args[1]

	repos, err := getRepos(org)
	if err != nil {
		log.Fatal(err)
	}

	feed := newRSSFeed(org)

	for _, repo := range repos {
		item := parse(repo)
		feed.Items = append(feed.Items, &item)
	}

	if len(feed.Items) > 0 {
		feed.Created = feed.Items[0].Created
	}

	atom, err := feed.ToAtom()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(atom)
}

func newRSSFeed(org string) feeds.Feed {
	return feeds.Feed{
		Title:   fmt.Sprintf("%s github org activity", org),
		Link:    &feeds.Link{Href: fmt.Sprintf("https://github.com/%s", org)},
		Created: time.Now(),
	}
}

func getRepos(org string) ([]*github.Repository, error) {
	client := github.NewClient(nil)
	opts := &github.RepositoryListByOrgOptions{
		Type: "public",
	}
	repos, _, err := client.Repositories.ListByOrg(
		context.Background(), org, opts)
	if err != nil {
		return nil, err
	}
	return repos, nil
}

func parse(repo *github.Repository) feeds.Item {
	language := ""
	if repo.Language != nil {
		language = *repo.Language
	}
	description := ""
	if repo.Description != nil {
		description = *repo.Description
	}
	return feeds.Item{
		Title: fmt.Sprintf("%s created %s (%s)",
			*repo.Owner.Login,
			*repo.Name,
			language),
		Link: &feeds.Link{Href: fmt.Sprintf("https://github.com/%s",
			*repo.FullName)},
		Description: description,
		Created:     repo.CreatedAt.Time,
	}
}
