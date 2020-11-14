FROM scratch

WORKDIR $GOPATH/src/github.com/hucongyang/go-demo
COPY . $GOPATH/src/github.com/hucongyang/go-demo

EXPOSE 8888
CMD ["./go-demo"]
