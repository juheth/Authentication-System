package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateTasks(w http.ResponseWriter, r *http.Request) {
	var newTask Taks
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "insert a valid task")
	}
	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(Tasks) + 1
	Tasks = append(Tasks, newTask)

	w.Header().Set("contet-type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(newTask)
}
