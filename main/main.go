package main

import (
	"NewsFeeder/data"
	"fmt"
)

func main()  {
	feed := data.NewRepo()

	feed.Add(data.Newsfeed{
		"bingo",
		"bango",
	})

	fmt.Println(*feed)
}

