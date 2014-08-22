VERSION = 0.1.0

DEPS = \
	github.com/google/go-github/github \
	github.com/jessevdk/go-flags \
	github.com/robfig/config \
	github.com/wsxiaoys/terminal

all: deps install

install:
	go install -a github.com/pksunkara/hub

deps:
	go get -u $(DEPS)

goxc:
	$(shell echo '{\n "ArtifactsDest": "build",\n "ConfigVersion": "0.9",' > $(GOXC_FILE))
	$(shell echo ' "PackageVersion": "$(VERSION)",\n "TaskSettings": {' >> $(GOXC_FILE))
	$(shell echo '  "bintray": {\n   "apikey": "",\n   "package": "hub",' >> $(GOXC_FILE))
	$(shell echo '   "repository": "utils",\n   "subject": "pksunkara"' >> $(GOXC_FILE))
	$(shell echo '  }\n }\n}' >> $(GOXC_FILE))
	$(GO_XC)

bintray:
	$(GO_XC) bintray
