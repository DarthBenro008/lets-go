package newsfeed

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

type Deleter interface {
	Delete(item Item) bool
}

type Updater interface {
	Update(item Item) bool
}

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

func (r *Repo) Update(item Item) bool {
	for index, presentItem := range r.Items {
		if item.Title == presentItem.Title {
			r.Items[index] = item
			return true
		}
	}
	return false
}

func (r *Repo) Delete(item Item) bool {
	for index, presentItems := range r.Items {
		if item.Title == presentItems.Title {
			frontList := r.Items[:index]
			backList := r.Items[index+1:]
			r.Items = append(frontList, backList...)
			return true
		}
	}
	return false
}
