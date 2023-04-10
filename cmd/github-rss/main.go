package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v51/github"
	"github.com/gorilla/feeds"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "%s: missing user parameter\n", os.Args[0])
		fmt.Printf("%s <GITHUB_USER>", os.Args[0])
		os.Exit(1)
	}
	user := os.Args[1]

	events, err := getEvents(user)
	if err != nil {
		log.Fatal(err)
	}

	feed := newFeed(user)
	for _, event := range events {
		items := parse(event)
		if items != nil {
			for _, item := range items {
				feed.Items = append(feed.Items, item)
			}
		}
	}
	feed.Created = feed.Items[0].Created
	atom, err := feed.ToAtom()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(atom)
}

func newFeed(login string) feeds.Feed {
	return feeds.Feed{
		Title:   login + " github activity",
		Link:    &feeds.Link{Href: "https://github.com/" + login},
		Created: time.Now(),
	}
}

func getEvents(login string) ([]*github.Event, error) {
	client := github.NewClient(nil)
	events, _, err := client.Activity.ListEventsReceivedByUser(
		context.Background(), login, true, nil)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func parse(event *github.Event) (items []*feeds.Item) {
	switch *event.Type {
	case "WatchEvent":
		item := feedWatch(event)
		if item != nil {
			items = append(items, item)
		}
	case "CreateEvent":
		item := feedCreate(event)
		if item != nil {
			items = append(items, item)
		}
	case "PushEvent":
		items = append(items, feedPush(event)...)
	}
	return items
}

func feedPush(e *github.Event) (items []*feeds.Item) {
	payload := e.Payload().(*github.PushEvent)
	for _, commit := range payload.Commits {
		item := &feeds.Item{
			Title: fmt.Sprintf("%s pushed to %s",
				*e.Actor.Login,
				*e.Repo.Name),
			Link: &feeds.Link{Href: fmt.Sprintf("https://github.com/%s/commit/%s",
				*e.Repo.Name,
				*commit.SHA)},
			Description: *commit.Message,
			Created:     e.CreatedAt.Time,
		}
		items = append(items, item)
	}
	return
}

func feedWatch(e *github.Event) *feeds.Item {
	payload := e.Payload().(*github.WatchEvent)
	if *payload.Action != "started" {
		return nil
	}
	return &feeds.Item{
		Title: fmt.Sprintf("⭐ %s starred %s",
			*e.Actor.Login,
			*e.Repo.Name),
		Link:    &feeds.Link{Href: "https://github.com/" + *e.Repo.Name},
		Created: e.CreatedAt.Time,
	}
}

func feedCreate(e *github.Event) *feeds.Item {
	payload := e.Payload().(*github.CreateEvent)
	if *payload.RefType != "repository" {
		return nil
	}
	return &feeds.Item{
		Title: fmt.Sprintf("%s created %s",
			*e.Actor.Login,
			*e.Repo.Name),
		Link:    &feeds.Link{Href: "https://github.com/" + *e.Repo.Name},
		Created: e.CreatedAt.Time,
	}
}
