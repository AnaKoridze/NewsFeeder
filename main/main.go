package main

import (
	"github.com/AnaKoridze/NewsFeeder/app"
	_ "github.com/lib/pq"
)

func main() {

	app.New().Run()
}
