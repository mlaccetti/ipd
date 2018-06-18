FROM golang:1.10.3-stretch as build
ARG TRAVIS_TAG
WORKDIR /go/src/github.com/mlaccetti/ipd2
COPY cmd/ /go/src/github.com/mlaccetti/ipd2/cmd/
COPY internal/ /go/src/github.com/mlaccetti/ipd2/internal/
COPY Makefile Gopkg.* index.html /go/src/github.com/mlaccetti/ipd2/
RUN \
  curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
  make deps
RUN make

FROM scratch
ARG TRAVIS_TAG
COPY --from=build /go/src/github.com/mlaccetti/ipd2/build/ipd2-${TRAVIS_TAG}-linux_amd64 /ipd2
ENTRYPOINT ["/ipd2"]
CMD ["--verbose"]