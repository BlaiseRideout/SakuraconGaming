
//const config = Load();
//

function LoadConfig() {
  let config = {};
  const AddToConfig = function(partialConfig) {
    for(const [key, value] of Object.entries(partialConfig)){
      if(key.charAt(0) !== '_')
        config[key] = value;
    }
  };

  const defaults = require('./defaultConfig');
  AddToConfig(defaults);

  const appDataPath = process.platform === 'win32' ?
      process.env.APPDATA :
      process.env.HOME +
      (process.platform == 'darwin' ?
        '/Library/Preferences' :
        '/.local/share');

  const userConfigPath = appDataPath + "/SakuraconGaming/config";

  try {
    const userConfig = require(userConfigPath);
    AddToConfig(userConfig);
  }
  catch(e) {
    console.log("No user config found at", userConfigPath);
  }

  return config;
}

module.exports = LoadConfig();
