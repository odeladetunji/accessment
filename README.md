# accessment
Golang

Docker is used to containerize the application

The root of the project contains docker-compose.yml and Dockerfile

To start the app run 

Install docker on your machine

cd into the root directory and run 

`docker build .`

`docker-compose build .`

Please you have to provide the following env variables in the docker-compose.yml file
`GITACCESSTOKEN`
`OWNER`
`REPONAME`

these are currently empty strings, they are however required

Please use your personal Git Access token for `GITACCESSTOKEN`



