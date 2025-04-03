FROM scratch AS build

ADD alpine-minirootfs-3.21.3-x86_64.tar /

ARG VERSION="1.0"

SHELL ["/bin/sh", "-c"]

RUN apk update && apk add --no-cache go git build-base

WORKDIR /app

COPY main.go .

ENV VERSION=$VERSION

RUN go build -o builder main.go && \
    ./builder && \
    cp index.html /tmp/index.html

FROM nginx:alpine

COPY --from=build /tmp/index.html /usr/share/nginx/html/index.html

HEALTHCHECK --interval=30s --timeout=3s --retries=3 \
  CMD wget -q -O- http://localhost:80 || exit 1

CMD ["nginx", "-g", "daemon off;"]
    