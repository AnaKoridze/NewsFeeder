package data

import "testing"

func TestRepo_Add(t *testing.T) {
	feed := NewRepo()

	feed.Add(Newsfeed{
		"bingo",
		"bango",
	})

	if (len(feed.Newsfeeds) == 0) {
		t.Errorf("not added")
	}
}

func TestRepo_GetAllNews(t *testing.T) {

}