package request

type ResetPasswordRequest struct {
	Username    string `json:"username" validate:"required,notblank" xml:"username" form:"username" query:"username"`
	NewPassword string `json:"newPassword" validate:"required,notblank,password" xml:"newPassword" form:"newPassword" query:"newPassword"`
}
