module.exports = {
  properties: {
    username: {
      required: true,
      pattern: /^[a-z\d][a-z\d\-]*$/i,
      message: 'Username may only contain alphanumeric characters or dashes and cannot begin with a dash'
    },
    password: {
      required: true,
      pattern: /^(?=.*[a-z])(?=.*\d).{7,}/,
      message: 'Must contain one lowercase letter, one number, and be at least 7 characters long',
      hidden: true
    },
    description: {
      required: true,
      message: 'Description must be provided'
    },
    homepage: {
      required: true,
      format: 'url',
      message: 'Must be a url'
    },
  }
};
