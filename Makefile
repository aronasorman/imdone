release:
	go build
	rm ~/bin/imdone  # mac wants you to delete the old one first for signature reasons
	cp ./imdone ~/bin/imdone

rust_release:
	cargo build --release
	rm ~/bin/imdone  # mac wants you to delete the old one first for signature reasons
	cp target/release/imdone ~/bin/imdone