package users

type Users struct {
	ID        int    `json:"id"`
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Age       int    `json:"Age"`
}

var users = []Users{
	{
		ID:        1,
		Firstname: "juheth",
		Lastname:  "fernando",
		Age:       17,
	},
}
