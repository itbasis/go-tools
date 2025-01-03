#!/bin/bash

builderGoPath=builder/main.go

go run ${builderGoPath} \
	build \
	--debug \
	--build-version="${version}" \
	--build-os=windows \
	--build-arch=amd64 \
	--output=".itbasis/builder.exe" \
	${builderGoPath}

go run ${builderGoPath} \
	build \
	--debug \
	--build-version="${version}" \
	--build-os=linux \
	--build-arch=amd64 \
	--output=".itbasis/builder" \
	${builderGoPath}
