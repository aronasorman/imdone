release:
	cargo build --release
	rm ~/bin/imdone  # mac wants you to delete the old one first for signature reasons
	cp target/release/imdone ~/bin/imdone