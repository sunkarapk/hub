var app = require('flatiron').app;

app.u = {};

app.u.user = function (cb) {
  if (typeof cb == 'function') {
    app.u.route(function (err, r) {
      if (err) cb(err);
      else cb(null, r.split('/')[0]);
    });
  } else {
    return cb.split('/')[0];
  }
}

app.u.repo = function (cb) {
  if (typeof cb == 'function') {
    app.u.route(function (err, r) {
      if (err) cb(err);
      else cb(null, r.split('/')[1]);
    });
  } else {
    return cb.split('/')[1];
  }
}

app.u.remotes = function (remote, cb) {
  if (typeof remote == 'function') {
    cb = remote;
    remote = null;
  }
  var remotes = {};

  app.git.remote('-v', function (err, stdout, stderr) {
    if (err || stderr) return cb(err || stderr);
    stdout.split('\n').forEach(function (line) {
      line = line.split('\t');
      if (line[1] && !remotes[line[0]]) {
        remotes[line[0]] = app.u.url(line[1].split(' ')[0]);
      }
    });

    if (remote) {
      cb(null, remotes[remote]);
    } else {
      cb(null, remotes);
    }
  });
}

app.u.route = function (cb) {
  app.u.remotes(function (err, r) {
    if (err) return cb(err);
    if (r.origin) {
      cb(null, r.origin);
    } else {
      cb(null, process.cwd().split('/').slice(-2).join('/'));
    }
  });
}

app.u.url = function (r) {
  if (r.indexOf(':')) {
    return r.split(':')[1];
  } else if (r.indexOf('github.com/')) {
    return r.split('github.com/')[1];
  } else {
    return r;
  }
}
