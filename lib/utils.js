var app = require('flatiron').app;

app.u = {};

var remotes = {};

app.u.private = function (cb) {
  var username = app.config.get('username')
    , auth = app.config.get('auth');

  if (auth && username) {
    var me = app.gh.client(auth.token).me();
    me.login = username;
    cb(null, me);
  } else {
    cb(new Error('You need to be authorized for this action'));
  }
};

app.u.user = function (cb) {
  if (typeof cb == 'function') {
    app.u.route(function (err, r) {
      if (err) cb(err);
      else cb(null, r.split('/')[0]);
    });
  } else if (cb) {
    return process.cwd().split('/').slice(-2,-1)[0];
  } else {
    return cb.split('/')[0];
  }
};

app.u.repo = function (cb) {
  if (typeof cb == 'function') {
    app.u.route(function (err, r) {
      if (err) cb(err);
      else cb(null, r.split('/')[1]);
    });
  } else if (cb)  {
    return process.cwd().split('/').slice(-1)[0];
  } else {
    return cb.split('/')[1];
  }
};

app.u.remotes = function (remote, cb) {
  if (typeof remote == 'function') {
    cb = remote;
    remote = null;
  }

  function ret(r) {
    cb(null, (remote ? remotes[remote] : remotes));
  }

  if (Object.keys(remotes).length > 0) {
    ret();
  } else {
    app.git.remote('-v', function (err, stdout, stderr) {
      if (err || stderr) return cb(err || stderr);
      stdout.split('\n').forEach(function (line) {
        line = line.split('\t');
        if (line[1] && !remotes[line[0]]) {
          remotes[line[0]] = app.u.url(line[1].split(' ')[0]);
        }
      });
      ret();
    });
  }
};

app.u.route = function (cb) {
  if (typeof cb == 'function') {
    app.u.remotes(function (err, r) {
      if (err) return cb(err);
      if (r.origin) {
        cb(null, r.origin);
      } else {
        cb(null, process.cwd().split('/').slice(-2).join('/'));
      }
    });
  } else {
    return process.cwd().split('/').slice(-2).join('/');
  }
};

app.u.url = function (r) {
  if (r.indexOf('github.com/')) {
    return r.split('github.com/')[1];
  } else if (r.indexOf(':')) {
    return r.split(':')[1];
  } else {
    return r;
  }
};

app.u.issue = function (s, cb) {
  if (s[0] == '#') s = s.slice(1);
  if (parseInt(s)) {
    app.u.route(function (err, route) {
      cb(err, { n: parseInt(s), r: route });
    });
  } else if (s.indexOf('#')) {
    s = s.split('#');
    if (s[0].indexOf('/') == -1) {
      s[0] = app.config.get('username') + '/' + s[0];
    }
    cb(null, { n: parseInt(s[1]), r: s[0] });
  } else {
    cb(new Error('Should be a valid issue'));
  }
};
