# GolangDockerTest

Test project for deploying Go applications to Docker (local and production)

## Setup Instructions

Follow these setup instructions to build this application, and deploy it to a local Docker container

### Building the Go Application

Run the following command to create a binary for this application.

```unix
go install .
```

You should now see a binary in the `bin` folder, which you can also use to run the application. 
Also run this command to ensure this module is Go-gettable. This is important in order to download the project and its dependencies in external environments, like Docker!

```unix
go get github.com/ahsu1230/GolangDockerTest
```

### Deploying to Local Docker

You can view Docker configurations in the `Dockerfile` of the root directory.

Use this to build a Docker image. The image is built from the directory that contains the `Dockerfile`. It will then give the image a name/tag `tester`. Use `docker images` to see a list of previously built images.

```unix
docker build . -t tester
docker images
```

Use this command to start and run the target Docker image in a new container. Here we are running image `tester` and naming the container `test_docker`. `--rm` means to delete the container if the app is finished running.

```unix
docker run --name testdocker --rm --publish 6060:8080 tester
```

Check the DockerHub to see that the container is running. The `publish` command means that the Docker container's port `:8080` is exposed to the external `:6060` port. So you may only interact with `localhost:6060`.

Use this curl command to verify:

```
curl http://localhost:6060/ping  <- works!
curl http://localhost:8080/ping  <- does not work
```

Stop the container using:

```
docker stop testdocker
```

### Fetching Docker instance from DockerHub

Setup Automated Builds by connecting Docker to a Github repository. When new commits are merged, automatic Docker builds will be triggered. Once builds are successfully triggered on DockerHub, anyone with docker installed can fetch and run this image as well.

```unix
docker run ahsu1230/golang_docker_test
```

## Using Docker Compose

If you want to create multiple services, use `docker-compose` to spin them all up together. Every service should have its own `Dockerfile`, which `docker-compose` can use to spin up that resource. In addition, `docker-compose` can also help you spin up basic images like `MySql`. To run `docker-compose`, use this command in the same directory as the `docker-compose.yml`.

```unix
docker-compose build
docker-compose up
```

The webserver should be started at port 6000, which maps to the container's port 8080. The database is now available at port 3307 which maps to another container's port 3306. Use curl to access the webserver living on the container `curl http://localhost:6000`.

Once you are ready to clean up the docker containers, run `docker-compose rm`.

## Resources

- Go doc on deploying to Docker <https://blog.golang.org/docker>
- <https://stackoverflow.com/questions/47837149/build-docker-with-go-app-cannot-find-package>
- Setup Automated Builds on Docker <https://docs.docker.com/docker-hub/builds/>
- Using Docker Compose Up <https://www.youtube.com/watch?v=Qw9zlE3t8Ko>
- Waiting on services <https://dev.to/hugodias/wait-for-mongodb-to-start-on-docker-3h8b>