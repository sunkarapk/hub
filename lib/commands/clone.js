var app = require('flatiron').app;

module.exports = function (repo, cb) {
  var cmd
    , username = app.config.get('username')
    , auth = app.config.get('auth');

  if (!repo) {
    app.showHelp('clone');
    return cb();
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

  app.git.clone(cmd + repo, cb);
};

module.exports.usage = [
  'Clone github repo'
];
