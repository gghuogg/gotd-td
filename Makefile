test:
	@./go.test.sh
.PHONY: test

coverage:
	@./go.coverage.sh
.PHONY: coverage

generate:
	go generate
	go generate ./...
.PHONY: generate

download_schema:
	go run ./cmd/dltl -base https://raw.githubusercontent.com/tdlib/td -branch master -dir td/generate/scheme -f telegram_api.tl -o _schema/tdlib.tl
	go run ./cmd/dltl -base https://raw.githubusercontent.com/telegramdesktop/tdesktop -branch dev -dir Telegram/SourceFiles/mtproto/scheme -f api.tl -o _schema/tdesktop.tl
	go run ./cmd/dltl -base https://raw.githubusercontent.com/telegramdesktop/tdesktop -branch dev -dir Telegram/SourceFiles/mtproto/scheme -f layer.tl -o _schema/layer.tl
	go run ./cmd/dltl -base https://raw.githubusercontent.com/telegramdesktop/tdesktop -branch dev -dir Telegram/SourceFiles/mtproto/scheme -f api.tl -merge _schema/legacy.tl,_schema/layer.tl -o _schema/telegram.tl
.PHONY: download_schema

download_public_keys:
	go run ./cmd/dlkey -o internal/mtproto/_data/public_keys.pem
.PHONY: download_public_keys

download_e2e_schema:
	go run ./cmd/dltl -f secret_api.tl -o _schema/encrypted.tl
.PHONY: download_e2e_schema

download_tdlib_schema:
	go run ./cmd/dltl -f td_api.tl -o _schema/tdapi.tl
.PHONY: download_tdlib_schema

check_generated: generate
	git diff --exit-code
.PHONY: check_generated

fuzz_telegram:
	go run -modfile=_tools/go.mod github.com/dvyukov/go-fuzz/go-fuzz -bin mtproto/telegram-fuzz.zip -workdir _fuzz/handle_message
.PHONY: fuzz_telegram

fuzz_telegram_build:
	cd mtproto && go run -modfile=../_tools/go.mod github.com/dvyukov/go-fuzz/go-fuzz-build -func FuzzHandleMessage -tags fuzz -o telegram-fuzz.zip
.PHONY: fuzz_telegram_build

fuzz_telegram_clear:
	rm -f _fuzz/handle_message/crashers/*
	rm -f _fuzz/handle_message/suppressions/*
.PHONY: fuzz_telegram_clear

fuzz_rsa:
	go run -modfile=_tools/go.mod github.com/dvyukov/go-fuzz/go-fuzz -bin internal/crypto/rsa-fuzz.zip -workdir _fuzz/rsa
.PHONY: fuzz_rsa

fuzz_rsa_build:
	cd internal/crypto && go run -modfile=../../_tools/go.mod github.com/dvyukov/go-fuzz/go-fuzz-build -func FuzzRSA -tags fuzz -o rsa-fuzz.zip
.PHONY: fuzz_rsa_build

fuzz_rsa_clear:
	rm -f _fuzz/rsa/crashers/*
	rm -f _fuzz/rsa/suppressions/*
.PHONY: fuzz_rsa_clear

fuzz_flow:
	go run -modfile=_tools/go.mod github.com/dvyukov/go-fuzz/go-fuzz -bin internal/exchange/flow-fuzz.zip -workdir _fuzz/flow
.PHONY: fuzz_flow

fuzz_flow_build:
	cd internal/exchange && go run -modfile=../../_tools/go.mod github.com/dvyukov/go-fuzz/go-fuzz-build -func FuzzFlow -tags fuzz -o flow-fuzz.zip
.PHONY: fuzz_flow_build

fuzz_flow_clear:
	rm -f _fuzz/flow/crashers/*
	rm -f _fuzz/flow/suppressions/*
.PHONY: fuzz_flow_clear
