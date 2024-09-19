package controllers

import (
	"bytes"
	"devbook_app/src/config"
	"devbook_app/src/cookies"
	"devbook_app/src/models"
	resp "devbook_app/src/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	loginUrl := fmt.Sprintf("%s/login", config.APIURL)
	fmt.Printf("Calling API %s\n", loginUrl)
	response, err := http.Post(loginUrl, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("Error logging in %v\n", err)
		http.Error(w, "Error logging in.", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		errorMessage := fmt.Sprintf("Error calling API. Status code: %d\n", response.StatusCode)
		http.Error(w, errorMessage, response.StatusCode)
		return
	}

	var authData models.AuthData
	if err = json.NewDecoder(response.Body).Decode(&authData); err != nil {
		resp.JSON(w, http.StatusUnprocessableEntity, resp.ErrorResponse{Message: err.Error()})
		return
	}

	if err = cookies.Save(w, authData.ID, authData.Token, authData.Username); err != nil {
		resp.JSON(w, http.StatusUnprocessableEntity, resp.ErrorResponse{Message: err.Error()})
		return
	}

	resp.JSON(w, http.StatusOK, nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Clear(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
