# go-web
A simple and opinionated web framework mostly based on the go standard library, server side includes and javascript import maps

## usage
```docs being updated on a regular basis, also becoming more detailed```

Make sure have `go` installed, preferabbly as recent as your platform of choice allows.

Install [air](https://github.com/cosmtrek/air) then run `air .air.toml`. This will build the executable into `tmp/` then run it. Details of what air runs are in the `build.sh` file.

## developer tooling
- preconfigured `.air.toml` to be able to use it out of the box. Documentation on options available in the [air](https://github.com/cosmtrek/air) repository. Just run `air server.go .air.toml` to start the server app.
- automatic page reload with a simple socket signalling enabled by [gorilla/websocket](https://github.com/gorilla/websocket)

## simple router
Even though Go STD provides this functionality out of the box, the [router](https://github.com/julienschmidt/httprouter) ads a few features for convenince and developer experience
Examples for how to handle a few kinds of routes
- html output
- JSON API output
- JWT examples, both issuing a new token at login and reading its claims from the request

## built in go templates for rendering nested html templates
Go provides the `html/template package to handle html templating. The `templates` folder contains all layouts and partials, where a layout defines a page layout while a partial is a reusable html snippet. Also, templates are embededd into the built binary with `embed.FS` to simplify delivery to production.

## easily serve static files (css/js/images/etc)
App is easily serving static files from `public` folder like this: `/public/some.file` available on `/public/some.file` url. Also, static files are embedded into binary.

## better css
Basic design system included based on [CubeCSS](https://cube.fyi/). Key principles:
- use progressive enhancement
- structure the CSS in these four groups: composition styles, utilities, blocks and exceptions
The most important part is that we should try to guide the browser to do what it does best (rendering) in a context that it finds itself in.

Design tokens are in a few files: defined colors, spacing values and text sizes, along with a global reset and global styling to bring all browsers on the same page. These are used to build the actual, fluid, styling of the pages.

## package and minify CSS
This done with `lightningcss` which can be installed like this `npm i -g lightningcss-cli` or any other way that makes the cli tool available to `npx`

## example configuration use 
Use c`config/config.go` to showcase some examples of how configuration will be done. Later environment variables will also be transitioned here.

## use javascript import maps 
JavaScript dependencies are installed with `npm` into the `assets/` folder, then minified/copied if needed into the `public` folder when the server starts. They are configured in the `config.toml` so that they become available as JavaScript import maps in the HTML head.
Where needed, [ES Module Shims](https://ga.jspm.io/npm:es-module-shims@1.5.1/dist/es-module-shims.js) is helping fill in the functionalöity in unsupportive browsers.

## middleware
Example of middleware implementations, starting with a request logger

## upcoming features
- session handling, start with cookies
- server side inclundes
- database example, Posgres with [pgx](https://github.com/jackc/pgx)
