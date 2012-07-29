var app = require('flatiron').app;

module.exports = function (sha, cb) {
  if (sha && typeof sha == 'string') {
    sha = sha.split('@');
    if (sha.length == 2) {
      require('./fetch')(sha[0], function (err) {
        app.git['cherry-pick'](sha[1], cb);
      });
    } else {
      return cb(1);
    }
  } else {
    return cb(1);
  }
}
