FROM golang:1.12 as build
COPY . /app
WORKDIR /app
ENV CGO_ENABLED=0
RUN go mod vendor
RUN go build

FROM busybox 
COPY --from=build /app/taildir /bin/
USER nobody
ENTRYPOINT ["taildir"]