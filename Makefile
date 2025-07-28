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

install-reflex:
	@if ! command -v reflex -h > /dev/null; then \
		read -p "Go's reflex is not installed. It's needed to hot reload. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/cespare/reflex@latest; \
			if [ ! -x "$$(which reflex -h)" ]; then \
				echo "Go's reflex installation failed. Exiting..."; \
				exit 1; \
			fi; \
		fi; \
	fi

dev: install-reflex
	pnpm --dir app dev & \
		cd server && reflex -r '\.go$$' -s -- sh -c "go run ./cmd/api"

monitoring:
	@docker compose -f ./infra/monitoring/docker-compose.yml up -d; \
		bash -c "trap 'trap - SIGINT SIGTERM ERR; \
			docker compose -f ./infra/monitoring/docker-compose.yml down -v; \
			exit 1' SIGINT SIGTERM ERR; \
			docker logs -f monitoring-alloy-1"

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
