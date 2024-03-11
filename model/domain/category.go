package domain

/*
	- struct disini adalah data model yang merupakan representasi dari table
*/
type Category struct {
	Id   int
	Name string
}

func (c *Category) TableName() string {
	return "categories"
}
