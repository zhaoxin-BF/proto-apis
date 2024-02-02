PROJECTS := $(shell find . -name "buf.gen.yaml" -type f | xargs dirname | xargs basename)
BUILD_PROJECTS :=

ifdef BUILD_PROJECTS
	buildCommand := "./build.sh \"$(BUILD_PROJECTS)\""
else
	buildCommand := "./build.sh"
endif

.PHONY: build
build:
	@bash -c $(buildCommand)

.PHONY: buf
buf:
	buf generate
#	cd proto/cloud && \
#	for file in $$(find . -type f -name "buf.gen.yaml" | tr -s "\n" " ");do \
#		echo "build " $${file}; \
#		path=$${file%/*}; \
#		buf generate --template $$file --path $$path; \
#	done
#
#.PHONY: test
#test:
#	@bash -c $(buildCommand)

.PHONY: build-proto
build-proto:
	cd proto/cloud && \
	for file in $$(find . -type f -name "buf.gen.yaml" | tr -s "\n" " ");do \
		echo "build " $${file}; \
		path=$${file%/*}; \
		buf generate --template $$file --path $$path; \
	done

.PHONY: test-build
test-build:
	cd proto/cloud && buf generate --template ./$(dir)/buf.gen.yaml --path ./$(dir)