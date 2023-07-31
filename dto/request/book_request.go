package request

type BookRequest struct {
	Name   string `json:"name" xml:"name" form:"name" query:"name"`
	Author string `json:"author" xml:"author" form:"author" query:"author"`
}
