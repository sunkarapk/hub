var path = require('path')
  , fs = require('fs')
  , app = require('flatiron').app;

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

  repo = repo.split('/');

  if (repo.length == 2) {
    if (fs.existsSync(path.join(process.cwd(), repo[1]))) {
      return cb(new Error('destination path \'' + repo[1] + '\' already exists and is not an empty directory'));
    }
    app.git.clone(cmd + repo.join('/'), cb);
  } else {
    return cb(1);
  }
};
