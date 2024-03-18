FROM golang:1.20 as build

WORKDIR /code
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app


FROM scratch as final
LABEL author "@ia-Simon [GitHub]"

COPY --from=build /app /

ENTRYPOINT [ "/app" ]