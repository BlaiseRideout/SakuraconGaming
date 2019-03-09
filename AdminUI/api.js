const config = require('./config');

const fetch = require('node-fetch');

async function RequestObject(endpoint, method, data) {
  const fullEndpoint = config.baseURL + endpoint;
  if(data === undefined)
    data = {};
  if(method === undefined)
    method = 'GET';

  const response = await fetch(fullEndpoint, {
    'method': method,
    'body': data
  });
  const json = await response.json();
  return JSON.parse(json);
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
  return RequestObject('admin/controllers', 'DELETE', {"ID":id});
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
  return RequestObject('admin/consoles', 'DELETE', {"ID": data});
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