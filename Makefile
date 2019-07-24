APPNAME := $(shell basename `pwd`)
VERSION := $(shell grep Version version.go | grep -Eo "[0-9\.]+")
SRCS := $(shell find . -name "*.go" -type f )
LDFLAGS := -ldflags="-s -w \
	-extldflags \"-static\""
XBUILD_TARGETS := \
	-os="windows linux darwin" \
	-arch="386 amd64" 
DIST_DIR := dist
README := README.*
EXTERNAL_TOOLS := \
	github.com/mitchellh/gox

help: ## ドキュメントのヘルプを表示する。
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: $(SRCS) ## ビルド
	go build $(LDFLAGS) -o bin/$(APPNAME) .

install: build ## インストール
	go install

xbuild: $(SRCS) bootstrap ## クロスコンパイル
	gox $(LDFLAGS) $(XBUILD_TARGETS) --output "$(DIST_DIR)/{{.Dir}}$(VERSION)_{{.OS}}_{{.Arch}}/{{.Dir}}"

archive: xbuild ## クロスコンパイルしたバイナリとREADMEを圧縮する
	find $(DIST_DIR)/ -mindepth 1 -maxdepth 1 -a -type d \
		| while read -r d; \
		do \
			cp $(README) $$d/ ; \
			cp LICENSE $$d/ ; \
		done
	cd $(DIST_DIR) && \
		find . -maxdepth 1 -mindepth 1 -a -type d  \
		| while read -r d; \
		do \
			../archive.sh $$d; \
		done

test: ## テストコードを実行する
	go test -cover ./...

clean: ## バイナリ、配布物ディレクトリを削除する
	-rm -rf bin
	-rm -rf $(DIST_DIR)

bootstrap: ## 外部ツールをインストールする
	for t in $(EXTERNAL_TOOLS); do \
		echo "Installing $$t ..." ; \
		GO111MODULE=off go get $$t ; \
	done

.PHONY: help build install xbuild archive release lint test clean bootstrap docker-build docker-push
