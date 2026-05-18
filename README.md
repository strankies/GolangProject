# GolangProjects

API REST simple en Go usando gorilla/mux.

## Endpoints

- `GET /users` - Listar todos los usuarios
- `GET /users/{id}` - Obtener usuario por ID

## Ejecución

```bash
go run main.go
```

El servidor corre en `http://localhost:8080`

## Estructura del proyecto

```
GolangProjects/
├── api/
│   └── router.go
├── controllers/
│   └── usercontroller.go
├── models/
│   └── user.go
├── utils/
├── main.go
├── go.mod
└── go.sum
```