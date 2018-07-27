TAG=v0.0.1-$(shell date +%s)
docker:
	@docker run --rm --name golang -v $(shell pwd):/go/src/harbor_upgrade --workdir /go/src/harbor_upgrade golang:1.8.3-alpine3.6 /bin/sh -c "apk add --update build-base && go build --ldflags -w -o harbor_upgrade"
image: docker
	docker build -t harbor.enncloud.cn/qinzhao-harbor/syncdb:$(TAG) .
run:
	go run main.go  -config config.yaml