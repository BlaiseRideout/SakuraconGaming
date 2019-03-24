const config = require('./config');

const fetch = require('node-fetch');
const FormData = require('form-data');

async function RequestObject(endpoint, method, data) {
  const fullEndpoint = config.baseURL + "/" + endpoint;
  if(method === undefined)
    method = 'GET';

  var args = {
    'method': method,
  };

  if(method !== 'GET' && data !== undefined)
    args['body'] = data;
  const response = await fetch(fullEndpoint, args);
  const text = await response.text();
  try {
    const json = JSON.parse(text);
    return json;
  }
  catch(err) {
    var object = {};
    if(data instanceof FormData)
      data.forEach((value, key) => {object[key] = value});
    else
      object = data;

    console.log("ERROR in", method, "to", endpoint, ":", "\"" + text + "\"", "with data:", object);
    return null;
  }
}

//////////////////////////
// CONTROLLER FUNCTIONS //
//////////////////////////

// Returns a list of controller types and the current count.
exports.ListControllerTypes = function() {
  if(config.dummyData)
    return [
      {
        "ID": 0,
        "Name":"Gamecube",
        "Image":"/images/controllers/gamecube.png",
        "Count":5
      }
    ];
  return RequestObject('admin/controllers');
}

// Creates a new controller
exports.CreateController = function(data) {
  if(config.dummyData)
    return [
      {
        "ID": 0,
        "Name":"Gamecube",
        "Image":"/images/controllers/gamecube.png",
        "Count":5
      }
    ];
  return RequestObject('admin/controllers', 'POST', data);
}

// Changes a given controller
exports.ChangeController = function(id, data) {
  if(config.dummyData)
    return [
      {
        "ID": id,
        "Name":"Gamecube",
        "Image":"/images/controllers/gamecube.png",
        "Count":5
      }
    ];
  return RequestObject('admin/controllers/' + id, 'POST', data);
}

// Returns a list of controller types and the current count.
exports.DeleteController = function(id) {
  if(config.dummyData)
    return [];
  return RequestObject('admin/controllers/' + id, 'DELETE');
}

///////////////////////
// CONSOLE FUNCTIONS //
///////////////////////

// Returns a list of console types and the current count.
exports.ListConsoles = function() {
  if(config.dummyData)
    return [
      {
        "ID": 0,
        "Controllers":[
          {
            "ID": 0,
            "Name":"Gamecube",
            "Image":"/images/controllers/gamecube.png",
            "Count":5
          }
        ],
        "Name":"Gamecube",
        "Image":"/images/controllers/gamecube.png"
      }
    ];
  return RequestObject('admin/consoles');
}

exports.CreateConsole = function(data) {
  if(config.dummyData)
    return [
      {
        "ID": 0,
        "Name":"Gamecube",
        "Image":"/images/controllers/gamecube.png"
      }
    ];
  return RequestObject('admin/consoles', 'POST', data);
}

exports.ChangeConsole = function(id, data) {
  if(config.dummyData)
    return [
      {
        "ID": id,
        "Name":"Gamecube",
        "Image":"/images/controllers/gamecube.png"
      }
    ];
  return RequestObject('admin/consoles/' + id, 'POST', data);
}

exports.DeleteConsole = function(id) {
  if(config.dummyData)
    return [ ];
  return RequestObject('admin/consoles/' + id, 'DELETE');
}

/////////////////////////////////
// CONSOLECONTROLLER FUNCTIONS //
/////////////////////////////////

exports.AddConsoleController = function(data) {
  if(config.dummyData)
    return [
      {
        "ID": 0,
        "ConsoleID":0,
        "ControllerID":0
      }
    ];
  return RequestObject('admin/console_controllers', 'POST', data);
}


///////////////////////
// STATION FUNCTIONS //
///////////////////////

// Returns a list of stations
exports.ListStations = function() {
  if(config.dummyData)
    return [
      {
        "ID": 0,
        "ConsoleID":0,
        "X":0,
        "Y":0
      }
    ];
  return RequestObject('admin/stations');
}

exports.CreateStation = function(data) {
  if(config.dummyData)
    return [
      {
        "ID": 0,
        "ConsoleID":0,
        "X":0,
        "Y":0
      }
    ];
  return RequestObject('admin/stations', 'POST', data);
}

exports.ChangeStation = function(id, data) {
  if(config.dummyData)
    return [
      {
        "ID": 0,
        "ConsoleID":0,
        "X":0,
        "Y":0
      }
    ];
  return RequestObject('admin/stations/' + id, 'POST', data);
}

exports.DeleteStation = function(id) {
  if(config.dummyData)
    return [ ];
  return RequestObject('admin/stations/' + id, 'DELETE');
}

