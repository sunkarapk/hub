var app = require('flatiron').app;

module.exports = function (sha, cb) {
  var cmd
    , username = app.config.get('username')
    , auth = app.config.get('auth');

  sha = sha.split('@');
  if (sha.length == 2) {
    require('./fetch')(sha[0], function (err) {
      app.git['cherry-pick'](sha[1], cb);
    });
  } else {
    return cb(new Error('Invalid sha given'));
  }
}

module.exports.usage = [
  'Add github url as remote'
];
