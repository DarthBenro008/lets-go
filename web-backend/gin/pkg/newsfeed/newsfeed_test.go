package newsfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{"Hello", "This is test", "Test URL"})
	if len(feed.Items) != 1 {
		t.Error("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	if len(feed.GetAll()) != 1 {
		t.Error("Did not fetch all newsfeeds")
	}
}
