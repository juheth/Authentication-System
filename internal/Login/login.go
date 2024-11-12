package login

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

// LoginFormHandler handles the login form rendering
func LoginFormHandler(w http.ResponseWriter, r *http.Request) {
	tpl := `<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Login</title>
    <style>
        body {
            font-family: 'Roboto', sans-serif;
            background: linear-gradient(135deg, #333, #555);
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
            color: #fff;
        }

        .login-container {
            background-color: #222;
            padding: 40px 30px;
            border-radius: 15px;
            box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5);
            max-width: 400px;
            width: 100%;
            text-align: center;
            position: relative;
            overflow: hidden;
        }

        h2 {
            color: #fff;
            font-weight: 700;
            text-transform: uppercase;
            font-size: 28px;
            margin-bottom: 20px;
        }

        .form-group {
            margin-bottom: 20px;
            text-align: left;
        }

        label {
            display: block;
            margin-bottom: 10px;
            color: #ddd;
            font-weight: 600;
            text-align: left;
        }

        input[type="text"] {
            width: 100%;
            padding: 12px;
            margin-bottom: 20px;
            border: none;
            border-radius: 5px;
            box-sizing: border-box;
            font-size: 16px;
            background: #444;
            color: #fff;
        }

        input[type="text"]::placeholder {
            color: #bbb;
        }

        .register-link {
            color: #007BFF;
            text-decoration: none;
            font-weight: 600;
            display: block;
            margin-top: 20px;
        }

        .register-link:hover {
            text-decoration: underline;
        }

        .btn {
            background-color: #007BFF;
            color: #fff;
            padding: 15px 20px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            font-size: 18px;
            width: 100%;
            transition: background-color 0.3s, color 0.3s;
        }

        .btn:hover {
            background-color: #0056b3;
        }
    </style>
</head>
<body>
    <div class="login-container">
        <h2>Login</h2>
        {{if .Message}}
        <p class="message">{{.Message}}</p>
        {{end}}
        <form action="/login" method="post" class="login-form">
            <div class="form-group">
                <label for="username">Usuario:</label>
                <input type="text" id="username" name="username" value="{{.Username}}" required>
            </div>
            <div class="form-group">
                <label for="lastname">Apellido:</label>
                <input type="text" id="lastname" name="lastname" value="{{.Lastname}}" required>
            </div>
            {{if .ShowRegisterLink}}
            <p class="register-link">El usuario no existe. Por favor, <a href="/register?username={{.Username}}&lastname={{.Lastname}}">registre su usuario aquí</a>.</p>
            {{end}}
            <button type="submit" class="btn">Iniciar sesión</button>
        </form>
    </div>
</body>
</html>`

	// Create the template and parse the HTML
	t, err := template.New("login-form").Parse(tpl)
	if err != nil {
		http.Error(w, "Error al analizar la plantilla", http.StatusInternalServerError)
		return
	}

	// Structure for passing data to the template
	type PageData struct {
		Message          string
		Username         string
		Lastname         string
		ShowRegisterLink bool
	}
	data := PageData{}

	// Check if the "registered" query parameter is set
	if r.URL.Query().Get("registered") == "true" {
		data.Message = "Usuario registrado con éxito. Por favor, inicie sesión."
	}

	// Check if "showRegisterLink" is set to true
	if r.URL.Query().Get("showRegisterLink") == "true" {
		data.ShowRegisterLink = true
		data.Username = r.URL.Query().Get("username")
		data.Lastname = r.URL.Query().Get("lastname")
	}

	// Execute the template with data and send the result to the ResponseWriter
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Error al ejecutar la plantilla", http.StatusInternalServerError)
		return
	}
}
