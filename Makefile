export GOBIN=$(shell pwd)/.bin

$(GOBIN)/depcheck: go.mod
	go install github.com/v-standard/go-depcheck/cmd/depcheck

.PHONEY: depcheck
depcheck: $(GOBIN)/depcheck
	go vet -vettool=$(GOBIN)/depcheck ./...
