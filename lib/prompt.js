module.exports = {
  properties: {
    username: {
      name: 'username',
      validator: /^[a-z\d][a-z\d\-]*$/i,
      warning: 'Username may only contain alphanumeric characters or dashes and cannot begin with a dash'
    },
    password: {
      name: 'password',
      validator: /^(?=.*[a-z])(?=.*\d).{7,}/,
      warning: 'Must contain one lowercase letter, one number, and be at least 7 characters long',
      hidden: true
    }
  }
};
