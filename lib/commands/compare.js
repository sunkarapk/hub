var app = require('flatiron').app;

module.exports = function (arg, cb) {
  var username = app.config.get('username');

  app.u.route(function (err, route) {
    if (err) return cb(new Error('Not a valid github repository'));

    if (arg && typeof cb == 'function') {
      app.git.browse('https://github.com/' + route + '/compare/' + arg, cb);
    } else {
      return cb(new Error('Not a valid github comparison'));
    }
  });
}
