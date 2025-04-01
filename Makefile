.PHONY: help
help:
	@printf "\nmake targets:\n\n"
	@printf "  Help:\n"
	@printf "    help .................. This, what this is?\n"
	@printf "  Code Tooling:\n"
	@printf "    generate .............. Run go generate\n"
	@printf "  Linting:\n"
	@printf "    lint .................. Run all lints, ensures generated code is up-to-date\n"
	@printf "  Testing:\n"
	@printf "    test .................. Run all tests, generate all code, run all lints\n"
	@printf "    quicktest ............. Run all tests, quickly!\n"
	@printf "  Dev Server:\n"
	@printf "    server ................ Serveux le HTTP\n"
	@printf "\n"

.PHONY: generate
generate:
	go generate -x ./...

.PHONY: server
server:
	go run ./main.go

.PHONY: lint
lint:
	go tool gofumpt -l -w .
	go tool revive -config .revive.toml -formatter friendly ./...
	go tool golangci-lint run
	go tool govulncheck ./...

.PHONY: test-run-actual
test-run-actual:
	if printf "%s" "${GOTESTFLAGS}" | egrep -q -- "-race"; then \
		export GORACE="atexit_sleep_ms=50" ; \
	fi ; \
	go test ./... -shuffle=on ${GOTESTFLAGS} -coverpkg=./... -coverprofile test-coverage.out.tmp
	@[ -n "$$(which perl)" ] && perl -le 'print "🟩" x ($$ENV{"TERM"} ? (qx/tput cols/ / 2) : 40 )' || true

.PHONY: test
test: generate lint test-run test-wrapup

.PHONY: test-run
test-run:
	@# CGO required for -race
	CGO_ENABLED=1 $(MAKE) GOTESTFLAGS="-race" test-run-actual

# wow -race takes a while
.PHONY: quicktest
quicktest: quicktest-run test-wrapup

.PHONY: quicktest-run
quicktest-run:
	@# Ensure we can run without CGO
	CGO_ENABLED=0 $(MAKE) GOTESTFLAGS="" test-run-actual

.PHONY: test-wrapup
test-wrapup: test-coverage-process test-coverage-update-readme

.PHONY: test-coverage-process
test-coverage-process:
	@printf "Processing test-coverage:"
	@printf " filtering coverage"
	@egrep -v 'mock_[a-zA-Z0-9_-]+\.go' test-coverage.out.tmp | egrep -v '/id/[a-zA-Z0-9_-]+_id.go' | grep -v '_enum.go' | grep -v '_test.go' > test-coverage.out
	@printf " ✔"
	@rm test-coverage.out.tmp
	@# Get the main branch contents of test-coverage.html and put it
	@# in test-coverage.main.html to allow side-by-side comparisons
	@if git show-branch main 2>/dev/null >/dev/null; then git show main:test-coverage.html > test-coverage.main.html; fi
	@# Get the current branch (HEAD) contents of test-coverage.html and put it
	@# in test-coverage.HEAD.html to allow side-by-side comparisons
	@git show HEAD:test-coverage.html > test-coverage.HEAD.html || true
	@# Generate test-coverage.html from test-coverage.out
	@printf ", generating html"
	@go tool cover -html=test-coverage.out -o test-coverage.html
	@printf " ✔"
	@# Generate test-coverage.txt with per-file coverage statistics
	@perl -lane '($$f, $$c) = (m/<option value="file\d+">(\S+) (.*)<\/option>/); print "$$f => $$c" if $$f' test-coverage.html  | sort > test-coverage-exists.txt
	@# Generate empty test-coverage data for all .go files that have function
	@# definitions, we ignore files that just define types
	@find * -name "*.go" -exec egrep -l '^func ' "{}" \; | sort | sed -e 's|^|napkindrawing.com/pokecache/|; s/$$/ => (0.0%)/' | egrep -v 'napkindrawing\.com/.*((_enum|mock_[^/]*|_test)\.go|/id/[^/]+_id.go)' > test-coverage-all-empty.txt
	@# Join actual existing coverage with empty coverage so we see files having 0% coverage
	@# join produces output with two types of lines:
	@#    - file has a row in both existing & all-empty coverage files:
	@#      - produce six columns, first three from existing, next three from all-empty
	@#    - file has no row in existing, but a row in all-empty:
	@#      - produce three spaces, then the three rows from all-empty
	@join -1 1 -2 1 -o 1.1,1.2,1.3 -a 2 -o 2.1,2.2,2.3 test-coverage-exists.txt test-coverage-all-empty.txt | perl -lape 's/^\s+//; s/ napkindrawing.com.*$$//' | sort > test-coverage.txt
	@# Cleanup intermediate files
	@rm test-coverage-exists.txt test-coverage-all-empty.txt
	@printf " ... done\n"

.PHONY: test-coverage-update-readme
test-coverage-update-readme:
	@printf "Processing test-coverage README updates ..."
	@TEST_COVERAGE="$$(go tool cover -func=test-coverage.out  | grep total: | awk '{print $$3}')" ; \
		sed -i '' -e "s/^Current Test Coverage: .*/Current Test Coverage: $${TEST_COVERAGE}/g" README.md
	@printf '# overridden lints for golangci-lint\n\n```\n' > README.nolint.md
	@git grep --no-color '//nolint:' | egrep '^[^:]+\.go:' | perl -lape 's|^.*//nolint:|//nolint:|; s| //.*$$||; s|//nolint:||; s/,\s*/\n/g' | sort | uniq -c >> README.nolint.md
	@printf '```\n\n' >> README.nolint.md
	@printf '# overridden lints for revive\n\n```\n' >> README.nolint.md
	@git grep --no-color '//revive:disable' | egrep '^[^:]+\.go:' | perl -lape 's|^.*//revive:disable[^:]*:|:|; s/^$$|^([^:].*)$$/(OVERRIDE MODULE NOT SPECIFIED)/g; s/^:(\S+).*$$/$$1/g;' | sort | uniq -c >> README.nolint.md
	@printf '```\n' >> README.nolint.md
	@printf " done\n"



