package controllers

import (
	"devbook_api/src/authentication"
	"devbook_api/src/db"
	"devbook_api/src/models"
	"devbook_api/src/repositories"
	"devbook_api/src/response"
	"devbook_api/src/security"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
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
	storedUser, err := repo.FindByEmail(user.Email)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(storedUser.Password, user.Password); err != nil {
		fmt.Printf("Error verifying password %v\n", err)
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(storedUser.ID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	userID := strconv.FormatUint(storedUser.ID, 10)

	authData := models.AuthData{
		ID:       userID,
		Token:    token,
		Username: storedUser.Name,
	}

	response.JSON(w, http.StatusOK, authData)
}
