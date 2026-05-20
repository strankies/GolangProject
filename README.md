# GolangProjects

API REST simple en Go usando gorilla/mux.

## Endpoints

- `GET /users` - Listar todos los usuarios
- `GET /users/{id}` - Obtener usuario por ID

## Ejecución local

```bash
go run main.go
```

El servidor corre en `http://localhost:8080`

## Docker

### Build de la imagen

```bash
docker build -t golang-projects .
```

### Ejecutar el contenedor

```bash
docker run -p 8080:8080 golang-projects
```

La API estará disponible en `http://localhost:8080`

### Versiones de la imagen

- `golang-projects:latest` - versión estándar
- `golang-projects:optimized` - versión optimizada (más pequeña y segura)

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
├── go.sum
├── Dockerfile
├── Dockerfile.optimized
└── README.md
```