FROM golang:1.17 as build-stage

WORKDIR /go/src/app
COPY /src .

RUN go build 

FROM golang:1.17 as prod-state

COPY --from=build-stage /go/src/app/backend /usr/local/bin

EXPOSE 8000/tcp

ENTRYPOINT ["/usr/local/bin/backend"]

