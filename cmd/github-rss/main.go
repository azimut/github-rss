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
	login := os.Args[1]

	events, err := getEvents(login)
	if err != nil {
		log.Fatal(err)
	}

	feed := newFeed(login)
	for _, event := range events {
		nnn
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

func newFeed(login string) *feeds.Feed {
	now := time.Now()
	feed := &feeds.Feed{
		Title:   fmt.Sprintf("%s github activity", login),
		Link:    &feeds.Link{Href: fmt.Sprintf("https://github.com/%s", login)},
		Created: now,
	}
	return feed
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
		items = append(items, feedWatch(event))
	case "CreateEvent":
		items = append(items, feedCreate(event))
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
			Created:     *e.CreatedAt,
		}
		items = append(items, item)
	}
	return items
}

func feedWatch(e *github.Event) *feeds.Item {
	payload := e.Payload().(*github.WatchEvent)
	if *payload.Action != "started" {
		return nil
	}
	return &feeds.Item{
		Title: fmt.Sprintf("%s starred %s",
			*e.Actor.Login,
			*e.Repo.Name),
		Link: &feeds.Link{Href: fmt.Sprintf("https://github.com/%s",
			*e.Repo.Name)},
		Created: *e.CreatedAt,
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
		Link: &feeds.Link{Href: fmt.Sprintf("https://github.com/%s",
			*e.Repo.Name)},
		Created: *e.CreatedAt,
	}
}
