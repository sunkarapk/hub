var app = require('flatiron').app;

module.exports = function (repo, cb) {
  app.u.private(function (err, me) {
    if (err) return cb(err);

    var fork = function (repo, callback) {
      me.fork(repo, function (err, data) {
        if (err) return callback(new Error('Unable to fork the repository'));
        callback(err, data);
      });
    };

    if (typeof repo == 'string') {
      fork(repo, function (err, data) {
        if (err) return cb(err);
        app.git.clone('--progress git@github.com:' + me.login + '/' + app.u.repo(repo), function (err) {
          if (err) return cb(err);
          app.git.remote('add upstream ' + (app.config.get('protocol') || 'git') + '://github.com/' + repo, function (err) {
            if (err) return cb(err);
            app.git.fetch('origin', cb);
          });
        });
      });
    } else {
      app.u.route(function (err, route) {
        if (err) return cb(new Error('Not a valid repository'));
        fork(route, function (err, data) {
          if (err) return cb(err);
          app.git.remote('rename origin upstream', function (err) {
            if (err) cb(err);
            app.git.remote('add origin git@github.com:' + me.login + '/' + app.u.repo(route), function (err) {
              if (err) return cb(err);
              app.git.fetch('origin', cb);
            });
          });
        });
      });
    }
  });
};
