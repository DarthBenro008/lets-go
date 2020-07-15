package newsfeed

type Item struct {
	Title string `json: "title`
	Post  string `json: "post"`
	Url   string `json: "url`
}

type Repo struct {
	Items []Item
}

func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	return r.Items
}

func (r *Repo) Update(item Item) {

}
