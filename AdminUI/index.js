// This file is required by the index.html file and will
// be executed in the renderer process for that window.
// All of the Node.js APIs are available in this process.

const config = require('./config');
const $ = require('jquery');

if(config.debug){
  $(() => {
    $('#config').text(JSON.stringify(config));
    $("#debug-info").css("display", "block");
  });
}
