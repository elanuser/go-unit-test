
.PYONY: clean
clean:
	rm -f cover_out cover_out.html

.PYONY: coverage
coverage:
	go test -cover ./...

.PYONY: coverage-html
coverage-html:
	go test -coverprofile=cover_out ./...
	go tool cover -html=cover_out -o cover_out.html	

bench-mark:
	go test -bench=.

