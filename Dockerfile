FROM golang:1.20 as build_stage

# Copy application data into image
COPY . /go/src/
WORKDIR /go/src/

# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download
RUN go mod tidy

# Build our application.
# RUN CGO_ENABLED=0 go build -o sender .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sender .

FROM scratch

COPY --from=build_stage /go/src/sender /bin/sender

EXPOSE 3000

ENTRYPOINT ["/bin/sender"]