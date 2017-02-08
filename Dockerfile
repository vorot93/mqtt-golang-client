from centurylink/ca-certs

ADD ./bin/app /usr/bin/app

ENTRYPOINT [ "/usr/bin/app" ]
