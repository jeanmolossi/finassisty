DIAGRAMS_DIR := docs/architecture

PUML_FILES := $(wildcard $(DIAGRAMS_DIR)/*.puml)
PNG_FILES  := $(patsubst %.puml,%.png,$(PUML_FILES))

PLANTUML := docker run --rm \
						-v "$(CURDIR)":/workspace \
						-w /workspace \
						plantuml/plantuml

PLANTUML_ARGS ?= -tpng

.PHONY: docs clean
docs: $(PNG_FILES)
	@echo "✅ Diagramas atualizados em $(DIAGRAMS_DIR)"

$(DIAGRAMS_DIR)/%.png: $(DIAGRAMS_DIR)/%.puml
	@echo "🖼️  Gerando $@"
	@$(PLANTUML) $(PLANTUML_ARGS) "$<"

clean:
	@rm -f $(PNG_FILES)
	@echo "🧹  Diagramas removidos"


.PHONY: pwa-deps api-deps dev
install-deps: pwa-deps api-deps
	@echo "✅ Dependencias instaladas"

pwa-deps:
	cd app && pnpm install

api-deps:
	cd server && go mod download

dev:
	pnpm --dir app dev & \
		cd server && go run ./cmd/api

install-go-lint:
	@if ! command golangci-lint -v > /dev/null; then \
		read -p "Go's linter is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.1.6; \
			if [ ! -x "$$(command golangci-lint -v)" ]; then \
				echo "Go linter installation failed. Exiting..."; \
				exit 1; \
			fi; \
		fi; \
	fi

install-js-lint:

lint: install-go-lint install-js-lint
	@golangci-lint run ./...
	@pnpm --dir app run tsc --noEmit
