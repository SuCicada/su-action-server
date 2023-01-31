include .env
export

SOURCES ?= $(shell find . -name "*.go" -type f)

TARGET=dist/su-action-server
all: deploy

DEPLOY_PATH ?= ~/APP/su-action-server/

info:
	@echo $(DEPLOY_HOST) $(DEPLOY_PATH)

deploy: build upload

# ========== build ==========
build: $(TARGET)
$(TARGET): $(SOURCES)
	make fmt
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -ldflags "-s -w" -o $(TARGET)

fmt: $(SOURCES)
	@hash gofumpt > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go install mvdan.cc/gofumpt; \
	fi
	gofumpt -l -w $(SOURCES)


# ========== upload ==========

define upload
	scp $(1) $(DEPLOY_HOST):$(DEPLOY_PATH)
endef
define command
	ssh $(DEPLOY_HOST) $(1)
endef


upload:
	$(call command, "supervisorctl stop su-action-server")
	$(call upload, $(TARGET))
	$(call upload, notification.gohtml)
	$(call command, "supervisorctl start su-action-server")
	$(call command, "supervisorctl tail su-action-server")



debug:
	./act-debug.sh
