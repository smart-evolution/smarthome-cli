GOCMD=go
GOLINT=golint
GOFMT=gofmt

.PHONY: install
install:
	$(GOCMD) get golang.org/x/lint/golint
	$(GOCMD) get golang.org/x/lint/golint

.PHONY: deploy
deploy:
	./scripts/install/install.sh

.PHONY: version
version:
	git tag $(V)
	./scripts/changelog.sh
	git add ./docs/changelogs/CHANGELOG_$(V).md
	git commit --allow-empty -m "Build $(V)"
	git tag --delete $(V)
	git tag $(V)

.PHONY: lint
lint:
	$(GOLINT) ./...
	$(GOCMD) vet ./...

.PHONY: fix
fix:
	$(GOFMT) -w .

.PHONY: help
help:
	@echo  'Available tasks:'
	@echo  '* Installation:'
	@echo  '- install         - Phony task that installs script globally'
	@echo  '- deploy          - Deploys script'
	@echo  ''
	@echo  '* Quality:'
	@echo  '- lint            - Phony task that runs all linting tasks'
	@echo  ''
	@echo  '* Release:'
	@echo  '- version         - Phony task. Creates changelog from latest'
	@echo  '                    git tag till the latest commit. Creates commit'
	@echo  '                    with given version (as argument) and tags'
	@echo  '                    this commit with this version. Version has to'
	@echo  '                    be passed as `V` argument with ex. `v0.0.0`'
	@echo  '                    format'.
	@echo  '                    Example: $ make version V=v0.0.0.
	@echo  ''

