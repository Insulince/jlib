dir ?= .

# Recipe surface lays out the API surface for the optional directory, dir. If dir is not set, all packages are listed.
surface:
	@for pkg in $$(go list ./pkg/${dir}/...); do \
		go doc $$pkg; \
	done

compile:
	@go build ./pkg/${dir}/...
