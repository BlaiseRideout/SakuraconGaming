const config = require('./config');
const { RenderTemplate, FillConsoleSelects, Draggable } = require('./util');
const {
  ListStations,
  AddStation,
  ChangeStation,
  DeleteStation
} = require('./api');
const $ = require('jquery');

$(() => {

var RefreshList;
var RefreshCallbacks;

RefreshCallbacks = function() {
  $(".station .delete").click(function() {
    const station = $(this).parents(".station");
    const id = $(station).data("id");
    if(confirm("Are you sure you want to delete station \"" + id + "\"?"))
      DeleteStation(id);
  });
  Draggable(".station", (station) => {
    const id = station.data("id");
    const x = station.offset()["left"];
    const y = station.offset()["top"];
    station.data("x", x);
    station.data("y", y);
    ChangeStation(id, {
      "X": x,
      "Y": y
    });
    RefreshCallbacks();
  },
 ".header, img");
  $("#create-station").click(function() {
    CreateStation();
    RefreshList();
  });
  $(".station .console-select").change(function() {
    const station = $(this).parents(".station");
    const id = $(station).data("id");
    const newVal = $(this).val();
    if(newVal !== "SelectName")
      ChangeStation(id, {
        "CosoleID": newVal
      });
  })
}

RefreshList = function() {
  RenderTemplate(
    './mustache/stations.mst',
    {'Stations':ListStations()},
    (html) => {
      $('#stations').html(html);
      FillConsoleSelects();
      $(".station").each(function(i, station) {
        const x = $(station).data("x");
        const y = $(station).data("y");
        if(x !== undefined)
          $(station).css({"left":x});
        if(y !== undefined)
          $(station).css({"top":y});
        const consoleSelect = $(station).find("console-select");
        const console = consoleSelect.data("console");
        if(console !== undefined)
          consoleSelect.val(console);
      });
      RefreshCallbacks();
    }
  );
}

$("#new-console-type").submit(function(e) {
  const formData = new FormData(this);
  CreateConsole(formData);
  RefreshList();
  e.preventDefault();
});

RefreshList();

});
