onAfterBootstrap((e) => {
  const utils = require(`${__hooks}/utils.js`);

  utils.hello('world');
});
