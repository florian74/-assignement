FROM golang:1.14

ENV GOPATH=/go
ENV GOBIN=/go/bin

WORKDIR $GOPATH/src/assignement


# Copy the code into the container
COPY . .
# Copy and download dependency using go mod
RUN go get -d -v ./...

# Build the application
RUN go install ./...
#RUN go install github.com/florian74/assignement/cmd/...


# This container exposes port 8080 to the outside world
EXPOSE 8080 8081

# Run the executable


RUN ["chmod", "+x", "/go/bin/fplserve"]
CMD ["fplserve"]