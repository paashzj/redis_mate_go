FROM ttbb/base:go AS build
COPY . /opt/sh/compile
WORKDIR /opt/sh/compile/pkg
RUN go build -o redis_mate .


FROM ttbb/redis:nake

LABEL maintainer="shoothzj@gmail.com"

COPY docker-build /opt/sh/redis/mate

COPY --from=build /opt/sh/compile/pkg/redis_mate /opt/sh/redis/mate/redis_mate

CMD ["/usr/bin/dumb-init", "bash", "-vx", "/opt/sh/redis/mate/scripts/start.sh"]
