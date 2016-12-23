#! /bin/bash

pkg=github.com/udhos/yname

step=0

msg() {
    step=$((step+1))
    echo >&2 $step. $*
}

get() {
    i=$1
    msg fetching $i
    go get $i
    msg fetching $i - done
}

get gopkg.in/yaml.v2
#get honnef.co/go/simple/cmd/gosimple

src=`find . -type f | egrep '\.go$'`

msg fmt
gofmt -s -w $src
msg fix
go tool fix $src
msg vet
go tool vet .

pushd $GOPATH/src/$pkg >/dev/null
examples=`echo examples/*`
popd >/dev/null

#echo examples: $examples

msg install
go install $pkg
for i in $examples; do
    msg install $pkg/$i
    go install $pkg/$i
done

# go get github.com/golang/lint/golint
l=$GOPATH/bin/golint
lint() {
    msg lint
    # golint cant handle source files from multiple packages
    pushd $GOPATH/src/$pkg >/dev/null
    $l *.go
    for i in $examples; do
	msg lint $i
	$l $i/*.go
    done
    popd >/dev/null
}
[ -x "$l" ] && lint

# go get honnef.co/go/simple/cmd/gosimple
s=$GOPATH/bin/gosimple
simple() {
    msg simple - this is slow, please standby
    # gosimple cant handle source files from multiple packages
    pushd $GOPATH/src/$pkg >/dev/null
    $s *.go
    for i in $examples; do
	msg simple $i
	$s $i/*.go
    done
    popd >/dev/null
}
[ -x "$s" ] && simple

msg test yname
go test github.com/udhos/yname
