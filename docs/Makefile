# Note to developers/testers

# in order to run this make file correctly, you must set the following env vars
## To build in a test repo:
## To simulate a CI-driven release
# TAGGED_VERSION

GCLOUD_PROJECT_ID ?= solo-test-236622


#----------------------------------------------------------------------------------
# Base
#----------------------------------------------------------------------------------
# determine if we are doing a release (same logic as with repo root Makefile)
# RELEASE := "true"
# ifeq ($(TAGGED_VERSION),)
# 	TAGGED_VERSION := v
# 	RELEASE := "false"
# endif

# remove the "v" prefix
VERSION ?= $(shell echo $(TAGGED_VERSION) | cut -c 2-)

# Passed by cloudbuild
GCLOUD_PROJECT_ID := $(GCLOUD_PROJECT_ID)
BUILD_ID := $(BUILD_ID)

TEST_IMAGE_TAG := test-$(BUILD_ID)
GCR_REPO_PREFIX := gcr.io/$(GCLOUD_PROJECT_ID)

# this allows us to host docs.solo.io/gloo/vX.Y.Z or docs.solo.io/supergloo/vA.B.C
PRODUCT_SCOPE := gloo
DOCS_VERSION := latest

#----------------------------------------------------------------------------------
# Docs
#----------------------------------------------------------------------------------

.PHONY: site-common
site-common: clean
	# if [ ! -d themes/hugo-theme-soloio ]; then git clone https://github.com/solo-io/hugo-theme-soloio themes/hugo-theme-soloio && git --git-dir=themes/hugo-theme-soloio/.git checkout b49aed308e7c03edd2640fe52afd745d37fa0cbc; fi
	cp pending-theme/shortcodes/protobuf.html themes/hugo-theme-soloio/layouts/shortcodes/
	go run cmd/generate_changelog_doc.go gen-changelog-md gloo > content/changelog/gloo-changelog.sk.md
	go run cmd/generate_changelog_doc.go gen-changelog-md glooe > content/changelog/glooe-changelog.sk.md

.PHONY: site-test
site-test: site-common
	go run cmd/generate_changelog_doc.go gen-version-scope-data --no-scope
	hugo --config docs.toml

.PHONY: site-release
site-release: site-common
	go run cmd/generate_changelog_doc.go gen-version-scope-data --product gloo --call-latest
	hugo --config docs.toml
	mv site site-latest
	go run cmd/generate_changelog_doc.go gen-version-scope-data --product gloo --version $(VERSION)
	hugo --config docs.toml
	mv site site-versioned

# .PHONY: deploy-site
# deploy-site: site
# 	firebase deploy --only hosting:gloo-docs

.PHONY: serve-site
serve-site: site-test
	hugo --config docs.toml server -D

.PHONY: clean
clean:
	rm -fr ./site ./resources ./site-latest ./site-versioned

# Uses https://github.com/gjtorikian/html-proofer
# Does not require running site; just make sure you generate the site and then run it
# Install with gem install html-proofer
# Another option we could use is wget: https://www.digitalocean.com/community/tutorials/how-to-find-broken-links-on-your-website-using-wget-on-debian-7
.PHONY: check-links
check-links:
	htmlproofer ./site/ --empty-alt-ignore  --allow-hash-href --alt-ignore "/img/Gloo-01.png" --url-ignore "/localhost/,/github.com/solo-io/solo-projects/,/developers.google.com/,/getgrav.org/,/github.com/solo-io/gloo/projects/,/developer.mozilla.org/"

# If on fedora, run
#     sudo dnf -y install gcc ruby-devel rubygems zlib-devel
# to install html-proofer deps (only works with gcc, not clang!)
install-tools:
	gem install html-proofer

.PHONY: docker-push-docs
docker-push-docs: site-release manifest
	docker build \
		--build-arg VERSION=latest \
		--build-arg PRODUCT_SCOPE=$(PRODUCT_SCOPE) \
		--build-arg FROM_DIR=./site-latest \
		-t $(GCR_REPO_PREFIX)/gloo-docs:latest .
	docker push $(GCR_REPO_PREFIX)/gloo-docs:latest
	docker build \
		--build-arg VERSION=$(VERSION) \
		--build-arg PRODUCT_SCOPE=$(PRODUCT_SCOPE) \
		--build-arg FROM_DIR=./site-versioned \
		-t $(GCR_REPO_PREFIX)/gloo-docs:$(VERSION) .
	docker push $(GCR_REPO_PREFIX)/gloo-docs:$(VERSION)

.PHONY: manifest
manifest:
	helm template install/helm \
		--set docsVersion=$(VERSION) \
		--set image.tag=$(VERSION) \
		> install/manifest-versioned.yaml
	helm template install/helm \
		--set docsVersion=latest \
		--set image.tag=latest \
		> install/manifest-latest.yaml

