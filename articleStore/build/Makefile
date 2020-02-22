# Makefile for service

#include custom.mk

# Exporting GO system Variable
export GOPATH=${PWD}/../../../../
GOFMT=gofmt -w
GOBUILD=go build -i -v

# Exporting service specific variables
GOGETFILES= golang.org/x/net github.com/gorilla/mux github.com/sirupsen/logrus github.com/mattn/go-sqlite3
SERVICE_NAME=articleStore
#CLIENT_DIR=${PWD}/../client
SERVER_DIR=server
MODELS_DIR=models
#SERVICE_DIR=genHandler
#ADV_SERVICE_DIR=advHandler
MAIN_SERVICE_DIR=handler
APP_DIR=app

SERVICE_PATH=${PWD}/../../$(SERVICE_NAME)
#COMMON_REST_DIR_PATH=../../common/rest

APPTARGET=$(SERVICE_PATH)/$(APP_DIR)/bin/$(SERVICE_NAME)

#all: client server service 
all: models server service 
	@echo "GOPATH: $(GOPATH)"

get: 
	@echo "go get all missing 3rd party dependencies..."
	go get $(GOGETFILES)
	for i in $(articleStore_SUBDIRS); \
	do \
		echo "go-get for module $$i" ;  \
		(cd $$i ;  \
		make get; \
		) ;  \
	done

#client: 
models: 
	@echo "Building service client module..."
#	$(GOFMT) $(SERVICE_PATH)/$(COMMON_REST_DIR_PATH)/$(SERVICE_NAME)/$(MODELS_DIR)/*.go
	$(GOFMT) $(SERVICE_PATH)/$(MODELS_DIR)/*.go
#	$(GOFMT) $(SERVICE_PATH)/$(COMMON_REST_DIR_PATH)/$(SERVICE_NAME)/$(CLIENT_DIR)/*.go
#	$(GOBUILD) $(SERVICE_PATH)/$(COMMON_REST_DIR_PATH)/$(SERVICE_NAME)/$(MODELS_DIR)/*.go
	$(GOBUILD) $(SERVICE_PATH)/$(MODELS_DIR)/*.go
#	$(GOBUILD) $(SERVICE_PATH)/$(COMMON_REST_DIR_PATH)/$(SERVICE_NAME)/$(CLIENT_DIR)/*.go

modules: 
	@echo "Building all service modules..."
	for i in $(articleStore_SUBDIRS); \
	do \
		echo "Building in module $$i" ; \
		(cd $$i ; \
		make ; \
		) ; \
	done

server: modules
	@echo "Building service server module..."
#	$(GOFMT) $(SERVICE_PATH)/$(COMMON_REST_DIR_PATH)/$(SERVICE_NAME)/$(MODELS_DIR)/*.go
	$(GOFMT) $(SERVICE_PATH)/$(MODELS_DIR)/*.go
#	$(GOFMT) $(SERVICE_PATH)/$(SERVER_DIR)/$(SERVICE_DIR)/*.go
#	$(GOFMT) $(SERVICE_PATH)/$(SERVER_DIR)/$(ADV_SERVICE_DIR)/*.go
	$(GOFMT) $(SERVICE_PATH)/$(SERVER_DIR)/$(MAIN_SERVICE_DIR)/*.go
#	$(GOBUILD) $(SERVICE_PATH)/$(COMMON_REST_DIR_PATH)/$(SERVICE_NAME)/$(MODELS_DIR)/*.go
	$(GOBUILD) $(SERVICE_PATH)/$(MODELS_DIR)/*.go
#	$(GOBUILD) $(SERVICE_PATH)/$(SERVER_DIR)/$(SERVICE_DIR)/*.go
#	$(GOBUILD) $(SERVICE_PATH)/$(SERVER_DIR)/$(ADV_SERVICE_DIR)/*.go
	$(GOBUILD) $(SERVICE_PATH)/$(SERVER_DIR)/$(MAIN_SERVICE_DIR)/*.go

service: server
	@echo "Generating service executable..."
	$(GOFMT) $(SERVICE_PATH)/$(APP_DIR)/*.go
	$(GOBUILD) -o $(APPTARGET) $(SERVICE_PATH)/$(APP_DIR)/*.go

clean: ## Remove previous build
	@echo "Clean previous build"
	rm -fv $(APPTARGET)
	for i in $(articleStore_SUBDIRS); \
	do \
		echo "Clean in module $$i" ; \
		(cd $$i ; \
		make clean; \
		) ; \
	done

