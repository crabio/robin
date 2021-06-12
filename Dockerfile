############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Create application user
ENV USER=robin
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR $GOPATH/src/mypackage/myapp/

# Copy source files
COPY . .

# Fetch dependencies.
RUN go get -d -v

# Build the binary.
RUN CGO_ENABLED=0 go build -o /go/bin/main

############################
# STEP 2 build a small image
############################
FROM scratch

# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable.
COPY --from=builder /go/bin/main /go/bin/main

# Add ENV params
ENV NATS_HOSTNAME=localhost
ENV NATS_PORT=4222

ENV SECRETS_FOLDER_PATH=/usr/local/share/robin/secrets

ENV GOOGLE_REDIRECT_URL=http://localhost:9000/auth

# Run the main binary.
ENTRYPOINT ["/go/bin/main"]