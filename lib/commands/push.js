var app = require('flatiron').app;

module.exports = function (repo, ref, cb) {
  if (arguments.length == 3 && repo.indexOf(',') != -1) {
    require('async').forEachSeries(repo.split(','), function (r, callback) {
      app.git.push(r + ' ' + ref, callback)
    }, function (err) {
      cb(new Error('Unable to push to remote repositories'));
    });
  } else {
    arguments[arguments.length - 1](1);
  }
}
