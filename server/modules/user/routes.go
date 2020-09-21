package user

import (
	"net/http"

	"github.com/PrathyushaLakkireddy/go-rest-api-example/server/common"
	"github.com/PrathyushaLakkireddy/go-rest-api-example/server/modules/user/controllers"
)

var Routes = []common.Route{
	{
		Method:  http.MethodPost,
		Path:    "/user",
		Name:    "Createuser",
		Handler: controllers.CreateUser,
	},
	{
		Method:  http.MethodPut,
		Path:    "/user/{userID}",
		Name:    "UpdateUser",
		Handler: controllers.UpdateUser,
	},
}
