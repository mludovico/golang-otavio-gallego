package models

import (
	"devbook_app/src/cookies"
	httpclient "devbook_app/src/http_client"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Following []User    `json:"following,omitempty"`
	Followers []User    `json:"followers,omitempty"`
	Posts     []Post    `json:"posts,omitempty"`
}

func GetUserCompleteData(r *http.Request, userID int64) (User, error) {
	userChannel := make(chan User)
	followingChannel := make(chan []User)
	followersChannel := make(chan []User)
	postsChannel := make(chan []Post)

	go getUser(r, userChannel, userID)
	go getFollowing(r, followingChannel, userID)
	go getFollowers(r, followersChannel, userID)
	go getPosts(r, postsChannel, userID)

	var (
		user      User
		following []User
		followers []User
		posts     []Post
	)

	for i := 0; i < 4; i++ {
		select {
		case retrievedUser := <-userChannel:
			if retrievedUser.ID == 0 {
				return User{}, errors.New("error retrieving user")
			}
			user = retrievedUser
		case retrievedFollowing := <-followingChannel:
			if retrievedFollowing == nil {
				return User{}, errors.New("error retrieving following")
			}
			following = retrievedFollowing
		case retrievedFollower := <-followersChannel:
			if retrievedFollower == nil {
				return User{}, errors.New("error retrieving followers")
			}
			followers = retrievedFollower
		case retrievedPosts := <-postsChannel:
			if retrievedPosts == nil {
				return User{}, errors.New("error retrieving posts")
			}
			posts = retrievedPosts
		}
	}

	user.Following = following
	user.Followers = followers
	user.Posts = posts

	return user, nil
}

func getUser(r *http.Request, channel chan<- User, userID int64) {
	cookie, err := cookies.Read(r)
	if err != nil {
		channel <- User{}
	}

	c := httpclient.NewClient(cookie["token"])
	endpoint := fmt.Sprintf("/user/%d", userID)
	response, err := c.Get(endpoint)
	if err != nil {
		channel <- User{}
	}
	defer response.Body.Close()

	var user User
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		channel <- User{}
	}
	channel <- user
}

func getFollowing(r *http.Request, channel chan<- []User, userID int64) {
	cookie, err := cookies.Read(r)
	if err != nil {
		channel <- nil
	}

	c := httpclient.NewClient(cookie["token"])
	endpoint := fmt.Sprintf("/user/%d/following", userID)
	response, err := c.Get(endpoint)
	if err != nil {
		channel <- nil
	}
	defer response.Body.Close()

	following := []User{}
	if err := json.NewDecoder(response.Body).Decode(&following); err != nil {
		channel <- nil
	}
	channel <- following
}

func getFollowers(r *http.Request, channel chan<- []User, userID int64) {
	cookie, err := cookies.Read(r)
	if err != nil {
		channel <- nil
	}

	c := httpclient.NewClient(cookie["token"])
	endpoint := fmt.Sprintf("/user/%d/followers", userID)
	response, err := c.Get(endpoint)
	if err != nil {
		channel <- nil
	}
	defer response.Body.Close()

	followers := []User{}
	if err := json.NewDecoder(response.Body).Decode(&followers); err != nil {
		channel <- nil
	}
	channel <- followers
}

func getPosts(r *http.Request, channel chan<- []Post, userID int64) {
	cookie, err := cookies.Read(r)
	if err != nil {
		channel <- nil
	}

	c := httpclient.NewClient(cookie["token"])
	endpoint := fmt.Sprintf("/user/%d/posts", userID)
	response, err := c.Get(endpoint)
	if err != nil {
		channel <- nil
	}
	defer response.Body.Close()

	posts := []Post{}
	if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
		channel <- nil
	}
	channel <- posts
}
