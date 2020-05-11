# Setup Instructions

Follow these setup instructions to build this application, and deploy it to a local Docker container

## Building the Go Application

Run the following command to create a binary for this application.

```unix
go install .
```

You should now see a binary in the `bin` folder, which you can also use to run the application. 
Also run this command to ensure this module is Go-gettable. This is important in order to download the project and its dependencies in external environments, like Docker!

```unix
go get github.com/ahsu1230/GolangDockerTest
```

## Deploying to Local Docker

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

## Fetching Docker instance from DockerHub

Setup Automated Builds by connecting Docker to a Github repository. When new commits are merged, automatic Docker builds will be triggered. Once builds are successfully triggered on DockerHub, anyone with docker installed can fetch and run this image as well.

```unix
docker run ahsu1230/golang_docker_test
```