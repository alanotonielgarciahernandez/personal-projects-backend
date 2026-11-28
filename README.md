<div align="center">
  <img src="./logo.svg" alt="Personal Projects" height="100">

  # Personal Projects Backend

  **Go backend that provides project data and Git operations for the Personal Projects frontend.**  
  **Backend en Go que provee datos de proyectos y operaciones Git para el frontend de Personal Projects.**

  [![Go](https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev/)
  [![Gin](https://img.shields.io/badge/Gin-Web_Framework-00ADD8?style=flat-square)](https://gin-gonic.com/)
  [![GraphQL](https://img.shields.io/badge/GraphQL-E10098?style=flat-square&logo=graphql&logoColor=white)](https://graphql.org/)
  [![MongoDB](https://img.shields.io/badge/MongoDB-4EA94B?style=flat-square&logo=mongodb&logoColor=white)](https://www.mongodb.com/)
  [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://opensource.org/licenses/MIT)

  [![Live Demo](https://img.shields.io/badge/Live_Demo-projects.alanotonielgarciahernandez.com-blue?style=flat-square)](https://projects.alanotonielgarciahernandez.com)
</div>

---

## What it does / Qué hace

- Stores and serves project metadata through a GraphQL API.  
  *Almacena y sirve metadatos de proyectos a través de una API GraphQL.*

- Allows downloading a project at any specific commit.  
  *Permite descargar un proyecto en cualquier commit específico.*

- Performs Git operations using libgit2.  
  *Realiza operaciones Git usando libgit2.*

## Endpoints / Endpoints

Full list available in [ENDPOINTS.md](./ENDPOINTS.md).  
*Lista completa disponible en [ENDPOINTS.md](./ENDPOINTS.md).*

Main routes / Rutas principales:

- `POST /query` — GraphQL endpoint (projects, issues, branches, commits, etc.)  
  *Endpoint GraphQL (proyectos, issues, branches, commits, etc.)*

- `GET /get-archive` — Download a project at a specific commit  
  *Descargar un proyecto en un commit específico*

- `GET /get-files` — Get the file tree of a project  
  *Obtener el árbol de archivos de un proyecto*

- `GET /get-content` — Get the content of a file  
  *Obtener el contenido de un archivo*

## Tech Stack / Tecnologías

- **Language / Lenguaje:** Go
- **Web Framework:** Gin
- **API:** GraphQL (gqlgen)
- **Database / Base de datos:** MongoDB
- **Git:** libgit2 (via git2go)

## How to run it / Cómo correrlo

### Requirements / Requisitos

- Go 1.20+
- MongoDB
- libgit2 1.5.0

### Environment variables / Variables de entorno

```env
DB_URI=mongodb://localhost:27017
REPOSITORIES_DIR=/path/to/repositories/
AUTHORIZATION_KEY=your-secret-key
```

### Steps / Pasos

```bash
git clone https://github.com/alanotonielgarciahernandez/personal-projects-backend.git
cd personal-projects-backend
go mod tidy
go run .
```

## Technical decisions / Decisiones técnicas

- GraphQL was chosen to give the frontend flexible queries over project data.  
  *Se eligió GraphQL para dar al frontend consultas flexibles sobre los datos de los proyectos.*

- libgit2 is used for Git operations instead of shelling out to the git CLI.  
  *Se usa libgit2 para las operaciones Git en lugar de llamar al CLI de git.*

- MongoDB stores project metadata while the actual repositories live on disk.  
  *MongoDB guarda los metadatos de los proyectos mientras que los repositorios reales viven en disco.*

## Related / Relacionado

Frontend: [personal-projects-frontend](https://github.com/alanotonielgarciahernandez/personal-projects-frontend)

## License / Licencia

MIT
