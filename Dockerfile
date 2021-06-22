FROM golang:1.14

ENV GOPATH=/go

WORKDIR $GOPATH/src/assignement


# Copy the code into the container
COPY . .
# Copy and download dependency using go mod
RUN go get -d -v ./...

RUN echo `ls -R src`

# Build the application
RUN go build ./...


# This container exposes port 8080 to the outside world
EXPOSE 8080 8081

# Run the executable
CMD ["/cmd/fplserve"]