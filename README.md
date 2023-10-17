# simple-go-docker-gcp



To start :  `go run main`\
To stop : `ctr + c`

or if you want to run this from docker:
start your docker daemon
create image: `docker build -t sgdg-app:v1 .`
create and run container `docker run -d -p 8080:9000 --name simple-go-docker-gcp-api sgdg-app:v1`
make sure it is running: `docker container ls`
On your browser: `localhost:8080`