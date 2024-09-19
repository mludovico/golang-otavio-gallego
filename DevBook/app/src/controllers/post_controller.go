package controllers

import (
	"devbook_app/src/cookies"
	httpclient "devbook_app/src/http_client"
	"devbook_app/src/response"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "Error reading request body")
		return
	}

	fmt.Printf("reqBody: %v\n", string(body))

	cookie, err := cookies.Read(r)
	if err != nil {
		response.JSON(w, http.StatusUnauthorized, nil)
		return
	}

	c := httpclient.NewClient(cookie["token"])

	resp, err := c.Post("/posts", string(body))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	defer r.Body.Close()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		response.JSON(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	response.JSON(w, http.StatusCreated, nil)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.Read(r)
	if err != nil {
		response.JSON(w, http.StatusUnauthorized, nil)
		return
	}

	c := httpclient.NewClient(cookie["token"])

	vars := mux.Vars(r)
	postID := vars["postID"]

	resp, err := c.Post(fmt.Sprintf("/posts/%s/like", postID), "")
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		response.JSON(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	response.JSON(w, http.StatusOK, nil)
}

func UnlikePost(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.Read(r)
	if err != nil {
		response.JSON(w, http.StatusUnauthorized, nil)
		return
	}

	c := httpclient.NewClient(cookie["token"])

	vars := mux.Vars(r)
	postID := vars["postID"]

	resp, err := c.Delete(fmt.Sprintf("/posts/%s/like", postID), "")
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		response.JSON(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	response.JSON(w, http.StatusOK, nil)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "Error reading request body")
		return
	}

	cookie, err := cookies.Read(r)
	if err != nil {
		response.JSON(w, http.StatusUnauthorized, nil)
		return
	}

	c := httpclient.NewClient(cookie["token"])

	vars := mux.Vars(r)
	postID := vars["postID"]

	resp, err := c.Put(fmt.Sprintf("/posts/%s", postID), string(body))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		response.JSON(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	response.JSON(w, http.StatusOK, nil)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.Read(r)
	if err != nil {
		response.JSON(w, http.StatusUnauthorized, nil)
		return
	}

	c := httpclient.NewClient(cookie["token"])

	vars := mux.Vars(r)
	postID := vars["postID"]

	resp, err := c.Delete(fmt.Sprintf("/posts/%s", postID), "")
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		response.JSON(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	response.JSON(w, http.StatusOK, nil)
}
