package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var UpdateTask Taks

	if err != nil {
		fmt.Fprintf(w, "invalid ID")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "succesty")

	}
	json.Unmarshal(reqBody, &UpdateTask)

	for i, t := range Tasks {
		if t.ID == taskID {
			if t.ID == taskID {
				Tasks = append(Tasks[:i], Tasks[i+1:]...)
				UpdateTask.ID = taskID
				Tasks = append(Tasks, UpdateTask)
				fmt.Fprintf(w, "succesfully", taskID)
			}
		}
	}

}
