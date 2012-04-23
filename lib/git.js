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
        var git = spawn('git', cmd.split(' '));

        git.stdout.pipe(process.stdout, {end: false});
        git.stderr.pipe(process.stderr, {end: false});

        git.on('exit', function (code) {
          cb(code);
        });
      }

      ['clone'].forEach(function (e) {
        app.git[e] = function (cmd, cb) {
          proxy(e + ' ' + cmd, cb);
        };
      });

      callback();
    }
  });
};
