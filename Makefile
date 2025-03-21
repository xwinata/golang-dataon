.PHONY: clean all init generate

all: build/main

build/main: main.go generated
	@echo "Building..."
	go build -o $@ $<

clean:
	rm -rf specs
# rm -rf mocks

init: clean generate
	go mod tidy
	go mod vendor

generate: generate_specs

generate_specs: openapi.yml
	@echo "Generating files..."
	mkdir specs || true
	oapi-codegen --package specs -generate types,server,spec $< > specs/api.gen.go

# generate_mocks:
# 	@echo "Generating mocks..."
# 	mockery