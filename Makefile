help:
	@more Makefile

deps: .git/hooks/prepare-commit-msg
	which dep || go get -u github.com/golang/dep/...
	dep ensure

.git/hooks/prepare-commit-msg:
	cp -f _misc/git/hooks/prepare-commit-msg .git/hooks/prepare-commit-msg

save:
	dep ensure -update
