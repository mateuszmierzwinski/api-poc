FROM scratch
MAINTAINER mateuszmierzwinski@gmail.com

COPY domain.crt /
COPY domain.key /
COPY api-poc-linux-amd64 /
COPY static /static/

EXPOSE 8443
ENTRYPOINT ["/api-poc-linux-amd64"]