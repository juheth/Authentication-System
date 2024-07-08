package users

type Users struct {
	ID        int    `json:"id"`
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Email     string `json:"Email"`
	Age       int    `json:"Age"`
	Password  string `json:"Password"`
}

var users = []Users{
	{
		ID:        1,
		Firstname: "juheth",
		Lastname:  "fernando",
		Email:     "juheth@gmail.com",
		Age:       17,
		Password:  "juheth123",
	},
}
