APP_NAME := wareflow
APP_DIR := $(CURDIR)/cmd/
CURRENT_CONFIG := local.yaml
CONFIG_DIR := ./config/


.PHONY: all
all: build

.PHONY: build
build:
	@echo "Сборка приложения: $(APP_NAME)"
	export CONFIG_WARE_FLOW="$(CONFIG_DIR)$(CURRENT_CONFIG)"
	@export CONFIG_WARE_FLOW="$(CONFIG_DIR)$(CURRENT_CONFIG)" && go build -C $(APP_DIR)/$(APP_NAME) -o $(APP_NAME)

.PHONY: run
run:
	@echo "Запуск приложения: $(APP_NAME)"
	@export CONFIG_WARE_FLOW="$(CONFIG_DIR)$(CURRENT_CONFIG)" && go run $(APP_DIR)$(APP_NAME)/main.go

.PHONY: clean
clean:
	@echo "Очищаю рабочую директорию..."
	rm -f $(APP_DIR)/$(APP_NAME)/$(APP_NAME)
	rm -rf ./coverage.out