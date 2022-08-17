#!/bin/bash
~/go/bin/arelo -p '**/*.go' -p '**/*.html' -p '**/*.css' -p '**/*.js' -i 'public/.*' -i 'tmp/.*' -i 'vendor/.*' -i 'testdata/.*' -i 'assets/node_modules/.*' -i '**/*_test.go' -- go run .
