FROM registry.access.redhat.com/ubi8/go-toolset
USER root
RUN mkdir /build/
ADD . /build/
RUN cd /build/ && go build

FROM registry.access.redhat.com/ubi8/ubi-minimal
COPY --from=0 /build/HelloBeego /usr/local/bin/app
CMD app
EXPOSE 8080