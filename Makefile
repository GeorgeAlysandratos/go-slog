.PHONY: build clean

build:
	@go test

clean:
	@rm -rf slog_test_*.txt
	@rm -rf go-slog.test_*.txt
