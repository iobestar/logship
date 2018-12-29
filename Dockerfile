FROM alpine:3.2

MAINTAINER iobestar <ivica.obestar@gmail.com>

COPY logship /bin/logship

EXPOSE 11034
ENTRYPOINT ["/bin/logship"]
