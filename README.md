# gweb
A simple and opinionated web framework mostly based on the go standard library, server side includes and javascript import maps

## usage
Install [air](https://github.com/cosmtrek/air) then run `air server.go .air.toml`. This will build the executable into `tmp/` then run it. Details of what is run is in the `build.sh` file.

## features
- watch and restart the server with `air`
- simple [router](https://github.com/julienschmidt/httprouter) for convenince and developer experience 
- built in go templates for rendering nested html templates
- serve static files from `public` folder 
- minify css `brew install tdewolff/tap/minify`
- parse and use `config.toml`
- use javascript import maps installed with `npm` and configured in `config.toml`

## upcoming features
- automatic page reload
- better css tooling
- session handling, start with cookies
- server side inclundes

## experiment watch with `run-dev.sh`
```
~/go/bin/arelo -p '**/*.go' -p '**/*.html' -p '**/*.css' -p '**/*.js' -i 'public/.*' -i 'tmp/.*' -i 'vendor/.*' -i 'testdata/.*' -i 'assets/node_modules/.*' -i '**/*_test.go' -- go run .
```
