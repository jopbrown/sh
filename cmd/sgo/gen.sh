#!/usr/bin/env bash

set -e
set -x

pkgname=sgolib
pkg_dir=internal/${pkgname}

rm -rf ${pkg_dir}
mkdir -p ${pkg_dir}

pushd ${pkg_dir}

yaegi extract -name ${pkgname} `cat ../../extract.txt | grep -v '^#'`

sed -i \
-e 's;logFatal;log.Fatal;' \
-e 's;logFatalf;log.Fatalf;' \
-e 's;logFatalln;log.Fatalln;' \
-e 's;logLogger;log.Logger;' \
-e 's;logNew;log.New;' \
log.go

sed -i \
-e 's;osExit;os.Exit;' \
-e 's;osFindProcess;os.FindProcess;' \
os.go

cat > ${pkgname}.go <<EOF
package ${pkgname}

import "reflect"

var Symbols = make(map[string]map[string]reflect.Value, `ls *.go | wc -l`)
EOF

popd