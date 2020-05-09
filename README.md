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

## Resources