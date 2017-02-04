help:
	@more Makefile

deps: config.yml
	which dep || go get -u github.com/golang/dep/...
	dep ensure

save:
	dep ensure -update
