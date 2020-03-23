package data

import "testing"

func TestRepo_Add(t *testing.T) {
	feed := NewRepo()

	feed.Add(Newsfeed{
		"bingo",
		"bango",
	})

	if (len(feed.Newsfeeds) != 1) {
		t.Errorf("not added")
	}
}

func TestRepo_GetAllNews(t *testing.T) {
	feed := NewRepo()

	feed.Add(Newsfeed{
		"bingo",
		"bango",
	})

	results := feed.GetAllNews()

	if (len(results) != 1) {
		t.Errorf("get not working")
	}
}