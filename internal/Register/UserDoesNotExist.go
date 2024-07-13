package register

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

func UserDoesNotExist(w http.ResponseWriter, r *http.Request) {

	// analiza los datos enviados al formulario
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "error 400", http.StatusBadRequest)
		return
	}

	//  obtener los valores del formulario
	username := r.Form.Get("username")
	lastname := r.Form.Get("lastname")

	// consulta sql para verificar si el usuario existe
	Verify := "SELECT username FROM Users WHERE username = ? AND lastname = ? LIMIT 1"

	var result string
	// lo escanea y manda el resultado
	err = db.QueryRow(Verify, username, lastname).Scan(&result)

	switch {
	case err == sql.ErrNoRows:
		// si no existe el usuario manda este mensaje:
		tpl := `<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Usuario no encontrado</title>
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

        p {
            font-size: 16px;
            margin-bottom: 20px;
        }

        a {
            color: #007BFF;
            text-decoration: none;
            font-weight: 600;
            transition: color 0.3s;
        }

        a:hover {
            color: #0056b3;
            text-decoration: underline;
        }
    </style>
</head>
<body>
    <div class="container">
        <h2>Usuario no encontrado</h2>
        <p>El usuario no existe. Por favor, <a href="/register?username={{.Username}}&lastname={{.Lastname}}">registre su usuario aquí</a>.</p>
    </div>
</body>
</html>
`

		data := struct {
			Username string
			Lastname string
		}{
			Username: username,
			Lastname: lastname,
		}

		// Crea una nueva plantilla "user-not-found" y analizar tpl
		t, err := template.New("user-not-found").Parse(tpl)
		if err != nil {
			http.Error(w, "Error al renderizar el mensaje", http.StatusInternalServerError)
			return
		}
		// ejecuta los datos enviados y manda un resultado al ResponseWriter(w)
		err = t.Execute(w, data)
		if err != nil {
			http.Error(w, "Error al renderizar el mensaje", http.StatusInternalServerError)
			return
		}

	case err != nil:
		// si no funciona manda un error 500
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	default:
		// si funciona manda el mensake con el resultado obtenido del formulario
		fmt.Fprintf(w, "Bienvenido, %s!", result)
	}
}