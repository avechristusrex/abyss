FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /abyss ./cmd/abyss
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /abyss /abyss
EXPOSE 9000 8080
ENTRYPOINT ["/abyss"]
