FROM golang:alpine AS builder

ENV GO111MODULE=on  \
    CGO_ENABLED=0   \
    GOOS=linux  \
    GOARCH=amd64

# move to working directory build
WORKDIR /build

# copy and download dependencices with go.mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# COPY the source code from current dir to working dir into the container
COPY . .

# build the application
RUN go build -o main ./cmd/app/product/

# move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# copy binary from bulid to main folder
RUN cp /build/main .

# Start a new stage - building a small image from Scratch ###

FROM scratch

# copy the prebuilt binary from the previous stage
COPY --from=builder /dist/main /

# Expose the app port
EXPOSE 7070

# command to run
ENTRYPOINT [ "/main" ]