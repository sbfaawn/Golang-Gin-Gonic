package authentication

import (
	"github.com/zalando/gin-oauth2/zalando"
)

var TeamsAccess []zalando.AccessTuple

func init() {
	TeamsAccess = []zalando.AccessTuple{
		{"teams", "8dw72d7g79dsds2", "all-teams"},
	}
}
