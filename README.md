# CAAS

Cowsay-As-A-Service (CAAS) is an application to display cowsay.

```
   ~~~~~~~~~~~~~~~~~~~~~~~
 (  Cowsay-As-A-Service     )
   ~~~~~~~~~~~~~~~~~~~~~~~
          \   ^__^
           \  (oo)\_______
              (__)\       )\/\
                  ||----w |
                  ||     ||
```

## Description ##

Cowsay as a service.


## TLDR ##

To get CAAS running with minimal effort run the DockerHub image. 

        $ docker run -it --rm -p 80:80 rcompos/caas

## Installation ##

### Run Go Binary ###

Run CAAS application as a Go executable.

#### Requirements to run Go binary ####

The following are requirements to run the CAAS service as a Go executable.

* Go version 1.10
* Fortune (if running Go binary, optional)
MacOS: brew install fortune
Alpine: apk add git fortune
Debian: apt-get install fortune
CentOS: yum install fortune

* Environemntally variable: export GO111MODULE=on

To run CAAS as a Go executable, follow these steps to build the binary.  Note that the default dataset is 4082 lines.

1. __Change to the caas directory.__

        $ go version
        $ export GO111MODULE=on 
        $ cd path/to/caas

2. __Build the application binary.__

        $ go build

3. __Run the newly built Go binary.__

        $ ./caas

4. __Command-line options__



### Build Docker Image ###

Build CAAS Docker image.

#### Requirements to run Docker image ####

The following are requirements to run the CAAS service as a Docker container.

* Docker version 17.1

To run CAAS as a Docker container, follow these steps to build the image.

1. __Change to the caas directory.__

        $ docker version
        $ cd path/to/caas

2. __Build the Docker image.__

        $ docker build -t caas .

3. __Run the newly built image.__

        $ docker run -it --rm -p 80:80 caas


### Run DockerHub Image ###

The CAAS application can be run as a pre-built DockerHub image.  Note this image is a compact image built from Alpine Linux base image.

1. __Run application as Docker container.__

        $ docker run -it --rm -p 80:80 rcompos/caas


## Usage ##


1. View Histogram

Once you have the CAAS application running as a service, you can In your web browser, enter the following URL:

        http://localhost:80

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

        http://localhost:80

After the service is running, output and log files are created.

```
error.txt
```

Notes:
If building Go binary, the following produces a small images well-suited for containerization.  Substitute hardware architecture for GOOS (i.e. darwin, linux, etc).

	$ CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o caas .
