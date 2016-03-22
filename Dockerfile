FROM scratch

MAINTAINER Robert Pierce <robert.pierce@careerbuilder.com>

ADD hal-test-server /
EXPOSE 8000
ENTRYPOINT ["/hal-test-server"]
