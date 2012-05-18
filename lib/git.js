var path = require('path')
  , github = require('octonode')
  , exec = require('child_process').exec
  , spawn = require('child_process').spawn;

module.exports.init = function (callback) {
  var app = this;

  exec('which git', function (err, stdout, stderr) {
    if (err) {
      callback(new Error('Please install git on your system'));
    } else {
      app.git = {};
      app.gh = github;

      function proxy(cmd, cb) {
        var key
          , git = spawn('git', cmd.split(' '))
          , out = (typeof cb == 'function');

        if (out) {
          app.log.debug(('git ' + cmd).white.bold);

          git.stdout.pipe(process.stdout, {end: false});
          git.stderr.pipe(process.stderr, {end: false});
        } else {
          git.stdout.pipe(cb, {end: false});
          git.stderr.pipe(cb, {end: false});
        }

        git.on('exit', function (code) {
          if (out) cb(code);
          else cb.emit('exit', code);
        });
      }

      app.cmds.forEach(function (e) {
        app.git[e] = function (cmd, cb) {
          proxy(e + ' ' + cmd, cb);
        };
      });

      require('./repo');

      callback();
    }
  });
};
