# Cowsaid

Cowsaid is an application to display cowsay as a service (CAAS).

```
   ~~~~~~~~~~~~~~~~~~~~~~~
 (   Cowsay-As-A-Service   )
   ~~~~~~~~~~~~~~~~~~~~~~~
          \   ^__^
           \  (oo)\_______
              (__)\       )\/\
                  ||----w |
                  ||     ||
```

## Description ##

Cowsaid is a web service to provide Cowsay as a service.


## TLDR ##

To get Cowsaid running with minimal effort run the DockerHub image.  View at http://localhost:8080.

        $ docker run -it --rm -p 8080:80 rcompos/cowsaid

## Installation ##

### Run Go Binary ##

Run Cowsaid application as a Go executable.

#### Requirements to run Go binary ####

The following are requirements to run the Cowsaid service as a Go executable.

* Go version 1.10
* Fortune (if running Go binary, optional)
MacOS: brew install fortune
Alpine: apk add git fortune
Debian: apt-get install fortune
CentOS: yum install fortune

* Environemntally variable: export GO111MODULE=on

To run Cowsaid as a Go executable, follow these steps to build the binary.

1. __Change to the cowsaid directory.__

        $ go version
        $ export GO111MODULE=on 
        $ cd path/to/cowsaid

2. __Build the application binary.__

        $ GO111MODULE=on go build

3. __Run the newly built Go binary.__

        $ ./cowsaid

4. __Command-line options__



### Build Docker Image ###

Build Cowsaid Docker image.

#### Requirements to run Docker image ####

The following are requirements to run the Cowsaid service as a Docker container.

* Docker version 17.1

To run Cowsaid as a Docker container, follow these steps to build the image.

1. __Change to the cowsaid directory.__

        $ docker version
        $ cd path/to/cowsaid

2. __Build the Docker image.__

        $ docker build -t cowsaid .

3. __Run the newly built image.__

        $ docker run -it --rm -p 8080:80 cowsaid


### Run DockerHub Image ###

The Cowsaid application can be run as a pre-built DockerHub image.  Note this image is a compact image built from Alpine Linux base image.

1. __Run application as Docker container.__

        $ docker run -it --rm -p 8080:80 rcompos/cowsaid


## Usage ##


1. View Cowsay Message

Once you have the Cowsaid application running as a service, you can In your web browser, enter the following URL:

        http://localhost:8080

2. API Endpoints

The following API endpoints are available:

```
/api/v1
/api/v1/cowhist
/api/v1/cowsay
/api/v1/info
/api/v1/ping
/api/v1/bad
/api/v1/err
/api/v1/count
```

3. Files

View the input, output and log files.

        http://localhost:8080

After the service is running, output and log files are created.

```
error.txt
```

Notes:
If building Go binary, the following produces a small images well-suited for containerization.  Substitute hardware architecture for GOOS (i.e. darwin, linux, etc).

	$ GO111MODULE=on CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o cowsaid .
