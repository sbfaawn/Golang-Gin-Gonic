package request

type CredentialsRequest struct {
	Username string `json:"username" validate:"required,notblank" xml:"username" form:"username" query:"username"`
	Password string `json:"password" validate:"required,notblank" xml:"password" form:"password" query:"password"`
}
