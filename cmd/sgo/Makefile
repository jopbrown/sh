DIST_EXE := dist/sgo$(shell go env GOEXE)

TEST_GOFILE := $(wildcard _test/*.go)
TEST_CASE := $(addprefix case,$(TEST_GOFILE))
PKG_DIR := pkgs

all: 
	mkdir -p dist
	go build -trimpath -ldflags "-s -w" -o $(DIST_EXE) .
	upx $(DIST_EXE)

generate:
	./gen.sh

clean:
	rm -rf tmp dist $(PKG_DIR)

test: $(TEST_CASE)
	@echo "all pass"

$(TEST_CASE): case% : %
	@echo "test sgo $< ..."
	go run . $< _test/test.txt
	@echo "$< pass"
	@echo
