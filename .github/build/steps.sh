#!/usr/bin/env bash
# The build steps of every repo (TOOL_23 to TOOL_26), run by .github/workflows/build.yml in the root
# of the repo; the same script runs by hand.
#
#	steps.sh build|lint|test|cover|debt
#	steps.sh sbom [name=Dockerfile]...
#
# A Go module is every directory with a go.mod, a JavaScript package every directory with a
# package-lock.json - neither below node_modules nor below a directory whose name starts with a dot.
# build, lint and test exit with 1 where a step fails. cover and debt report and exit with 0; they
# exit otherwise only where the measuring itself did not happen. sbom builds each image named from
# its Dockerfile, and exits with 1 where sbom.md does not list the packages delivered.
#
# It needs go, golangci-lint v2 and deadcode on the PATH, and node with npm where there is a
# JavaScript package.
set -uo pipefail

#	the root of specs-and-deploy; .github/build holds the configurations of the linters
steps="$(cd "$(dirname "$0")/../.." && pwd)"

#	@implements SR_51@1.1.1
modules() {
	find . -name go.mod -not -path '*/node_modules/*' -not -path '*/.*/*' | sed 's#/go.mod$##' | sort
}

#	@implements SR_52@1.1.1
packages() {
	find . -name package-lock.json -not -path '*/node_modules/*' -not -path '*/.*/*' | sed 's#/package-lock.json$##' | sort
}

failed=0

#	run <dir> <command>...: run a command in a directory, and note when it fails
run() {
	local dir="$1"
	shift
	echo "== $dir: $*"
	if ! (cd "$dir" && "$@"); then
		echo "FAIL $dir: $*"
		failed=1
	fi
}

#	@implements SR_51@1.1.1, SR_52@1.1.1
build() {
	for m in $(modules); do
		run "$m" go build ./...
	done
	for p in $(packages); do
		run "$p" npm ci --no-audit --no-fund
		run "$p" npm run build
	done
}

#	the packages eslint.config.js imports, installed beside it once
eslint_install() {
	[ -x "$steps/node_modules/.bin/eslint" ] && return
	(cd "$steps" && npm install --no-save --no-audit --no-fund \
		eslint@10.14.0 @eslint/js@10.0.1 globals@17.13.0) || exit 2
}

#	@implements SR_53@1.1.1, SR_54@1.1.1
lint() {
	for m in $(modules); do
		run "$m" golangci-lint run -c "$steps/.github/build/.golangci.yml" ./...
		run "$m" golangci-lint run -c "$steps/.github/build/.golangci.yml" --build-tags integration ./...
	done
	if [ -n "$(packages)" ]; then
		eslint_install
	fi
	for p in $(packages); do
		run "$p" "$steps/node_modules/.bin/eslint" -c "$steps/.github/build/eslint.config.js" .
	done
}

#	@implements SR_55@1.1.1, SR_56@1.1.1
test() {
	for m in $(modules); do
		run "$m" go test -race -timeout 30m ./...
		run "$m" go test -race -timeout 30m -tags integration ./...
	done
	for p in $(packages); do
		run "$p" npm ci --no-audit --no-fund
		run "$p" npm test
	done
}

#	@implements SR_57@1.1.1
cover() {
	local profile
	profile="$(mktemp)"
	for m in $(modules); do
		echo "== $m: test coverage of MT and SwIT, with -coverpkg=./..."
		if ! (cd "$m" && go test -tags integration -coverpkg=./... -coverprofile="$profile" ./... >/dev/null); then
			echo "a test failed; the coverage is of the tests that ran"
		fi
		(cd "$m" && go tool cover -func="$profile") || exit 2
	done
	for p in $(packages); do
		echo "== $p: test coverage"
		if (cd "$p" && npm pkg get scripts.coverage | grep -qv '{}'); then
			(cd "$p" && npm ci --no-audit --no-fund && npm run coverage) || exit 2
		else
			echo "no script coverage in package.json"
		fi
	done
	rm -f "$profile"
}

#	@implements SR_58@1.1.1
debt() {
	for m in $(modules); do
		echo "== $m: what nothing uses, and a cyclomatic complexity of 15 and more"
		(cd "$m" && golangci-lint run -c "$steps/.github/build/debt.golangci.yml" --tests=false ./...)
		[ $? -le 1 ] || exit 2
		echo "== $m: what no program reaches (deadcode)"
		(cd "$m" && deadcode ./...) || exit 2
	done
}

#	@implements SR_67@1.1.1
sbom() {
	local bin names=()
	bin="$(mktemp -d)/sbom"
	(cd "$steps/tools" && GOTOOLCHAIN=auto go build -o "$bin" ./sbom/cmd/sbom) || exit 2
	for image in "$@"; do
		run . docker build -t "${image%%=*}" -f "${image#*=}" .
		names+=(-image "${image%%=*}")
	done
	[ "$failed" = 0 ] || exit 2
	run . "$bin" -check ${names[@]+"${names[@]}"}
}

case "${1:-}" in
build | lint | test)
	"$1"
	exit "$failed"
	;;
sbom)
	shift
	sbom "$@"
	exit "$failed"
	;;
cover | debt)
	"$1"
	;;
*)
	echo "usage: steps.sh build|lint|test|cover|debt|sbom" >&2
	exit 2
	;;
esac
