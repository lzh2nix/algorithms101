#!/usr/bin/env bash
echo "mode: atomic" > coverage.txt
for d in $(go list ./...); do
    go test -race -covermode=atomic -coverprofile=profile.out $d
    if [ -f profile.out ]; then
        cat profile.out |grep -v "mode: atomic" >> coverage.txt
        rm profile.out
    fi
done
