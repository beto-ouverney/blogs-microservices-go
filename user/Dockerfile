FROM golang

WORKDIR /app-user

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8201

CMD ["main"]