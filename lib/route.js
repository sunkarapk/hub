var app = require('flatiron').app;

app.route = function (cb) {
  var cwd = process.cwd().split('/');

  cb(null, cwd.slice(-2).join('/'));
}
