FROM golang

# Copy the local files
ADD . /go/src/learning-go/main

# Build the outyet command inside the container.
# Manually fetch dependencies
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/gorm/dialects/sqlite

RUN go install github.com/gin-gonic/gin
RUN go install github.com/jinzhu/gorm
RUN go install github.com/jinzhu/gorm/dialects/sqlite
RUN go install learning-go/main

# Run the main command by default when the container starts.
ENTRYPOINT /go/bin/main

EXPOSE 8080