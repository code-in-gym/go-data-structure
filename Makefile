test:
	go vet ./... && go test ./... -count=1 -cover 

