package entities

//Article Structure
type Article struct {
	Title string `json:"title" binding:"min=2,required"`
	Body  string `json:"post" binding:"min=2,required"`
	URL   string `json:"url" binding:"url"`
	ID    int    `json:"id" binding:"gte=1"`
}
