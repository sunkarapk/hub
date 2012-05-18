var app = require('flatiron').app;

module.exports = function (user, cb) {
  var users
    , username = app.config.get('username')
    , auth = app.config.get('auth');

  if (user && typeof user == 'string' && user.indexOf(',') != -1) {
    users = user.split(',');

    app.u.remotes(function (err, remotes) {
      if (err) return cb(new Error('Unable to get remotes of the repository'));
      require('async').forEachSeries(users, function (user, callback) {
        if (!remotes[user]) {
          require('./remote')('add', user, callback);
        }
      }, function (err) {
        if (err) return cb(new Error('Unable to resolve all the remotes'));
        app.git.fetch('--multiple ' + users.join(' '), cb);
      });
    });
  } else {
    return cb(1);
  }
}
