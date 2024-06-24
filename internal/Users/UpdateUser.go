package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UsersID, err := strconv.Atoi(vars["id"])

	var UpdateUser Users

	if err != nil {
		fmt.Fprintf(w, "invalid id")
	}
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "succesty")
	}
	json.Unmarshal(reqBody, &UpdateUser)

	for i, r := range users {
		if r.ID == UsersID {
			if r.ID == UsersID {
				users = append(users[:i], users[i:1]...)
				UpdateUser.ID = UsersID
				users = append(users, UpdateUser)
				println(w, "se actualizo correctamente", UsersID)
			}
		}
	}

}
