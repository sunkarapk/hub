var app = require('flatiron').app;

module.exports = function (action, user) {
  var cmd
    , username = app.config.get('username');

  if (action == 'add' && arguments.length == 3) {
    var cb = arguments[2];
    cmd = 'add ' + user + ' ';

    if (app.argv.private || user == 'origin') {
      cmd += 'git@github:';
    } else {
      cmd += (app.config.get('protocol') || 'git') + '://github.com/';
    }

    if (user == 'origin') {
      if (username) {
        user = username;
      } else {
        return cb(new Error('No username found to add origin remote'));
      }
    }

    if (user.indexOf('/') == -1) {
      app.u.repo(function (err, repo) {
        if (err) return cb(new Error('Not a valid repository'));
        app.git.remote(cmd + user + '/' + repo, cb);
      });
    } else {
      app.git.remote(cmd + user, cb);
    }
  } else {
    return arguments[arguments.length - 1](1);
  }
}
