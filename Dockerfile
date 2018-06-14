FROM scratch

ENV verbose=true

COPY build/ipd2-linux_amd64 /

ENTRYPOINT ["/ipd2-linux_amd64"]

CMD ["--verbose"]