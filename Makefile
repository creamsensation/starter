dev: dev-starter dev-assets

dev-starter:
	reflex --all=false -r '(\.go)' -s go run app/starter/*.go

dev-assets:
	cd assets && $(MAKE) dev