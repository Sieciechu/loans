FROM golang:1.13 AS builder

WORKDIR /app
COPY . .
RUN go test ./... && ./bin/build.sh


FROM golang:1.13-alpine AS runner
WORKDIR /app
COPY --from=builder /app/loans /app/loans

CMD ./loans

