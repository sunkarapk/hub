var app = require('flatiron').app;

module.exports = function (user, cb) {
  var cmd, users
    , username = app.config.get('username')
    , auth = app.config.get('auth');

  if (user.indexOf(',') != -1) {
    users = user.split(',');
  } else {
    users = [user];
  }

  require('async').forEachSeries(users, function (user, callback) {
    require('./remote')('add', user, callback)
  }, function (err) {
    if (err) cb(err);
    if (users.length > 1) {
      app.git.fetch('--multiple ' + users.join(' '), cb);
    } else {
      app.git.fetch(users[0], cb);
    }
  });
}

module.exports.usage = [
  'Add github url as remote'
];
