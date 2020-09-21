package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/PrathyushaLakkireddy/go-rest-api-example/server/db"
	"github.com/PrathyushaLakkireddy/go-rest-api-example/server/modules/user/helpers"
	"github.com/PrathyushaLakkireddy/go-rest-api-example/server/modules/user/models"
	"github.com/PrathyushaLakkireddy/go-rest-api-example/server/utils"
)

// CreateUser creates a new user, it reads json data as an input and stores in db
func CreateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var user models.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		msg := "Error while reading input body"

		utils.ReturnErrorResponse(http.StatusBadRequest, msg, "", nil, nil, res)
		return
	}

	// Helper function to generate encrypted password hash
	passwordHash := helpers.GeneratePasswordHash(user.Password)

	if passwordHash == "" {
		msg := "Error occurred while hashing the password"

		utils.ReturnErrorResponse(http.StatusBadRequest, msg, "", nil, nil, res)
		return
	}

	user.ID = bson.NewObjectId()
	user.Password = passwordHash

	err := db.CreateUser(user)
	if err != nil {
		msg := "Error occurred while creating user"

		utils.ReturnErrorResponse(http.StatusBadRequest, msg, "", nil, nil, res)
		return
	}

	msg := "User created successfully"
	utils.ReturnSuccessReponse(http.StatusCreated, msg, user.ID, res)

}

// UpdateUser updates a particular user details by taing userID as query params
// and user data in the format of json
func UpdateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	userID := bson.ObjectIdHex(params["userID"])

	var user models.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		msg := "Error while reading input body"

		utils.ReturnErrorResponse(http.StatusBadRequest, msg, "", nil, nil, res)
		return
	}

	// Helper function to generate encrypted password hash
	passwordHash := helpers.GeneratePasswordHash(user.Password)

	if passwordHash == "" {
		msg := "Error occurred while hashing the password"

		utils.ReturnErrorResponse(http.StatusBadRequest, msg, "", nil, nil, res)
		return
	}

	query := bson.M{
		"_id": userID,
	}

	update := bson.M{
		"$set": user,
	}

	err := db.UpdateUser(query, update)
	if err != nil {
		if err.Error() == mgo.ErrNotFound.Error() {
			msg := "User not found"

			utils.ReturnErrorResponse(http.StatusNotFound, msg, "", nil, nil, res)
			return
		}

		msg := "Error occurred while updating the user"

		utils.ReturnErrorResponse(http.StatusBadRequest, msg, "", nil, nil, res)
		return
	}

	msg := "User updated successfully"
	utils.ReturnSuccessReponse(http.StatusOK, msg, userID, res)

}

// GetUser returns user single user details
// it reads userID as query params and returns data accordingly
func GetUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	userID := bson.ObjectIdHex(params["userID"])

	query := bson.M{
		"_id": userID,
	}

	selector := bson.M{
		"_id":   1,
		"name":  1,
		"email": 1,
	}

	user, err := db.GetUser(query, selector)
	if err != nil {
		if err.Error() == mgo.ErrNotFound.Error() {
			msg := "User not found"

			utils.ReturnErrorResponse(http.StatusNotFound, msg, "", nil, nil, res)
			return
		}

		msg := "Error occurred while getting user details"

		utils.ReturnErrorResponse(http.StatusBadRequest, msg, "", nil, nil, res)
		return
	}

	msg := "Your request processed successfully"
	utils.ReturnSuccessReponse(http.StatusOK, msg, user, res)

}

// GetAllUsers returns list of users from the database
func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	query := bson.M{}
	selector := bson.M{
		"_id":   1,
		"name":  1,
		"email": 1,
	}

	users, err := db.GetAllUsers(query, selector)
	if err != nil {
		if err.Error() == mgo.ErrNotFound.Error() {
			msg := "User not found"

			utils.ReturnErrorResponse(http.StatusNotFound, msg, "", nil, nil, res)
			return
		}

		msg := "Error occurred while getting user details"
		utils.ReturnErrorResponse(http.StatusBadRequest, msg, "", nil, nil, res)
		return
	}

	msg := "Your request processed successfully"
	utils.ReturnSuccessReponse(http.StatusOK, msg, users, res)
}

// DeleteUser will deletes the user details
// it takes userID from query params and deletes that particular user details
func DeleteUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	userID := bson.ObjectIdHex(params["userID"])

	query := bson.M{
		"_id": userID,
	}

	err := db.DeleteUser(query)
	if err != nil {
		if err.Error() == mgo.ErrNotFound.Error() {
			msg := "User not found"

			utils.ReturnErrorResponse(http.StatusNotFound, msg, "", nil, nil, res)
			return
		}

		msg := "Error occurred while deleting user"
		utils.ReturnErrorResponse(http.StatusBadRequest, msg, "", nil, nil, res)
		return
	}

	msg := "User deleted successfully"
	utils.ReturnSuccessReponse(http.StatusOK, msg, nil, res)
}
