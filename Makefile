APPNAME = s
OUTDIR = pkg

# Allow user to override cross compilation scope
OSARCH ?= darwin/386 darwin/amd64 dragonfly/amd64 freebsd/386 freebsd/amd64 freebsd/arm linux/386 linux/amd64 linux/arm netbsd/386 netbsd/amd64 netbsd/arm openbsd/386 openbsd/amd64 plan9/386 plan9/amd64 solaris/amd64 windows/386 windows/amd64
DIRS ?= darwin_386 darwin_amd64 dragonfly_amd64 freebsd_386 freebsd_amd64 freebsd_arm linux_386 linux_amd64 linux_arm netbsd_386 netbsd_amd64 netbsd_arm openbsd_386 openbsd_amd64 plan9_386 plan9_amd64 solaris_amd64 windows_386 windows_amd64

all:
	go build .

compile:
	gox -osarch="$(OSARCH)" -output "$(OUTDIR)/$(APPNAME)-{{.OS}}_{{.Arch}}/$(APPNAME)"
	@for dir in $(DIRS) ; do \
		(cp README.md $(OUTDIR)/$(APPNAME)-$$dir/README.md) ;\
		(cp LICENSE $(OUTDIR)/$(APPNAME)-$$dir/LICENSE) ;\
		(cp -r autocomplete $(OUTDIR)/$(APPNAME)-$$dir/autocomplete) ;\
		(cd $(OUTDIR) && zip -q $(APPNAME)-$$dir.zip -r $(APPNAME)-$$dir) ;\
		echo "make $(OUTDIR)/$(APPNAME)-$$dir.zip" ;\
	done

install:
	go install .

uninstall:
	go clean -i

docker:
	docker build -t $(APPNAME) .
