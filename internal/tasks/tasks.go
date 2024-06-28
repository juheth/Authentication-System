package tasks

type Taks struct {
	ID      int    `json:"ID"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type AllTaks []Taks

var Tasks = AllTaks{
	{
		ID:      1,
		Name:    "task the juheth",
		Content: "movie",
	},
}
