# go-web
A simple and opinionated web framework mostly based on the go standard library, server side includes and javascript import maps

## usage
Install [air](https://github.com/cosmtrek/air) then run `air server.go .air.toml`. This will build the executable into `tmp/` then run it. Details of what is run is in the `build.sh` file.

## features
- watch and restart the server with `air` configured in `.air.toml`
- simple [router](https://github.com/julienschmidt/httprouter) for convenince and developer experience 
- built in go templates for rendering nested html templates
- serve static files from `public` folder 
- better css tooling enabled by a full project including design tokens in `assets/css-toolchain`
- build and minify CSS with [parcel-css](https://github.com/parcel-bundler/parcel-css)
- parse and use `config.toml`
- use javascript import maps installed with `npm` and configured in `config.toml`
- automatic page reload with a simple socket signalling enabled by [gorilla/websocket](https://github.com/gorilla/websocket)

## upcoming features

- session handling, start with cookies
- server side inclundes
