package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UsersID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "invalid ID")
		return
	}

	for i, r := range users {
		if r.ID == UsersID {
			if r.ID == UsersID {
				users = append(users[:i], users[i:1]...)
				println(w, "se elimino correctamente", UsersID)
			}
		}
	}
}
