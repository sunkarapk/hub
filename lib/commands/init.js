var app = require('flatiron').app;

module.exports = function (cb) {
  var username = app.config.get('username');

  if (typeof cb == 'function' && app.argv.g) {
    if (username) {
      app.u.repo({}, function (err, repo) {
        if (err) return cb(new Error('Not a valid repository'));
        app.git.init('.', function (err) {
          if (err) return cb(err);
          require('./remote')('add', 'origin', cb);
        });
      });
    } else {
      cb(new Error('You need to provide username for this action'));
    }
  } else {
    arguments[arguments.length - 1](1);
  }
}
