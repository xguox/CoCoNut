FROM golang:1.11.3
ENV GO111MODULE=on GIN_MODE=release
WORKDIR /apps
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
COPY ./config/conf.docker.yml ./config/conf.yml
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 9876
CMD [ "/apps/coconut" ]
