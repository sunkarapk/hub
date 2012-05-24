var app = require('flatiron').app;

module.exports = function (repo, cb) {
  var username = app.config.get('username')
    , auth = app.config.get('auth');

  if (typeof repo == 'function') {
    cb = repo;
    repo = null;
  }

  var ask = function (repo, me) {
    app.prompt.get(['description', 'homepage'], function (err, result) {
      if (err) return cb(new Error('Unable to get information about the new repository'));
      me.repos({
        name: repo,
        description: result.description,
        homepage: result.homepage,
        private: app.argv.private
      }, function (err, data) {
        if (err) return cb(new Error('Unable to create repository on github'));
        app.log.info('Created repository at ' + data.html_url.magenta);
        require('./remote')('add', 'origin', cb);
      });
    });
  };

  if (auth && username) {
    var me = app.gh.client(auth.token).me();

    if (repo) {
      if (repo.indexOf('/')) {
        var repos = repo.split('/');
        me = me.org(repos[1]);
        repo = repos[0];
      }
      ask(repo, me);
    } else {
      app.u.repo(function (err, repo) {
        if (err) return cb(new Error('Not a valid github repository'));
        ask(repo, me);
      });
    }
  } else {
    cb(new Error('You need to be authorized to create a repository'));
  }
}
