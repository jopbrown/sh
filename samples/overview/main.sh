#!/usr/bin/env bash

set -e
set -x

echo "$#"
echo "$0"
echo "$1"

for arg in $@; do
	echo $arg
done

pwd

echo "$PATH"

rm -rf tmp*

ls
ls *.go

mkdir -p tmp tmp2
cp *.go tmp/.

cp -r tmp tmp3
mv tmp/*.go tmp2/.

pushd tmp2
touch 1.txt 2.txt

if [[ -e 1.txt ]]; then
    echo "exist"
fi

if [[ -f 1.txt ]]; then
    echo "file exist"
fi

if [[ -d 1.txt ]]; then
    echo "dir exist"
fi

ls *.txt
rm *.txt *.go

popd

pushd tmp3

grep 'func' main.go
cat *.go | grep 'func'

sed -e 's/func/xxxx/g' main.go > main2.txt
cat main.go | sed -e 's/func/ssss/g' >> main2.txt

tar zcvf xxx.tar.gz *
if [[ $? != 0 ]]; then
    exit -1
fi

echo "xxxxx" | base64

popd
