YEAR ?= $(shell date +"%Y")
DAY ?= $(shell date +"%-d")

run:
	go run ./$(YEAR)/day$(DAY)

build_go:
	.bin/build -d $(DAY) -y $(YEAR) -l go

build_rb:
	.bin/build -d $(DAY) -y $(YEAR) -l rb
