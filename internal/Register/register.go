package register

import (
	"database/sql"
	"html/template"
	"net/http"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func RegisterFormHandler(w http.ResponseWriter, r *http.Request) {
	// obtener los valores que se mandan al URL
	username := r.URL.Query().Get("username")
	lastname := r.URL.Query().Get("lastname")

	// formulario de registro
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
</html>
`

	// Crear una nueva plantilla llamada "register-form" y analizar el HTML de la plantilla
	t, err := template.New("register-form").Parse(tpl)
	if err != nil {
		http.Error(w, "Error al renderizar el formulario", http.StatusInternalServerError)
		return
	}

	data := struct {
		Username string
		Lastname string
	}{
		Username: username,
		Lastname: lastname,
	}

	// ejecuta los datos enviados y manda un resultado al ResponseWriter(w)
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Error al renderizar el formulario", http.StatusInternalServerError)
		return
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	// analiza los datos enviados al formulario
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error al analizar los datos del formulario", http.StatusBadRequest)
		return
	}

	// obtener los valores del formulario
	username := r.Form.Get("username")
	lastname := r.Form.Get("lastname")

	// sql para agregar el nuevo usuario
	Insert := "INSERT INTO Users (username, lastname) VALUES (?, ?)"
	_, err = db.Exec(Insert, username, lastname)
	if err != nil {
		http.Error(w, "Error al renderizar el formulario", http.StatusInternalServerError)
		return
	}

	// Redirigir al formulario de login con mensaje de registrado con éxito
	http.Redirect(w, r, "/?registered=true", http.StatusSeeOther)
}
