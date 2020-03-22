package data

type Newsfeed struct {
	Title string `json:"title"`
	Post string `json:"post"`
}

type Repo struct {
	Newsfeeds []Newsfeed
}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) Add(News Newsfeed) {
	r.Newsfeeds = append(r.Newsfeeds, News)
}

func (r *Repo) GetAllNews() []Newsfeed {
	return r.Newsfeeds
}