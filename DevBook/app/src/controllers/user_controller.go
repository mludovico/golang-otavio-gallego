package controllers

import (
	"bytes"
	"devbook_app/src/config"
	"devbook_app/src/cookies"
	httpclient "devbook_app/src/http_client"
	"devbook_app/src/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	userUrl := fmt.Sprintf("%s/user", config.APIURL)
	response, err := http.Post(userUrl, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("Error creating user %v\n", err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	if response.StatusCode != http.StatusCreated {
		errorMessage := fmt.Sprintf("Error calling API. Status Code: %d\n", response.StatusCode)
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	fmt.Printf("Response: %v\n", response)

	w.WriteHeader(http.StatusOK)
	responseString := fmt.Sprintf("{\"status\": \"success\", \"response\":\"%v\"}", response)
	w.Write([]byte(responseString))
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	callFollow("follow", w, r)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	callFollow("unfollow", w, r)
}

func callFollow(resourcePath string, w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.Read(r)
	if err != nil {
		http.Error(w, "Error reading cookie", http.StatusUnauthorized)
		return
	}

	token := cookie["token"]

	userID, err := strconv.ParseInt(mux.Vars(r)["userID"], 10, 64)
	if err != nil {
		http.Error(w, "Error getting user ID", http.StatusBadRequest)
		return
	}

	userUrl := fmt.Sprintf("/user/%d/%s", userID, resourcePath)
	c := httpclient.NewClient(token)
	resp, err := c.Post(userUrl, "")
	if err != nil {
		message := fmt.Sprintf("Error %sing user", resourcePath)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusNoContent {
		message := fmt.Sprintf("Error calling API. Status Code: %d\n", resp.StatusCode)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("Response: %v\n", resp)

	message := fmt.Sprintf("User %sed successfully", resourcePath)
	response.JSON(w, http.StatusOK, message)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.Read(r)
	if err != nil {
		http.Error(w, "Error reading cookie", http.StatusUnauthorized)
		return
	}

	token := cookie["token"]
	ID := cookie["ID"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	userUrl := fmt.Sprintf("/user/%s", ID)
	c := httpclient.NewClient(token)
	resp, err := c.Put(userUrl, string(body))
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		message := fmt.Sprintf("Error calling API. Status Code: %d\n", resp.StatusCode)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	formData := make(map[string]string)
	if err = json.Unmarshal(body, &formData); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	updatedUserName := formData["name"]
	currentUserName := cookie["username"]

	if updatedUserName != currentUserName && updatedUserName != "" {
		cookie["username"] = updatedUserName
		cookies.Save(w, ID, token, updatedUserName)
	}

	response.JSON(w, http.StatusOK, "Profile updated successfully")
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.Read(r)
	if err != nil {
		http.Error(w, "Error reading cookie", http.StatusUnauthorized)
		return
	}

	token := cookie["token"]
	ID := cookie["ID"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	userUrl := fmt.Sprintf("/user/%s/updatePassword", ID)
	c := httpclient.NewClient(token)
	resp, err := c.Post(userUrl, string(body))
	if err != nil {
		http.Error(w, "Error updating password", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusNoContent {
		message := fmt.Sprintf("Error calling API. Status Code: %d\n", resp.StatusCode)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("Response: %v\n", resp)

	response.JSON(w, http.StatusOK, "Password updated successfully")
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.Read(r)
	if err != nil {
		http.Error(w, "Error reading cookie", http.StatusUnauthorized)
		return
	}

	token := cookie["token"]
	ID := cookie["ID"]

	userUrl := fmt.Sprintf("/user/%s", ID)
	c := httpclient.NewClient(token)
	resp, err := c.Delete(userUrl, "")
	if err != nil {
		http.Error(w, "Error deleting account", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		message := fmt.Sprintf("Error calling API. Status Code: %d\n", resp.StatusCode)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("Response: %v\n", resp)

	cookies.Clear(w)
	response.JSON(w, http.StatusOK, "Account deleted successfully")
}
