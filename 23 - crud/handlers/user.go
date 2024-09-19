package handlers

import (
	"crud/db"
	"crud/entities"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating a new user")
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Error reading the request body"))
		return
	}

	var user entities.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
		w.Write([]byte("Error parsing the request body"))
		return
	}

	fmt.Println(user)

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to the database"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("insert into user (name, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Error preparing the statement"))
		return
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error executing the statement"))
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		w.Write([]byte("Error getting the last insert ID"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Created user %s with ID %d", user.Name, id)))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all users")
	db, err := db.Connect()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error connecting to the database"))
		return
	}

	defer db.Close()

	rows, err := db.Query("select * from user")
	if err != nil {
		w.Write([]byte("Error querying the database"))
		return
	}

	var users []entities.User
	for rows.Next() {
		var user entities.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error scanning the result"))
			return
		}

		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Error encoding the result"))
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting a user")
	userId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user ID"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to the database"))
		return
	}

	defer db.Close()

	rows, err := db.Query("select * from user where id = ?", userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error preparing the statement"))
		return
	}

	var user entities.User
	if rows.Next() {
		if err = rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error scanning the result"))
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(user); err != nil {
				w.Write([]byte("Error encoding the result"))
				return
			}

		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating a user")
	userId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user ID"))
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Error reading the request body"))
		return
	}

	var user entities.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
		w.Write([]byte("Error parsing the request body"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to the database"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("update user set name = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Error preparing the statement"))
		return
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email, userId)
	if err != nil {
		w.Write([]byte("Error executing the statement"))
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		w.Write([]byte("Error getting the rows affected"))
		return
	}

	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting a user")
	userId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user ID"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to the database"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("delete from user where id = ?")
	if err != nil {
		w.Write([]byte("Error preparing the statement"))
		return
	}

	defer statement.Close()

	result, err := statement.Exec(userId)
	if err != nil {
		w.Write([]byte("Error executing the statement"))
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		w.Write([]byte("Error getting the rows affected"))
		return
	}

	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
