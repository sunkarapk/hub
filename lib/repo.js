var app = require('flatiron').app;

app.repo = function (cb) {

  if (typeof cb == 'function') {
    cb(null, require('path').basename(process.cwd()));
  } else {
    return cb.split('/')[1];
  }
}
