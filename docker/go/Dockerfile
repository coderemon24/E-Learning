FROM golang:1.24-alpine

WORKDIR /app

# Needed for HTTPS + git fetch
RUN apk add --no-cache git ca-certificates tzdata

# (optional but helpful) if corporate/proxy env exists, allow build-time args
ARG HTTP_PROXY
ARG HTTPS_PROXY
ARG NO_PROXY
ENV HTTP_PROXY=$HTTP_PROXY \
    HTTPS_PROXY=$HTTPS_PROXY \
    NO_PROXY=$NO_PROXY

# Go module download stability
ENV GOPROXY=https://proxy.golang.org,direct

# Install air (pin version for reliability)
RUN go install github.com/air-verse/air@v1.61.7

EXPOSE 8080
CMD ["air"]