var app = require('flatiron').app;

module.exports = function (repo, cb) {
  var cmd
    , username = app.config.get('username')
    , auth = app.config.get('auth');

  if (!repo || typeof repo != 'string' || repo.indexOf(':') != -1) {
    return cb(1);
  }

  if (app.argv.private) {
    cmd = '--progress git@github.com:';
  } else {
    cmd = '--progress ' + (app.config.get('protocol') || 'git') + '://github.com/';
  }

  if (repo.indexOf('/') == -1) {
    if (auth && auth.token) {
      cmd = '--progress git@github.com:';
    }

    if (username) {
      repo = username + '/' + repo;
    } else {
      return cb(new Error('Cannot find path to ' + repo));
    }
  }

  if (repo.split('/').length == 2) {
    app.git.clone(cmd + repo, cb);
  } else {
    return cb(1);
  }
};
