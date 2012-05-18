var app = require('flatiron').app;

app.repo = function (cb) {
    cb(require('path').basename(process.cwd()));
}
