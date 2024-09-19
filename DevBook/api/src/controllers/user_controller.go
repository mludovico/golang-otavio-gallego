package controllers

import (
	"devbook_api/src/authentication"
	"devbook_api/src/db"
	"devbook_api/src/models"
	"devbook_api/src/repositories"
	"devbook_api/src/response"
	"devbook_api/src/security"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	user, err := repo.FindById(userId)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if user.ID == 0 {
		response.Error(w, http.StatusNotFound, errors.New("user not found"))
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	db, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	users, err := repo.FindAll(nameOrNick)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if len(users) == 0 {
		response.Error(w, http.StatusNotFound, errors.New("no users found"))
		return
	}

	response.JSON(w, http.StatusOK, users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	fmt.Printf("User before prepare: %v\n", user)
	if err = user.Prepare(models.InsertPreparation); err != nil {
		fmt.Printf("Error preparing user %v\n", err)
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	fmt.Printf("User after prepare: %v\n", user)

	db, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	fmt.Printf("User when calling repo: %v\n", user)
	userId, err := repo.Create(user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	user.ID = userId

	response.JSON(w, http.StatusCreated, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	userIdFromToken, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userId != userIdFromToken {
		response.Error(w, http.StatusForbidden, errors.New("you can only update your own user"))
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(models.UpdatePreparation); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	affectedRows, err := repo.Update(userId, user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if affectedRows == 0 {
		response.Error(w, http.StatusNotFound, errors.New("user not found"))
		return
	} else {
		response.JSON(w, http.StatusOK, affectedRows)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	userIdFromToken, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	fmt.Printf("userId: %d\nuserIdFromToken: %d\n", userId, userIdFromToken)

	if userId != userIdFromToken {
		response.Error(w, http.StatusForbidden, errors.New("you can only delete your own user"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	affectedRows, err := repo.Delete(userId)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if affectedRows == 0 {
		response.Error(w, http.StatusNotFound, errors.New("user not found"))
	} else {
		response.JSON(w, http.StatusOK, affectedRows)
	}
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	followedId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	fmt.Printf("followerId: %d\nfollowedId: %d\n", followerId, followedId)
	if followerId == followedId {
		response.Error(w, http.StatusForbidden, errors.New("you can't follow yourself"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	if err = repo.Follow(followerId, followedId); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	followedId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if followerId == followedId {
		response.Error(w, http.StatusForbidden, errors.New("you can't unfollow yourself"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	if err = repo.Unfollow(followerId, followedId); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func GetFollowers(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	followers, err := repo.FindFollowers(userId)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if len(followers) == 0 {
		response.JSON(w, http.StatusFound, []models.User{})
		return
	}

	response.JSON(w, http.StatusOK, followers)
}

func GetFollowing(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	following, err := repo.FindFollowing(userId)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if len(following) == 0 {
		response.JSON(w, http.StatusFound, []models.User{})
		return
	}

	response.JSON(w, http.StatusOK, following)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	requestedUserId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if userId != requestedUserId {
		response.Error(w, http.StatusForbidden, errors.New("you can only update your own password"))
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var password models.Password
	if err = json.Unmarshal(reqBody, &password); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = password.Prepare(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	currentStoredPassword, err := repo.GetCurrentUserPassword(userId)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(currentStoredPassword, password.Current); err != nil {
		response.Error(w, http.StatusUnauthorized, errors.New("current password does not match"))
		return
	}

	hashedPassword, err := security.Hash(password.New)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = repo.UpdatePassword(userId, string(hashedPassword)); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
