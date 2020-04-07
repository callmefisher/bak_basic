#!/bin/zsh

#str=`cat LICENSE`
#str=`awk '{print "//"$0}'  LICENSE`
#echo $str
#find src/* -name "*.go" -type f ! -path "src/github.com/*" ! -path "src/golang.org/*" | xargs sed -i '1i\ '' '
find src/* -name "*.go" -type f ! -path "src/github.com/*" ! -path "src/golang.org/*" |xargs sed -i '1r t'
find src/* -name "*.go" -type f ! -path "src/github.com/*" ! -path "src/golang.org/*" |xargs sed -i '1d'
#sed -i '1i\ '' ' README.md
#sed -i '1r t' README.md
#sed -i '1d' README.md
