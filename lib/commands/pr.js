var app = require('flatiron').app;

module.exports.usage = [
  'Work with github pull requests',
  '',
  'open  | o - Open a pull request',
  'merge | m - Merge a pull request',
  'close | c - Close a pull request',
  'pull  | p - Pull a pull request'
];

module.exports.c = module.exports.close = function (num, cb) {
  app.u.private(function (err, me) {
    if (err) return cb(err);

    app.u.issue(num, function (err, num) {
      if (err) return cb(err);

      me.pr(num.r, num.n).close(cb);
    });
  });
};

module.exports.m = module.exports.merge = function (num, cb) {
  app.u.private(function (err, me) {
    if (err) return cb(err);

    app.u.issue(num, function (err, num) {
      if (err) return cb(err);

      me.pr(num.r, num.n).merge(function (err, body) {
        if (body.sha && body.merged) {
          app.log.info(body.message.green.bold); cb();
        } else {
          cb(new Error(body.message));
        }
      });
    });
  });
};

module.exports.o = module.exports.open = function (cb) {
  app.u.private(function (err, me) {
    if (err) return cb(err);

    app.u.issue(num, function (err, num) {
      if (err) return cb(err);

      //TODO:master-branch, body editor
    });
  });
};

module.exports.p = module.exports.pull = function (num, cb) {
  app.u.private(function (err, me) {
    if (err) return cb(err);

    app.u.issue(num, function (err, num) {
      if (err) return cb(err);

      me.pr(num.r, num.n).info(function (err, pr) {
        if (err) return cb(err);

        app.git.remote(['pull', pr.head.repo.ssh_url, pr.head.ref].join(' '), cb);
      });
    });
  });
};
