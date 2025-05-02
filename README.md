<!-- markdownlint-disable heading-start-left first-line-h1 -->
<!-- markdownlint-capture -->
<!-- markdownlint-disable no-inline-html heading-increment -->

<div align='center'>
  <img src='./logo.svg' alt='Alan Otoniel García Hernández Personal Projects' height='100px'>

  #### Back-end helping to reach project data.
</div>

# Functions

## Saving data.
Save the project data into the database.

## Project download at any point.
Select the point where you would like to download a project to try it on your machine.

# Information
Started in June 26 of 2024, The Personal Projects back-end was developed after watching 36 over 55 hours of a ReactJS course on Udemy and having previously finished 14.5 hours of a Go course on Udemy.
With the need to build a portfolio to increase my chances of getting a job. I made the decision to pause the course and work on this project to build my portfolio and to test and expand my knowledge and skill.

## How has this project helped to build my portfolio?
The Personal Projects back-end source code shows my understanding and skill in action when using most of the tools provided by the technology used and by implementing the project with systems using different technologies as currently used by leading companies.
In addition, using this project to expose my work and make it understandable to all the public helps to showcase my knowledge.

## How has this project helped to test and expand my knowledge?
Trying to replicate the GitHub back-end made me learn how it works on the inside, how long it takes to replicate the same application only with the main functionalities and how long would it take to add extra functionalities like codespaces.

## How has this project helped to test and expand my skills?
Working in an application all by myself without following steps from a course or tutorial made me train my logical process.

## Why did I make this project?
I was interested in how GitHub works to bring data from Git repositories to a website where users can inspect many projects in an accessible way. Also, making a copy of an existing software helps to skip the search for ideas process to start directly with the development process.

## What's next?
Of course, I will continue working on other projects I have in mind, expanding even more my knowledge on many different technologies, but also I will be maintaining and optimizing this project because I will use it to someday exhibit my projects on my own website.

# How to Use

## Requirements
[ Go ]( https://go.dev ) and [ MongoDB ]( https://mongodb.com ) are required to start working in this project.
After the installation, install the required Go packages with the command:
```sh
go get
```

## Environment Variables
- **DB_URI**: Your MongoDB Cluster URI, e.g "mongodb://localhost:27017".
- **REPOSITORIES_DIR**: You need to specify the repositories folder from the root path, e.g "C:/path/to/repositories/"
Note: It is important that the path ends with "/".
- **AUTHORIZATION_KEY**: Secret key to make GraphQL mutations.
- MongoDB Database name and collections can be found in the "cfg/variables.go" file of this project with its default values that you can freely change to match your needs.

## Handler configuration
You can change the CORS allowed origins and trusted proxies configuration at the "handlers/handlers.go" file.

## libgit2
This project uses the package git2go which requires the version 1.5.0 of libgit2 to make git operations.  
The version 1.5.0 of libgit2 is not included in the project and should be installed manually to make builds of the project, you can find installing information on the [ libgit2 GitHub repository ]( https://github.com/libgit2/libgit2 ).

## GraphQL model
When updating the "schema.graphqls" file make sure to add to the database models containing the ID value in the "models_gen.go" file 'bson:"_id"'.

# Technologies
- Programming Language: [ Go ]( https://github.com/golang/go )
- Web Framework: [ Gin Web Framework ]( https://github.com/gin-gonic/gin )
- Database: [ MongoDB ]( https://github.com/mongodb/mongo )
- Query Language: [ GraphQL ]( https://graphql.org/ )
- GraphQL Server Library: [ 99designs/gqlgen ]( https://github.com/99designs/gqlgen )
