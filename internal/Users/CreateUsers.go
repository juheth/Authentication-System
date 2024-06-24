package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	var newUser Users

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Please insert a valid user", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &newUser)
	if err != nil {
		http.Error(w, "Error parsing user data", http.StatusBadRequest)
		return
	}

	newUser.ID = len(users) + 1
	users = append(users, newUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(users)
}
