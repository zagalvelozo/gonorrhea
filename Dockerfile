##
## Build
##
FROM golang:1.21 AS build
RUN mkdir -p $GOPATH/src/github.com/zagalvelozo/gonorrhea
ADD . $GOPATH/src/github.com/zagalvelozo/gonorrhea
WORKDIR $GOPATH/src/github.com/zagalvelozo/gonorrhea
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o /gonorrhea

##
## Run
##
FROM scratch
# Copy the ca-certificate.crt from the build stage
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /gonorrhea /bin/gonorrhea
CMD ["/bin/gonorrhea"]
