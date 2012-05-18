var app = require('flatiron').app;

module.exports = function (action, user, cb) {
  var cmd
    , username = app.config.get('username')
    , auth = app.config.get('auth');

  if (action == 'add') {
    cmd = 'add ' + user + ' ';

    if (app.argv.private) {
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
    return cb(new Error('Don\'t know `remote ' + action + '` command'));
  }
}

module.exports.usage = [
  'Add github url as remote'
];
