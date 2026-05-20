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

## Flujo de trabajo con Git (Best Practices)

### Rama principal
- `main` siempre debe estar en estado desplegable
- Todo el desarrollo ocurre en ramas de feature temporales

### Flujo de trabajo recomendado

#### 1. Preparación
```bash
git checkout main
git pull
```

#### 2. Crear rama feature
```bash
git checkout -b feature/nombre-descriptivo
```

**Convenciones de nombres:**
- `feature/` - nueva funcionalidad (ej: `feature/post-endpoint`)
- `bugfix/` - corrección de errores
- `refactor/` - reestructuración de código
- `docs/` - documentación
- `test/` - pruebas

#### 3. Desarrollar feature
- Haz tus cambios
- Commits frecuentes y descriptivos:
  ```bash
  git add .
  git commit -m "Descripción clara del cambio"
  ```

#### 4. Push al remote
```bash
git push -u origin feature/nombre-descriptivo
```

#### 5. Crear Pull Request
1. En GitHub: "Compare & pull request" o Pull Requests → New PR
2. Base: `main`, Comparar: `feature/nombre-descriptivo`
3. Título claro, descripción detallada
4. Asignar revisores (si aplica)
5. Crear PR

#### 6. Atender feedback
- Haz cambios basado en review
- `git push` (actualiza PR automáticamente)

#### 7. Merge después de aprobación
- Vía GitHub UI (recomendado: "Squash and merge")
- O localmente:
  ```bash
  git checkout main
  git merge --no-ff feature/nombre-descriptivo
  git push origin main
  ```

#### 8. Limpieza
- Rama local: `git branch -d feature/nombre-descriptivo`
- Rama remota: usualmente se borra automáticamente tras merge en GitHub
- Actualizar main local: `git checkout main && git pull`

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

## Comandos útiles

### Ver estado
```bash
git status
```

### Ver ramas
```bash
git branch
git branch -r  # ramas remotas
```

### Ver commits
```bash
git log --oneline -10
```

### Comparar con main
```bash
git diff main feature/nombre-de-tu-rama
```

### Actualizar main
```bash
git checkout main
git pull
```