ifneq ("$(wildcard .env)","")
	include .env
	export
endif

SOURCES ?= $(shell find . -name "*.go" -type f)

TARGET=dist/su-action-server
all: build

DEPLOY_PATH ?= ~/APP/su-action-server/

info:
	@echo $(DEPLOY_HOST) $(DEPLOY_PATH)
	@echo $(WEBHOOK_URL)

deploy: build upload

# ========== build ==========
build: $(TARGET)
$(TARGET): $(SOURCES)
	make fmt
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -ldflags "-s -w" -o $(TARGET)
	cp notification.gohtml dist/


fmt: $(SOURCES)
	@hash gofumpt > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go get -u mvdan.cc/gofumpt; \
		go install mvdan.cc/gofumpt; \
	fi
	gofumpt -l -w $(SOURCES)


# ========== upload ==========

define upload
	rsync -avz \
		--rsh="ssh -o StrictHostKeyChecking=no" \
		$(1) \
		${DEPLOY_HOST}:$(DEPLOY_PATH)/$(patsubst %,%,$(2))
endef
define command
	ssh -o StrictHostKeyChecking=no $(DEPLOY_HOST) $(1)
endef


upload:
	$(call command, "supervisorctl stop su-action-server")
	$(call command, "mkdir -p $(DEPLOY_PATH)")
	$(call upload, dist/)
	$(call upload, .env.pro, .env)
	$(call command, "supervisorctl start su-action-server")
	$(call command, "supervisorctl tail su-action-server")



act-debug:
	./act-debug.sh
