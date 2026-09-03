prefix = /usr/local
exec_prefix = $(prefix)
bindir = $(exec_prefix)/bin
BASHCOMPLETIONSDIR = $(exec_prefix)/share/bash-completion/completions
FISHCOMPLETIONSDIR = $(shell pkg-config fish --variable completionsdir)
ZSHCOMPLETIONSDIR = $(exec_prefix)/share/zsh/site-functions


RM = rm -f
INSTALL = install -D
MKDIRP = mkdir -p

.PHONY: install uninstall build clean default
default: build
build:
	@CGO_ENABLED=0 go build -trimpath
clean:
	@go clean
reinstall: uninstall install
install:
	@$(INSTALL) kmlint $(DESTDIR)$(bindir)/kmlint
	@$(MKDIRP) $(DESTDIR)$(BASHCOMPLETIONSDIR)
	@$(DESTDIR)$(bindir)/kmlint completion bash > $(DESTDIR)$(BASHCOMPLETIONSDIR)/kmlint
	@$(MKDIRP) $(DESTDIR)$(FISHCOMPLETIONSDIR)
	@$(DESTDIR)$(bindir)/kmlint completion fish > $(DESTDIR)$(FISHCOMPLETIONSDIR)/kmlint.fish
	@$(MKDIRP) $(DESTDIR)$(ZSHCOMPLETIONSDIR)
	@$(DESTDIR)$(bindir)/kmlint completion zsh > $(DESTDIR)$(ZSHCOMPLETIONSDIR)/_kmlint
uninstall:
	@$(RM) $(DESTDIR)$(bindir)/kmlint
	@$(RM) $(DESTDIR)$(BASHCOMPLETIONSDIR)/kmlint
	@$(RM) $(DESTDIR)$(FISHCOMPLETIONSDIR)/kmlint.fish
	@$(RM) $(DESTDIR)$(ZSHCOMPLETIONSDIR)/_kmlint
