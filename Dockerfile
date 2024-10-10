FROM golang:1.23.1

WORKDIR /site

COPY . .

RUN go mod download
RUN go build -o site

EXPOSE 8080

CMD ["./site"]