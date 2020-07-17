package entities

type book struct {
	ID     string `json:"id" binding:"gte=1,required" bson:"_id"`
	Title  string `json:"title" binding:"required,min=2" bson:"title"`
	Author string `json:"author" bson:"author,omitempty"`
}
