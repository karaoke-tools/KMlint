prefix = /usr/local
exec_prefix = $(prefix)
bindir = $(exec_prefix)/bin
BASHCOMPLETIONSDIR = $(exec_prefix)/share/bash-completion/completions


RM = rm -f
INSTALL = install -D
MKDIRP = mkdir -p

.PHONY: install uninstall build clean default
default: build
build:
	CGO_ENABLED=0 go build -trimpath
clean:
	go clean
reinstall: uninstall install
install:
	$(INSTALL) kmlint $(DESTDIR)$(bindir)/kmlint
	$(MKDIRP) $(DESTDIR)$(BASHCOMPLETIONSDIR)
	$(DESTDIR)$(bindir)/kmlint completion bash > $(DESTDIR)$(BASHCOMPLETIONSDIR)/kmlint
	@echo "================================="
	@echo ">> Now run the following command:"
	@echo "\tsource $(DESTDIR)$(BASHCOMPLETIONSDIR)/kmlint"
	@echo "================================="
uninstall:
	$(RM) $(DESTDIR)$(bindir)/kmlint
	$(RM) $(DESTDIR)$(BASHCOMPLETIONSDIR)/kmlint
