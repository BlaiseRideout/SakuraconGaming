const config = require('./config');
const { RenderTemplate, EditField } = require('./util');
const { ListControllerTypes, CreateController, ChangeController, DeleteController } = require('./api');
const $ = require('jquery');

$(() => {

var RefreshList;
var RefreshCallbacks;

RefreshCallbacks = function() {
  function UpdateCount(count, newCount){
    const controller = $(count).parents(".controller");
    const controllerId = controller.data("controllerId");
    $(count).text(newCount);
    if(newCount === 0 && confirm("Do you want to remove this controller type?")){
      DeleteController(controllerId);
      RefreshList();
    }
    else
      ChangeController(
        controllerId,
        {"Count": newCount}
      );
  }
  function CountButton(className, valueAdd){
    $(className).click(function() {
      const controller = $(this).parents(".controller");
      const newCount = Math.max(0, count.text() - -valueAdd);
      const count = controller.find(".count");
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
  RenderTemplate(
    './mustache/controllerlist.mst',
    {'Controllers':ListControllerTypes()},
    (html) => {
      $('#controller-list').html(html);
      RefreshCallbacks();
    }
  );
}

$("#new-controller-type").submit(function(e) {
  const formData = new FormData(this);
  CreateController(formData);
  RefreshList();
  e.preventDefault();
});

RefreshList();

});
