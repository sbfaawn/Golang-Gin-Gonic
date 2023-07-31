package response

type BaseResponse struct {
	Message string `json:"message" xml:"message" form:"message" query:"message"`
	Data    any    `json:"data" xml:"data" form:"data" query:"data"`
}
