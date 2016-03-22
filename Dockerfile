FROM scratch

MAINTAINER Robert Pierce <robert.pierce@careerbuilder.com>

ADD hal-test-server /
ADD fixtures/* /fixtures/
EXPOSE 8000
ENTRYPOINT ["/hal-test-server"]
