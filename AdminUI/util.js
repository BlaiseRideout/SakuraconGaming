const config = require('./config');

const $ = require('jquery');
const fs = require('fs');
const mustache = require('mustache');
const { ListControllerTypes, ListConsoles } = require('./api');

/*
 * if dialogs get fixed
const { dialog } = require('electron').remote;
*/

exports.RenderTemplate = function(path, requestData, cb) {
  fs.readFile(path, (err, template) => {
    if(err) {
      /*
       * if dialogs get fixed
      dialog.showMessageBox({
        "type":"error",
        "title":"RenderTemplate failed"
        "message":"Couldn't load " + path
      });
      */
      console.error(
        "RenderTemplate failed: Couldn't load " + path
      );
    }
    cb(mustache.render(template.toString("ascii"), requestData));
  });
};

exports.EditField = function(field, cb) {
  function ChangeTag(elem, newTag){
    const oldTag = elem.tagName;
    const newElem = document.createElement(newTag);
    $.each(elem.attributes, function() {
      $(newElem).attr(this.name, this.value);
    });
    $(newElem).data("oldTag", oldTag);
    const val = (oldTag.toLowerCase() === "input")?
      $(elem).val() : $(elem).text();

    if(newTag.toLowerCase() === "input")
      $(newElem).val(val);
    else
      $(newElem).text(val);

    $(elem).replaceWith(newElem);
    return newElem;
  }
  const oldEditField = $("#editfield");
  if(oldEditField.length)
    ChangeTag(oldEditField.get(0), oldEditField.data("oldTag"))
  const newElem = ChangeTag(field, "input");
  const Finish = function(){
    cb($(ChangeTag(this, $(this).data("oldTag"))).text());
  };
  $(newElem).keypress(function(e) {
    if(e.which === 13)
      Finish.bind(this)();
  });
  $(newElem).focusout(Finish);
};

exports.Draggable = function(elemSelector, draggedCb, handleSelector) {
  const handle = (handleSelector === undefined)?
                      $(elemSelector):
                      $(elemSelector).find(handleSelector);
  handle.unbind('mousedown');
  handle.mousedown(function(downEvent) {
    var prevX;
    var prevY;
    const curElem = (handleSelector === undefined)?
                      $(this):
                      $(this).parents(elemSelector);
    curElem.mousemove(function(moveEvent) {
      moveEvent.preventDefault();
      const newX = moveEvent.clientX;
      const newY = moveEvent.clientY;
      const prevOffset = curElem.offset();
      if(prevX !== undefined && prevY !== undefined){
        curElem.offset({
          "left": prevOffset["left"] + newX - prevX,
          "top":  prevOffset["top"]  + newY - prevY
        });
      }
      prevX = newX;
      prevY = newY;
    });
    const stop = function() {
      curElem.unbind('mouseup');
      curElem.unbind('mousemove');
      draggedCb(curElem);
    };
    curElem.mouseleave(stop);
    curElem.mouseup(stop);
  });
};

exports.FillControllerSelects = function() {
  exports.RenderTemplate('./mustache/controllerselect.mst',
    {'Controllers':ListControllerTypes()},
    (html) => {
      $(".controller-select").html(html);
    }
  );
};

exports.FillConsoleSelects = function() {
  exports.RenderTemplate('./mustache/consoleselect.mst',
    {'Consoles':ListConsoles()},
    (html) => {
      $(".console-select").html(html);
    }
  );
};
