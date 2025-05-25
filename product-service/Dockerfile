FROM golang:1.24-alpine

WORKDIR /app

# Install bash (needed for wait-for-it.sh)
RUN apk add --no-cache bash

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

# Download wait-for-it.sh script
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

RUN go build -o main .

# Wait for product-db:3306 to be ready, then start the app
CMD ["/wait-for-it.sh", "product-db:3306", "--", "./main"]