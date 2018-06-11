# Docker
## Basic Commands
The pull command fetches an image from Docker registry, much like `npm install`.
```
docker pull busybox
```

Use `images` command to see a list of all images on your system.
```
docker images
```

Use `run` to run a container based on the provided image.
```
docker run busybox
```

Use `ps` to see all the containers that are currently running.
```
docker ps
```

Attach `-a` to show all the containers that ran.
```
docker ps -a
```

We can SSH into a docker container using a `-it` flag
```
docker run -it busybox sh
```

We can remove containers using `rm`, even those who have already exited.
```
docker rm 91e3611bf6ae
```

We can delete all containers in one go. The `-q` flag means returning only numeric IDs and `-f` 
flag means filter ouput based on conditions provided.
```
docker ps -a -q -f status=exited
```

The command above will filter docker containers and return numeric IDs of those selected containers,
now we can go ahead and delete them.
```
docker rm $(docker ps -a -q -f status=exited)
```

Now if we want to run a container and immediately delete it afterward then we can use a flag `--rm`
```
docker run busybox --rm
```

## Web Applications
Let's pull an image from Prakhar repository, the guy who wrote the tutorial.
```
docker run prakhar1989/static-site
```

At this point, the container is not exposing any port. We need to run the container is a detached
mode so we can ask the container to publish ports. The flag `-d` means detached mode. The flag `-P`
will publish all exposed ports to random ports. Finally, the flag `--name` corresponds to a name we
want to give to the container.
```
docker run --detach --publish-all --name static-site prakhar1989/static-site
```

Now we ask the docker what are the ports.
```
docker port static-site
```

## Create Images
* Base images are images that have no parent image, usually images with an OS like ubuntu, busybox
or debian.
* Child images are images that build on base images and add additional functionality.
* Official images are images that are officially maintained and supported by people at Docker.
* Usere images are images created and shared by users.

Let's begin with creating a Dockerfile.
```Dockerfile
FROM golang:1.8 as build

LABEL authors="Calvin Feng"

# The public image golang:1.8 specified that all source code must go into /go/src/
COPY . /go/src/react-ts-go-todo
WORKDIR /go/src/react-ts-go-todo

EXPOSE 3000

RUN go get -u github.com/golang/dep/cmd/dep

RUN go install

CMD react-ts-go-todo
```

After Dockerfile has been specified, build the image using `build`, and the flag `-t` means tag name.
The actual tag comes after the colon. In this case, the tag is `beta`.
```
docker build -tag react-ts-go-todo:beta .
```

Now let's run it, using publish flag to forward 3000 to 8000 because we already exposed 3000.
```
docker run --publish 8000:3000 react-ts-go-todo:beta
```

Or run it in detach mode which is more convenient
```
docker run --publish 8000:3000 --detach react-ts-go-todo:beta
```

Stop them all
```
docker stop $(docker ps -q)
```