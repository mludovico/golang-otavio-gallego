package controllers

import (
	"devbook_app/src/cookies"
	httpclient "devbook_app/src/http_client"
	"devbook_app/src/models"
	"devbook_app/src/response"
	"devbook_app/src/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if cookies.IsValid(r) {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}
	utils.RespondWithTemplate(w, "login.html", nil)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithTemplate(w, "signup.html", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.Read(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	token := cookie["token"]
	username := cookie["username"]

	userID, err := strconv.ParseInt(cookie["ID"], 10, 64)
	if err != nil {
		userID = 0
	}

	c := httpclient.NewClient(token)

	var posts []models.Post
	response, err := c.Get("/posts")
	if err != nil {
		posts = []models.Post{}
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
		posts = []models.Post{}
	}

	utils.RespondWithTemplate(w, "home.html", struct {
		EditingPost  models.Post
		LoggedUserID int64
		Posts        []models.Post
		Header       string
	}{
		EditingPost:  models.Post{},
		LoggedUserID: userID,
		Posts:        posts,
		Header:       username,
	})
}

func LoadEditPostView(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.Read(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	token := cookie["token"]
	username := cookie["username"]

	userID, err := strconv.ParseInt(cookie["ID"], 10, 64)
	if err != nil {
		userID = 0
	}

	c := httpclient.NewClient(token)

	postID := mux.Vars(r)["postID"]

	var posts []models.Post
	response, err := c.Get("/posts")
	if err != nil {
		posts = []models.Post{}
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
		posts = []models.Post{}
	}

	var editingPost models.Post
	for _, post := range posts {
		if strconv.FormatInt(int64(post.ID), 10) == postID {
			editingPost = post
			break
		}
	}

	utils.RespondWithTemplate(w, "home.html", struct {
		EditingPost  models.Post
		LoggedUserID int64
		Posts        []models.Post
		Header       string
	}{
		EditingPost:  editingPost,
		LoggedUserID: userID,
		Posts:        posts,
		Header:       username,
	})
}

func LoadCreatePostView(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.Read(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	token := cookie["token"]
	username := cookie["username"]

	userID, err := strconv.ParseInt(cookie["ID"], 10, 64)
	if err != nil {
		userID = 0
	}

	c := httpclient.NewClient(token)

	response, err := c.Get("/posts")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	var posts []models.Post
	if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	utils.RespondWithTemplate(w, "home.html", struct {
		EditingPost  models.Post
		LoggedUserID int64
		Posts        []models.Post
		Header       string
	}{
		EditingPost: models.Post{
			AuthorID: uint64(userID),
		},
		LoggedUserID: userID,
		Posts:        posts,
		Header:       username,
	})
}
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	cookie, err := cookies.Read(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	token := cookie["token"]
	myID, err := strconv.ParseInt(cookie["ID"], 10, 64)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	c := httpclient.NewClient(token)

	endpoint := "/user?" + queryParams.Encode()

	response, err := c.Get(endpoint)
	if err != nil {
		if response.StatusCode == http.StatusUnauthorized {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}
	defer response.Body.Close()

	var users []models.User

	if response.StatusCode == http.StatusNotFound {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}
		if !strings.Contains(string(body), "no users found") {
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}
	}

	if response.StatusCode == http.StatusOK {
		if err := json.NewDecoder(response.Body).Decode(&users); err != nil {
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}
	}

	utils.RespondWithTemplate(w, "users_search.html", struct {
		MyID  int64
		Users []models.User
	}{
		MyID:  myID,
		Users: users,
	})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.Read(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	MyID, err := strconv.ParseInt(cookie["ID"], 10, 64)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	vars := mux.Vars(r)
	userID, err := strconv.ParseInt(vars["userID"], 10, 64)
	if err != nil {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	if MyID == userID {
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	user, err := models.GetUserCompleteData(r, userID)
	if err != nil {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	utils.RespondWithTemplate(w, "users.html", struct {
		MyID uint64
		User models.User
	}{
		MyID: uint64(MyID),
		User: user,
	})
}

func ListFollowers(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.Read(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	vars := mux.Vars(r)
	userID, err := strconv.ParseInt(vars["userID"], 10, 64)
	if err != nil {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	token := cookie["token"]
	myID, err := strconv.ParseInt(cookie["ID"], 10, 64)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	endpoint := fmt.Sprintf("/user/%d/followers", userID)
	c := httpclient.NewClient(token)
	resp, err := c.Get(endpoint)
	if err != nil {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	var followers []models.User
	if err := json.NewDecoder(resp.Body).Decode(&followers); err != nil {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	utils.RespondWithTemplate(w, "users_search.html", struct {
		MyID  int64
		Users []models.User
	}{
		MyID:  int64(myID),
		Users: followers,
	})
}

func ListFollowing(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.Read(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	vars := mux.Vars(r)
	userID, err := strconv.ParseInt(vars["userID"], 10, 64)
	if err != nil {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	token := cookie["token"]
	myID, err := strconv.ParseInt(cookie["ID"], 10, 64)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	endpoint := fmt.Sprintf("/user/%d/following", userID)
	c := httpclient.NewClient(token)
	resp, err := c.Get(endpoint)
	if err != nil {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	var following []models.User
	if err := json.NewDecoder(resp.Body).Decode(&following); err != nil {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	utils.RespondWithTemplate(w, "users_search.html", struct {
		MyID  int64
		Users []models.User
	}{
		MyID:  myID,
		Users: following,
	})
}

func Profile(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.Read(r)
	if err != nil {
		response.JSON(w, http.StatusUnauthorized, "Error reading cookie")
		return
	}

	userID, err := strconv.ParseInt(cookie["ID"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorResponse{Message: "Error getting user ID"})
		return
	}

	user, err := models.GetUserCompleteData(r, userID)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorResponse{Message: "Error getting user data"})
		return
	}

	utils.RespondWithTemplate(w, "profile.html", struct {
		User models.User
	}{
		User: user,
	})
}

func EditProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.Read(r)
	if err != nil {
		response.JSON(w, http.StatusUnauthorized, "Error reading cookie")
		return
	}

	userID, err := strconv.ParseInt(cookie["ID"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorResponse{Message: "Error getting user ID"})
		return
	}

	user, err := models.GetUserCompleteData(r, userID)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorResponse{Message: "Error getting user data"})
		return
	}

	utils.RespondWithTemplate(w, "edit_profile.html", struct {
		User models.User
	}{
		User: user,
	})
}

func EditPassword(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithTemplate(w, "edit_password.html", nil)
}
