# GO Commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Others
LINTER=golangci-lint

# Application
APPNAME=custom-wallet
APPPATH = ./cmd/$(APPNAME)
APPMAIN = main.go

# Swagger
SWAG_BUILD = swag
API_HANDLERS = ./pkg/api
API_MAIN =  api.go
API_DOCS = ./api
SWAG_FLAGS = --parseInternal=true --parseDependency=true

all: test build

test:
	@echo "Run tests $(APPNAME)"
	$(GOTEST) ./... -v

build:
	@echo "Build application $(APPNAME)"
	$(GOBUILD) -o $(APPPATH)/$(APPNAME) -v $(APPPATH)/$(APPMAIN)

lint:
	@echo "Run linter  $(LINTER)"
	$(LINTER) run

swag:
	@echo "Create swagger documentation in $(API_DOCS) folder"
	$(SWAG_BUILD) init -g $(API_MAIN) -d $(API_HANDLERS) -o $(API_DOCS) $(SWAG_FLAGS)

depend:
	# godoc
	# @echo "Get and install godoc"
	# $(GOGET) -u golang.org/x/tools/cmd/godoc
	# swag
	# @echo "Get and install swag"
	# $(GOGET) -u github.com/swaggo/swag/cmd/swag
	# gin-swagger
	# @echo "Get and install gin-swagger"
	# $(GOGET) -u github.com/swaggo/gin-swagger
	# $(GOGET) -u github.com/swaggo/files
	# go-migrate
	export GO111MODULE=off && $(GOGET) -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate
