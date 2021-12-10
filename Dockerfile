FROM golang:1.16 as builder

WORKDIR /app

COPY main.go main.go
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download
RUN CGO_ENABLED=0 go build -o /mybinary main.go


FROM scratch
COPY --from=builder /mybinary /app
CMD ["/app"]