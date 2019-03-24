const config = require('./config');
const { RenderTemplate, EditField, FillControllerSelects } = require('./util');
const {
  ListConsoles,
  CreateConsole,
  ChangeConsole,
  DeleteConsole,
  AddConsoleController,
  DeleteConsoleController
} = require('./api');
const $ = require('jquery');
const FormData = require('form-data');

$(() => {

var RefreshList;
var RefreshCallbacks;

RefreshCallbacks = function() {
  $(".console .delete").click(function() {
    const console = $(this).parents(".console");
    const name = $(console).find(".name").text();
    const consoleId = console.data("consoleId");
    if(confirm("Are you sure you want to delete \"" + name + "\"?"))
      DeleteConsole(consoleId);
  });
  $(".console .name").click(function() {
    EditField(this, function(name) {
      const console = $(name).parents(".console");
      const consoleId = console.data("consoleId");
      ChangeConsole(
        consoleId,
        {"Name": $(name).text()}
      );
      RefreshCallbacks();
    });
  });
  $(".console .add-controller-type").click(function() {
    const consoleObj = $(this).parents(".console");
    const controllerId = consoleObj.find(".controller-select").val();
    if(controllerId == "SelectName")
      return;
    const consoleId = $(consoleObj).data("consoleId");
    AddConsoleController({
      "ConsoleID": consoleId,
      "ControllerID": controllerId
    });
    RefreshList();
  });
  $(".console .delete-console-controller").click(function() {
    const console = $(this).parents(".console-controller");
    const consoleControllerId = $(console).data("consoleControllerId");
    DeleteConsoleController(consoleControllerId);
    RefreshList();
  });
}

RefreshList = function() {
  ListConsoles().then((consoles) => {
    if(typeof(consoles) === "object" && consoles.length > 0)
      RenderTemplate(
        './mustache/consolelist.mst',
        {
          'Consoles': consoles,
          'baseURL': config.baseURL,
        },
        (html) => {
          $('#console-list').html(html);
          RefreshCallbacks();
          FillControllerSelects();
        }
      );
    else {
      $('#console-list').html("<h1>No consoles</h1>");
      RefreshCallbacks();
      FillControllerSelects();
    }
  });
}

$("#new-console-type").submit(function(e) {
  const formData = new FormData(this);
  CreateConsole(formData);
  RefreshList();
  e.preventDefault();
});

RefreshList();

});
