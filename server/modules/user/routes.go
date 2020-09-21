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
	{
		Method:  http.MethodGet,
		Path:    "/user/{userID}",
		Name:    "GetUser",
		Handler: controllers.GetUser,
	},
	{
		Method:  http.MethodGet,
		Path:    "/users",
		Name:    "GetAllUsers",
		Handler: controllers.GetAllUsers,
	},
	{
		Method:  http.MethodDelete,
		Path:    "/users/{userID}",
		Name:    "DeleteUser",
		Handler: controllers.DeleteUser,
	},
}
