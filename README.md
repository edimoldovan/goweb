# gweb
A simple web framework based on go std, server side includes and javascript import maps

## usage
Install [air](https://github.com/cosmtrek/air) then run `air server.go .air.toml`. This will build the executable into `tmp/` then run it. Details of what is run is in the `build.sh` file.

## features

- watch and restart the server with `air`
- simple [router](https://github.com/julienschmidt/httprouter) for convenince and developer experience 
- built in go templates for rendering nested html templates
- serve static files from `public` folder 
