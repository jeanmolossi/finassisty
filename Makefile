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
