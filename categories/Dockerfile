FROM golang

WORKDIR /app-categories

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8200

CMD ["main"]