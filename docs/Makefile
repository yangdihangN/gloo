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
RELEASE := "true"
ifeq ($(TAGGED_VERSION),)
	TAGGED_VERSION := v
	RELEASE := "false"
endif

# remove the "v" prefix
VERSION ?= $(shell echo $(TAGGED_VERSION) | cut -c 2-)

## Passed by cloudbuild
#GCLOUD_PROJECT_ID := $(GCLOUD_PROJECT_ID)
#BUILD_ID := $(BUILD_ID)
#
#TEST_IMAGE_TAG := test-$(BUILD_ID)
GCR_REPO_PREFIX := gcr.io/$(GCLOUD_PROJECT_ID)



#----------------------------------------------------------------------------------
# Docs
#----------------------------------------------------------------------------------

site: clean pass-version-to-hugo
	# if [ ! -d themes/hugo-theme-soloio ]; then git clone https://github.com/solo-io/hugo-theme-soloio themes/hugo-theme-soloio && git --git-dir=themes/hugo-theme-soloio/.git checkout b49aed308e7c03edd2640fe52afd745d37fa0cbc; fi
	cp pending-theme/shortcodes/protobuf.html themes/hugo-theme-soloio/layouts/shortcodes/
	hugo --config docs.toml

.PHONY: deploy-site
deploy-site: site
	firebase deploy --only hosting:gloo-docs

.PHONY: serve-site
serve-site: site
	hugo --config docs.toml server -D

.PHONY: clean
clean:
	rm -fr ./site ./resources

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

.PHONY: docker-push
docker-push: site
	docker build \
		--build-arg VERSION=$(VERSION) \
		-t gcr.io/solo-test-236622/gloo-docs:$(VERSION) .
	docker push $(GCR_REPO_PREFIX)/gloo-docs:$(VERSION)

.PHONY: manifest
manifest:
	helm template install/helm \
		--set docsVersion=$(VERSION) \
		--set image.tag=$(VERSION) \
		> install/manifest.yaml

# do something more sophisticated as needed
# for now, we only use a single value in the Solo.yaml file
# during local development, the value must be an empty string to correctly resolve the relative links
# during a release or other version-scoped deployment, the version must be prefixed with a "/" to correctly match the nginx path
.PHONY: pass-version-to-hugo
pass-version-to-hugo:
	echo DocsVersion: > data/Solo.yaml
ifeq ($(RELEASE),"true")
	echo DocsVersion: "/$(VERSION)" > data/Solo.yaml
endif
