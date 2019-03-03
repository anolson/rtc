.PHONY: test watch

GOBIN=$(GOPATH)/bin
REFLEX := $(GOBIN)/reflex
GOTEST := $(GOBIN)/gotest

$(GOTEST):
	go get -v github.com/rakyll/gotest

$(REFLEX):
	go get -v github.com/cespare/reflex

setup: $(GOTEST) $(REFLEX)

test: $(GOTEST)
	$(GOTEST) ./...

watch: $(REXLEX)
	$(REFLEX) -r ".go" --decoration=none -- make test
