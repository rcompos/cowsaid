FROM golang:1.13 as build
WORKDIR /build
COPY . .
RUN GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o cowsaid ./cowsaid.go

FROM alpine:latest
LABEL maintainer="Ron Compos <rcompos@gmail.com>"
COPY --from=build /build/cowsaid /bin
RUN apk add fortune
COPY fortunes-alt /usr/share/fortunes-alt
RUN cd /usr/share/fortunes-alt; for f in `find . -type d | grep -v '^.$'`; do echo $f; strfile $f/$f $f/$f.dat; done
EXPOSE 80
ENTRYPOINT ["/bin/cowsaid"]

