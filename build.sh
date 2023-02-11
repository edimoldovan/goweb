#!/bin/bash

echo 'copying javascript files to public folder...';
cp assets/node_modules/flatpickr/dist/flatpickr.min.js public/js;
cp assets/js/web-component.js public/js/;
cp assets/js/is-land.js public/js/;

echo 'copying fonts...'
cp assets/fonts/* public/fonts/

go build -o ./tmp/main .
