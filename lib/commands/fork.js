var app = require('flatiron').app;

module.exports = function (cb) {
  var username = app.config.get('username')
    , auth = app.config.get('auth');

  if (auth && username) {
    var me = app.gh.client(auth.token).me();
    app.u.route(function (err, route) {
      if (err) return cb(new Error('Not a valid github repository'));
      me.fork(route, function (err, data) {
        if (err) return cb(err);
        app.git.remote('add -f ' + username + ' git@github.com:' + username + '/' + app.u.repo(route), cb);
      });
    });
  } else {
    cb(new Error('You need to be authorized for this action'));
  }
}
