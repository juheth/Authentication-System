package register

import (
	"database/sql"
	"html/template"
	"net/http"
)

var db *sql.DB

// SetDB sets the database connection
func SetDB(database *sql.DB) {
	db = database
}

// RegisterFormHandler handles the registration form rendering
func RegisterFormHandler(w http.ResponseWriter, r *http.Request) {
	// Get the username and lastname from the query parameters
	username := r.URL.Query().Get("username")
	lastname := r.URL.Query().Get("lastname")

	tpl := `<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Registro de Usuario</title>
    <style>
        body {
            font-family: 'Roboto', sans-serif;
            background: #f0f0f0;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
            color: #333;
        }

        .container {
            background-color: #fff;
            padding: 40px;
            border-radius: 10px;
            box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
            max-width: 400px;
            width: 100%;
            text-align: center;
        }

        h2 {
            color: #333;
            font-weight: 700;
            font-size: 24px;
            margin-bottom: 20px;
        }

        form {
            text-align: left;
        }

        label {
            color: #555;
            font-weight: 600;
            margin-bottom: 8px;
            display: block;
        }

        input[type="text"], input[type="submit"] {
            width: 100%;
            padding: 12px;
            margin-bottom: 20px;
            border: 1px solid #ccc;
            border-radius: 4px;
            box-sizing: border-box;
            font-size: 16px;
        }

        input[type="submit"] {
            background-color: #007BFF;
            color: #fff;
            border: none;
            cursor: pointer;
            transition: background-color 0.3s;
        }

        input[type="submit"]:hover {
            background-color: #0056b3;
        }
    </style>
</head>
<body>
    <div class="container">
        <h2>Registro de Usuario</h2>
        <form action="/register" method="post">
            <label for="username">Username:</label><br>
            <input type="text" id="username" name="username" value="{{.Username}}" required><br><br>
            <label for="lastname">Lastname:</label><br>
            <input type="text" id="lastname" name="lastname" value="{{.Lastname}}" required><br><br>
            <input type="submit" value="Registrar">
        </form>
    </div>
</body>
</html>`

	// Create and parse the template
	t, err := template.New("register-form").Parse(tpl)
	if err != nil {
		http.Error(w, "Error al renderizar el formulario", http.StatusInternalServerError)
		return
	}

	// Passing data to the template
	data := struct {
		Username string
		Lastname string
	}{
		Username: username,
		Lastname: lastname,
	}

	// Execute and send the result to the ResponseWriter
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Error al renderizar el formulario", http.StatusInternalServerError)
		return
	}
}

// RegisterHandler handles user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error al analizar los datos del formulario", http.StatusBadRequest)
		return
	}

	// Get values from the form
	username := r.Form.Get("username")
	lastname := r.Form.Get("lastname")

	// Comprobar si los campos están vacíos
	if username == "" || lastname == "" {
		http.Error(w, "Los campos de usuario y apellido son requeridos", http.StatusBadRequest)
		return
	}

	Insert := "INSERT INTO Users (username, lastname) VALUES (?, ?)"
	_, err = db.Exec(Insert, username, lastname)
	if err != nil {
		http.Error(w, "Error al registrar el usuario: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/?registered=true", http.StatusSeeOther)
}
