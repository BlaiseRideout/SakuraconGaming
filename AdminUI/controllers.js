const config = require('./config');
const { RenderTemplate, EditField } = require('./util');
const { ListControllerTypes, CreateController, ChangeController, DeleteController } = require('./api');
const $ = require('jquery');
const FormData = require('form-data');

$(() => {

var RefreshList;
var RefreshCallbacks;

RefreshCallbacks = function() {
  function UpdateCount(count, newCount){
    const controller = $(count).parents(".controller");
    const controllerId = controller.data("controllerId");
    if(newCount > 0) {
      ChangeController(
        controllerId,
        {"Count": newCount}
      );
      $(count).text(newCount);
    }
    else if(confirm("Do you want to remove this controller type?")) {
      DeleteController(controllerId);
      RefreshList();
    }
  }
  function CountButton(className, valueAdd){
    $(className).click(function() {
      const controller = $(this).parents(".controller");
      const count = controller.find(".count");
      const newCount = Math.max(0, count.text() - -valueAdd);
      UpdateCount(count, newCount);
    });
  }
  CountButton(".inc-count", +1);
  CountButton(".dec-count", -1);
  $(".controller .count").click(function() {
    EditField(this, function(count) {
      UpdateCount($(count).text());
      RefreshCallbacks();
    });
  });
  $(".controller .name").click(function() {
    EditField(this, function(name) {
      const controller = $(name).parents(".controller");
      const controllerId = controller.data("controllerId");
      ChangeController(
        controllerId,
        {"Name": $(name).text()}
      );
      RefreshCallbacks();
    });
  });
}

RefreshList = function() {
  ListControllerTypes().then((controllers) => {
    if(typeof(controllers) === "object" && controllers.length > 0)
      RenderTemplate(
        './mustache/controllerlist.mst',
        {'Controllers':controllers},
        (html) => {
          $('#controller-list').html(html);
          RefreshCallbacks();
        }
      );
    else {
      $('#controller-list').html("<h1>No controllers</h1>");
      RefreshCallbacks();
    }
  });
}

$("#new-controller-type").submit(function(e) {
  const formData = new FormData(this);
  CreateController(formData);
  RefreshList();
  e.preventDefault();
});

RefreshList();

});
