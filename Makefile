APPNAME = s
OUTDIR = pkg

# Allow user to override cross compilation scope
OSARCH ?= darwin/amd64 darwin/arm64 dragonfly/amd64 freebsd/386 freebsd/amd64 freebsd/arm linux/386 linux/amd64 linux/arm netbsd/386 netbsd/amd64 netbsd/arm openbsd/386 openbsd/amd64 plan9/386 plan9/amd64 solaris/amd64 windows/386 windows/amd64

all:
	go build .

compile:
	@for osarch in $(OSARCH) ; do \
		os=$${osarch%/*} ; arch=$${osarch#*/} ; ext="" ;\
		if [ "$$os" = "windows" ]; then ext=".exe" ; fi ;\
		dir="$(APPNAME)-$${os}_$${arch}" ;\
		echo "building $$osarch" ;\
		CGO_ENABLED=0 GOOS=$$os GOARCH=$$arch go build -o "$(OUTDIR)/$$dir/$(APPNAME)$$ext" . || exit 1 ;\
		(cp README.md $(OUTDIR)/$$dir/README.md) ;\
		(cp LICENSE $(OUTDIR)/$$dir/LICENSE) ;\
		(cd $(OUTDIR) && zip -q $$dir.zip -r $$dir) ;\
		echo "make $(OUTDIR)/$$dir.zip" ;\
	done

install:
	go install .

uninstall:
	go clean -i

docker:
	docker build -t $(APPNAME) .
