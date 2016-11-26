PREFIX =
DESCTDIR=
GOFLAGS=-ldflags "-X \"main.version=`git rev-list HEAD | wc -l`\" -X \"main.commitID=`git rev-list HEAD | head -1`\""
BINDIR=${PREFIX}/build

CONTENT_SRCS = $(wildcard )  main.go

APPS = leetCode
BLDDIR = build
export GOPATH = $(shell pwd)

all: ${APPS}

$(BLDDIR)/%:
	cd ./src/leetCode;  go build ${GOFLAGS} -o ../../build/bin/$(APPS) $(CONTENT_SRCS)

$(BINARIES): %: $(BLDDIR)/%
$(APPS): %: $(BLDDIR)/%


clean:
#	rm -fr $(BLDDIR)

.PHONY: install clean all
.PHONY: $(BINARIES)
.PHONY: $(APPS)
