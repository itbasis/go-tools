#!/bin/bash

builderGoPath=builder/main.go
builderOutput=.itbasis/builder

echo "os: ${OSTYPE}"
if [[ "$OSTYPE" == "msys" ]]; then
	builderOutput="${builderOutput}.exe"
fi

go run ${builderGoPath} \
	build \
	--debug \
	--output="${builderOutput}" \
	${builderGoPath}
