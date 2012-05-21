var app = require('flatiron').app;

module.exports = function (action, repo, path, cb) {
  if (arguments.length == 4 && action == 'add' && repo.indexOf(':') == -1) {
    var cmd = 'add ';

    if (app.argv.b) {
      cmd += '-b ' + app.argv.b + ' ';
    }

    if (app.argv.private) {
      cmd += 'git@github.com:'
    } else {
      cmd += (app.config.get('protocol') || 'git') + '://github.com/';
    }

    app.git.submodule(cmd + repo + ' ' + path, cb);
  } else {
    arguments[arguments.length - 1](1);
  }
}
