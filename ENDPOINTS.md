# Endpoints / Endpoints

## GraphQL — `POST` `/query`

Main GraphQL endpoint. Requires the `Authorization` header with the secret key for mutations.  
*Endpoint principal de GraphQL. Requiere el header `Authorization` con la clave secreta para las mutaciones.*

### Queries / Consultas

- `projects(page, filter)` — List projects  
  *Listar proyectos*

- `issues(page, filter)` — List issues  
  *Listar issues*

- `branches(page, project)` — List branches of a project  
  *Listar branches de un proyecto*

- `commits(page, project)` — List commits of a project  
  *Listar commits de un proyecto*

- `tags(page, project)` — List tags of a project  
  *Listar tags de un proyecto*

- `release(project, reference)` — Get a specific release  
  *Obtener un release específico*

### Mutations / Mutaciones

- `createProject` / `updateProject` / `deleteProject`
- `createIssue` / `updateIssue` / `deleteIssue`
- `createRelease` / `updateRelease` / `deleteRelease`

---

## REST Endpoints / Endpoints REST

### `GET` `/get-archive`

Download a project as an archive at a specific commit.  
*Descargar un proyecto como archivo en un commit específico.*

**Headers / Headers:**
- `Project`: project name
- `Reference`: commit hash or branch/tag name

---

### `GET` `/get-commit`

Get details of a specific commit.  
*Obtener detalles de un commit específico.*

**Headers / Headers:**
- `Project`: project name
- `Reference`: commit hash

---

### `GET` `/get-content`

Get the content of a file at a specific reference.  
*Obtener el contenido de un archivo en una referencia específica.*

**Headers / Headers:**
- `Project`: project name
- `Reference`: commit hash or branch/tag
- `File`: file path

---

### `GET` `/get-files`

Get the file tree of a project at a specific reference.  
*Obtener el árbol de archivos de un proyecto en una referencia específica.*

**Headers / Headers:**
- `Project`: project name
- `Reference`: commit hash or branch/tag
- `Path`: directory path (optional)

---

### `GET` `/get-importantfiles`

Get important files (README, license, etc.) of a project.  
*Obtener archivos importantes (README, licencia, etc.) de un proyecto.*

**Headers / Headers:**
- `Project`: project name
- `Reference`: commit hash or branch/tag

---

### `GET` `/uploads/:filename`

Serve uploaded images (project logos, backgrounds).  
*Servir imágenes subidas (logos y fondos de proyectos).*
