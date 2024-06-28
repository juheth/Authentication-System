package tasks

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "invalid ID")
		return
	}

	for i, t := range Tasks {
		if t.ID == taskID {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			fmt.Fprintf(w, "succesfully", taskID)
		}
	}
}
