# Auto-generated by repobuild, do not modify directly.

.gen-src/rpcz:
	mkdir -p .gen-src; if [[ ! -a . ]]; then mkdir -p .; fi; ln -f -s .. .gen-src/rpcz

.gen-src/rpcz/.dummy: .gen-src/rpcz
	if [[ ! -a .gen-src/rpcz/.dummy ]]; then touch .gen-src/rpcz/.dummy; fi

clean:
	rm -f .gen-src/rpcz/.dummy
	rm -f .gen-src/rpcz
	rm -rf .gen-obj
	rm -rf .gen-files
	rm -rf .gen-src

all:

.PHONY: clean all

.DEFAULT_GOAL=all

