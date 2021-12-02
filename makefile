YEAR ?= $(shell date +"%Y")
DAY ?= $(shell date +"%-d")

run:
	go run ./$(YEAR)/day$(DAY)
