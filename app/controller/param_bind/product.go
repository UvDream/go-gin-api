package parambind

type ProductAdd struct {
	Name string `form:"name" json:"name" validate:"required,NameValid"`
}
