FROM golang:1.23.5-alpine3.21

WORKDIR /app

RUN go install github.com/bokwoon95/wgo@latest

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
	--mount=type=cache,target=/root/.cache/go-build \
	go mod download

COPY . .

ENTRYPOINT ["wgo", "run", "."]

