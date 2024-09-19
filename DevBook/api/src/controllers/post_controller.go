package controllers

import (
	"devbook_api/src/authentication"
	"devbook_api/src/db"
	"devbook_api/src/models"
	"devbook_api/src/repositories"
	"devbook_api/src/response"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	fmt.Printf("reqBody: %v\n", string(reqBody))

	var post models.Post
	if err = json.Unmarshal(reqBody, &post); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorID = userID
	if err = post.Prepare(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostRepository(db)
	storedPost, err := postRepository.Create(post)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, storedPost)
}

func FindPost(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.ParseUint(mux.Vars(r)["postID"], 10, 64)
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

	postRepository := repositories.NewPostRepository(db)
	post, err := postRepository.FindPost(postID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if post.ID == 0 {
		fmt.Printf("Post not found. searched id:%v retrieved id: %v\n", postID, post.ID)
		fmt.Printf("Post: %v\n", post)
		response.JSON(w, http.StatusNotFound, nil)
		return
	}

	response.JSON(w, http.StatusOK, post)
}

func FindPosts(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostRepository(db)
	post, err := postRepository.FindPosts(userId)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	postID, err := strconv.ParseUint(mux.Vars(r)["postID"], 10, 64)
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

	postRepository := repositories.NewPostRepository(db)
	post, err := postRepository.FindPost(postID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if post.AuthorID != userID {
		response.Error(w, http.StatusForbidden, errors.New("you cannot update a post that is not yours"))
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var updatedPost models.Post
	if err = json.Unmarshal(reqBody, &updatedPost); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	post.Title = updatedPost.Title
	post.Description = updatedPost.Description

	if err = post.Prepare(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = postRepository.UpdatePost(int64(postID), post); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	postID, err := strconv.ParseUint(mux.Vars(r)["postID"], 10, 64)
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

	postRepository := repositories.NewPostRepository(db)
	post, err := postRepository.FindPost(postID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if post.AuthorID != userID {
		response.Error(w, http.StatusForbidden, errors.New("you cannot delete a post that is not yours"))
		return
	}

	if err = postRepository.DeletePost(int64(postID)); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func FindPostsByUser(w http.ResponseWriter, r *http.Request) {
	myID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	userId, err := strconv.ParseUint(mux.Vars(r)["userID"], 10, 64)
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

	postRepository := repositories.NewPostRepository(db)
	posts, err := postRepository.FindPostsByUser(userId, myID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, posts)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	postID, err := strconv.ParseUint(mux.Vars(r)["postID"], 10, 64)
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

	postRepository := repositories.NewPostRepository(db)
	if err = postRepository.LikePost(postID, userID); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func UnlikePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	postID, err := strconv.ParseUint(mux.Vars(r)["postID"], 10, 64)
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

	postRepository := repositories.NewPostRepository(db)
	if err = postRepository.UnlikePost(postID, userID); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
