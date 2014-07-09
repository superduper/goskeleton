PROJECT_NAME=skeleton

compile:
	gvp in go build $(PROJECT_NAME)

deps:
	gvp in gpm 
	gvp in gpm git

build: deps compile
