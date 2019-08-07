# TheWeatherDesk

TheWeatherDesk is a Go application to display tomorrow's forecasted high temperatures for location of each line in specified logfile.  The temperature values are displayed 
as a histogram and saved to tab separated file.

```
   ~~~~~~~~~~~~~~~~~~~~~~~
 (     TheWeatherDesk     )
   ~~~~~~~~~~~~~~~~~~~~~~~
          \   ^__^
           \  (oo)\_______
              (__)\       )\/\
                  ||----w |
                  ||     ||
```

## Description ##

For a given tab-separated text input file, TheWeatherDesk (TWD) search each line for IP address,
get tomorrow's weather forecast for high temperatures, display histogram.
Run web service with API endpoints.

The default input file is a log file with tab-separated values.  The default input file is included in the src directory (in.txt).  Each log line is expected to contain an IP address that will be used to geolocate.

When TheWeatherDesk is started, it begins processing the input file.  The default input files will take about 45 minutes to process.  The geolocation service (ip-api.com) will ban for over 150 requests per minute per source IP, so the process is slowed down accordingly.  Unban at http://ip-api.com/docs/unban

Note: TDW uses an IP and Point cache.  The cache is not purged of old values.

__Upstream services:__

Geolation: ip-api.com   
Weather: weather.gov   

## TLDR ##

To get TheWeatherDesk running with minimal effort run the DockerHub image.  Substitute _~/local_ with your local storage directory.  Copy input data file to _in.txt_.


        $ mkdir ~/local
        $ cp path/to/histogram_data ~/local/in.txt
        $ docker run -it --rm -v ~/local:/src -p 8080:8080 rcompos/twd-alpine

## Installation ##

### Run Go Binary ###

Run TWD application as a Go executable.

#### Requirements to run Go binary ####

The following are requirements to run the TWD service as a Go executable.

* Go version 1.10
* Fortune (if running Go binary, optional)
MacOS: brew install fortune
Alpine: apk add git fortune
Debian: apt-get install fortune
CentOS: yum install fortune

* Environemntally variable: export GO111MODULE=on

To run TWD as a Go executable, follow these steps to build the binary.  Note that the default dataset is 4082 lines.

1. __Change to the twd directory.__

        $ go version
        $ export GO111MODULE=on 
        $ cd path/to/twd

2. __Build the application binary.__

        $ go build

3. __Run the newly built Go binary.__

        $ ./twd -i path/input-file

4. __Command-line options__

```
Usage of twd:
  -b string
    	Bad lines file (default "./src/bad.txt")
  -c string
    	Histogram text chart file (default "./src/histogram.txt")
  -d string
    	Histogram data tab seperated file (default "./src/temperatures.tsv")
  -e string
    	Errors file (default "./src/error.txt")
  -h string
    	Histogram tsv bucket file (default "./src/histogram.tsv")
  -i string
    	Input file (default "./src/in.txt")
  -n int
    	Number of buckets (default 10)
  -p string
    	Output point cache file (default "./src/cache-point.csv")
  -q string
    	Output ip cache file (default "./src/cache-ip.csv")
```


### Build Docker Image ###

Build TWD Docker image.

#### Requirements to run Docker image ####

The following are requirements to run the TWD service as a Docker container.

* Docker version 17.1

To run TWD as a Docker container, follow these steps to build the image.

1. __Change to the twd directory.__

        $ docker version
        $ cd path/to/twd

2. __Build the Docker image.__

        $ docker build -t twd .

3. __Run the newly built image.  The number of buckets can be specified with the -n argument.  Substitute _~/local_ with your local storage directory.  Copy input data file to _in.txt_ __

        $ mkdir ~/local
        $ cp path/to/histogram_data ~/local/in.txt
        $ docker run -it --rm -v ~/local:/go/src/github.com/rcompos/twd/src -p 8080:8080 twd -n 10


### Run DockerHub Image ###

The TWD application can be run as a pre-built DockerHub image.  Note this image is a compact image built from Alpine Linux base image.

1. __Run application as Docker container. Substitute _~/local_ with your local storage directory.  Copy input data file to _in.txt_ __

        $ mkdir ~/local
        $ cp path/to/histogram_data ~/local/in.txt
        $ docker run -it --rm -v ~/local:/src -p 8080:8080 rcompos/twd-alpine


## Usage ##


1. View Histogram

Once you have the TWD application running as a service, you can In your web browser, enter the following URL:

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

        http://localhost:8080/files

After the service is running, output and log files are created.

```
bad.txt
cache-ip.csv
cache-point.csv
error.txt
histogram.tsv
histogram.txt
in.txt
temperatures.tsv
twd-log-2019-04-21-211402.out
```

Notes:
If building Go binary, the following produces a small images well-suited for containerization.  Substitute hardware architecture for GOOS (i.e. darwin, linux, etc).

	$ CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o twd .
