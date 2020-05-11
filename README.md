# GolangDockerTest

Test project for deploying Go applications and a React front-end to Docker (local and production).

## Pushing individual Docker containers

View the README in the `./web-server` page.

## Using Docker Compose

This project contains multiple services. Use `docker-compose` to spin them all up together. Every service should have its own `Dockerfile`, which `docker-compose` can use to spin up that resource. In addition, `docker-compose` can also help you spin up basic images like `MySql`. To run `docker-compose`, use this command in the same directory as the `docker-compose.yml`.

```unix
docker-compose build
docker-compose up
```

The webserver should be started on the hosts' (your computer's) port 6000, which maps to the container's port 8080. The database is also available at hosts' port 3307 which maps to another container's port 3306. Use curl to access the webserver living on the container `curl http://localhost:6000`. 

Once you are ready to clean up the docker containers, run `docker-compose rm`.

## Relationships between Docker-Compose services

If one server depends on another, follow this guide: <https://dev.to/hugodias/wait-for-mongodb-to-start-on-docker-3h8b>.

- You can have one server start before another using `depends_on`. 
- If one server needs to be "healthy" before another (like a database before a webserver), use `docker-compose-wait`. This is very useful if a service is "ready" but takes a while initializing before it can accept requests.
- Use `links` for containers to expose one another to each other. Host has access to all containers, but containers cannot access each other.

## Shell into Docker Containers

Let's say we want to see what's going on in our Docker Containers. For example, let's say I want to check the database status of one of my containers. I could simple use:

```unix
docker-compose exec db bash
```

This will open service `db` in a bash CLI. Because mysql is installed as a base, we can then use `mysql` commands once inside the container.

## Resources

- Go doc on deploying to Docker <https://blog.golang.org/docker>
- <https://stackoverflow.com/questions/47837149/build-docker-with-go-app-cannot-find-package>
- Setup Automated Builds on Docker <https://docs.docker.com/docker-hub/builds/>
- Using Docker Compose Up <https://www.youtube.com/watch?v=Qw9zlE3t8Ko>
- Waiting on services <https://dev.to/hugodias/wait-for-mongodb-to-start-on-docker-3h8b>