package domain

type Category struct {
	Id   int
	Name string
}

func (c *Category) TableName() string {
	return "categories"
}
