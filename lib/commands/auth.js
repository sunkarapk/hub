var fs = require('fs')
  , app = require('flatiron').app;

module.exports.usage = [
  'Manage github access tokens',
  '',
  'public  - Use public github data for given user',
  'private - Give access to private data using token',
  'destroy - Destory authorization and delete username'
];

module.exports.public = function (cb) {
  if (!app.config.get('auth')) {
    app.prompt.get(['username'], function (err, result) {
      app.log.info('Saved username as ' + result.username.magenta);
      app.commands.config.set('username', result.username, cb);
    });
  } else {
    cb(new Error('Please destroy authorization data before changing username'));
  }
};

module.exports.public.usage = [
  'Save username so that it\'s public data can be used',
  'This do not allow you to perform any private actions',
  '',
  'This will select the public usage mode'
];

module.exports.private = function (cb) {
  app.prompt.get(['username', 'password'], function (err, result) {
    app.gh.auth.config(result).login({
      note: 'HUB - terminal github',
      note_url: 'http://pksunkara.github.com/hub',
      scopes: ['user', 'repo', 'gist']
    }, function (err, id, token) {
      if (err) {
        cb(err);
      } else {
        app.log.info('Authenticated as ' + result.username.magenta);
        app.commands.config.set('auth', { id: id, token: token}, function () {
          fs.chmodSync(app.config.stores.file.file, 0600);
          app.commands.config.set('username', result.username, cb);
        });
      }
    });
  });
};

module.exports.private.usage = [
  'Authenticate to github and save access token',
  'This is required before any other private actions',
  '',
  'This will select the private usage mode'
];

module.exports.destroy = function (cb) {
  var username = app.config.get('username')
    , auth = app.config.get('auth');

  if (auth && username) {
    app.log.warn('You need to give github password to revoke authorization');
    app.prompt.get(['password', 'destroy'], function (err, result) {
      delete result.destroy;
      result.username = username;

      app.gh.auth.config(result).revoke(auth.id, function (err) {
        app.commands.config.delete('username', function () {
          app.commands.config.delete('auth', function () {
            app.log.info('Destroyed authorization and username data');
            cb();
          });
        });
      });
    });
  } else {
    cb(new Error('Either auth or username missing in config file'));
  }
};

module.exports.destroy.usage = [
  'Destroy the authorization token and username data',
  '',
  'This will select the general usage mode'
];
