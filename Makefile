GO=go
OUTPUT=./build/gitview
SRC=./src/gitview

all: build

build:
	$(GO) build -o $(OUTPUT) $(SRC)

run:
	$(GO) run $(SRC)