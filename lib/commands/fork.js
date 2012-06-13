var app = require('flatiron').app;

module.exports = function (repo, cb) {
  var username = app.config.get('username')
    , auth = app.config.get('auth');

  if (auth && username) {
    var me = app.gh.client(auth.token).me();

    var fork = function (repo, cb) {
      me.fork(repo, function (err, data) {
        if (err) return cb(new Error('Unable to fork the repository'));
        cb(err, data);
      });
    };

    if (typeof repo == 'string') {
      fork(repo, function (err, data) {
        if (err) return cb(err);
        app.git.clone('--progress git@github.com:' + username + '/' + app.u.repo(repo), function (err) {
          if (err) return cb(err);
          app.git.remote('add upstream ' + (app.config.get('protocol') || 'git') + '://github.com/' + repo, cb);
        });
      });
    } else {
      app.u.route(function (err, route) {
        if (err) return cb(new Error('Not a valid repository'));
        fork(route, function (err, data) {
          if (err) return cb(err);
          app.git.remote('rename origin upstream', function (err) {
            if (err) cb(err);
            app.git.remote('add origin git@github.com:' + username + '/' + app.u.repo(route), cb);
          });
        });
      });
    }
  } else {
    cb(new Error('You need to be authorized for this action'));
  }
}
