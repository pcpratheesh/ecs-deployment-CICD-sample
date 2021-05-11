FROM golang:latest
# Add Maintainer Info
# Optional field to let you identify yourself as the maintainer of this image. This is just a label (it used to be a dedicated Docker directive).
LABEL maintainer="Pratheesh"

# Set the Current Working Directory inside the container
# Define the default working directory for the command defined in the “ENTRYPOINT” or “CMD” instructions
RUN mkdir /app
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Specify commands to make changes to your Image and subsequently the Containers started from this Image. 
# This includes updating packages, installing software, adding users, creating an initial database, setting up certificates, etc.
# These are the commands you would run at the command line to install and configure your application. 
# This is one of the most important dockerfile directives
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

## Our project will now successfully build with the necessary go libraries included.
RUN go build -o main cmd/main.go


## Our start command which kicks off
## our newly created binary executable
CMD ["/app/main"]