package request

type BookRequest struct {
	Title  string `json:"title" validate:"required,notblank" xml:"title" form:"title" query:"title"`
	Author string `json:"author" validate:"required,notblank" xml:"author" form:"author" query:"author"`
}
