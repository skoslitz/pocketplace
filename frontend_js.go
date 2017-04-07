package main

const (
	frontendJs = `// Globals
var SCALE = 30;

var CANVAS = null;
var LOADING = null;
var PICKER = null;

var CONTEXT = null;
var CONNECTION = null;

var MY_COLOR = {r: 63, g: 110, b: 135};
var ON_COOLDOWN = false;

// Initialize
window.onload = function () {
  CANVAS = document.getElementById("canvas");
  LOADING = document.getElementById("loading");
  PICKER = document.getElementById("picker");

  setupCanvas();
  setupWebsocket();
  setupPicker();
};

function setupCanvas () {
  CONTEXT = CANVAS.getContext("2d");


  CANVAS.width = SIZE;
  CANVAS.height = SIZE;

  CANVAS.style.transform = "scale(" + SCALE + ")";

  CANVAS.onclick = function (e) {
    var x = Math.floor(e.pageX / SCALE);
    var y = Math.floor(e.pageY / SCALE);
    if (!ON_COOLDOWN) {
      sendPixel(x, y, MY_COLOR.r, MY_COLOR.g, MY_COLOR.b);
      toggleCooldown();
    }
  };

  CANVAS.oncontextmenu = function (e) {
    e.preventDefault();
  };
}

function toggleCooldown () {
  if (ON_COOLDOWN) { return; }
  if (COOLDOWN <= 0) { return; }

  ON_COOLDOWN = true;

  var el = document.getElementById("cooldown");

  var cooldownMs = COOLDOWN * 1000;
  var cooldownEnd = Date.now() + cooldownMs;

  el.style.display = "block";
  el.textContent = cooldownMs + "ms";

  var intervalId = setInterval(() => {
    el.textContent = (cooldownEnd - Date.now()) + "ms";
  }, 5);

  setTimeout(() => {
    ON_COOLDOWN = false;
    clearInterval(intervalId);
    el.style.display = "none";
  }, cooldownMs);
}

function setupWebsocket () {
  function processCmd (data) {
    var parts = data.split(" ");
    if (parts.length === 5) {
      setPixel(parts[0], parts[1], parts[2], parts[3], parts[4]);
    }
  }

  function fillCanvas (data) {
    var dv = new DataView(data);
    var image = CONTEXT.createImageData(SIZE, SIZE);

    var i, j, k;
    for (i = 0; i < (dv.byteLength / 3); i += 1) {
      j = i*3;
      k = i*4;
      image.data[k] = dv.getUint8(j);
      image.data[k+1] = dv.getUint8(j+1);
      image.data[k+2] = dv.getUint8(j+2);
      image.data[k+3] = 255;
    }

    CONTEXT.putImageData(image, 0, 0);

    LOADING.style.display = "none";
    CANVAS.style.display = "block";
    PICKER.style.display = "block";
  }

  var url = "ws://" + window.location.host + "/ws";

  CONNECTION = new WebSocket(url);
  CONNECTION.binaryType = "arraybuffer";

  CONNECTION.onmessage = function (e) {
    var data = e.data;

    if (typeof data === "string") {
      return processCmd(data);
    }

    if (data instanceof ArrayBuffer) {
      return fillCanvas(data);
    }
  };
}

function setupPicker () {
  var input = document.getElementById("picker-input");
  var preview = document.getElementById("picker-preview");
  var blue = document.getElementById("blueColor");
  var red = document.getElementById("redColor");
  var yellow = document.getElementById("yellowColor");
  var orange = document.getElementById("orangeColor");
  var teal = document.getElementById("tealColor");
  var pink = document.getElementById("pinkColor");
  var olive = document.getElementById("oliveColor");
  var green = document.getElementById("greenColor");
  var violet = document.getElementById("violetColor");
  var purple = document.getElementById("purpleColor");
  var brown = document.getElementById("brownColor");
  var grey = document.getElementById("greyColor");

  blue.onclick = function() {
    var value = "#3f6e87";
    var color = hexToRgb(value);

    if (color) {
      MY_COLOR = color;
      preview.style.backgroundColor = value;
    }
  }

  red.onclick = function() {
    var value = "#8e225c";
    var color = hexToRgb(value);

    if (color) {
      MY_COLOR = color;
      preview.style.backgroundColor = value;
    }
  }

  yellow.onclick = function() {
    var value = "#e1c94b";
    var color = hexToRgb(value);

    if (color) {
      MY_COLOR = color;
      preview.style.backgroundColor = value;
    }
  }

  orange.onclick = function() {
    var value = "#e06645";
    var color = hexToRgb(value);

    if (color) {
      MY_COLOR = color;
      preview.style.backgroundColor = value;
    }
  }

  teal.onclick = function() {
    var value = "#00937f";
    var color = hexToRgb(value);

    if (color) {
      MY_COLOR = color;
      preview.style.backgroundColor = value;
    }
  }

  pink.onclick = function() {
    var value = "#e55b94";
    var color = hexToRgb(value);

    if (color) {
      MY_COLOR = color;
      preview.style.backgroundColor = value;
    }
  }

  olive.onclick = function() {
    var value = "#9c8c42";
    var color = hexToRgb(value);

    if (color) {
      MY_COLOR = color;
      preview.style.backgroundColor = value;
    }
  }

  green.onclick = function() {
    var value = "#6e873f";
    var color = hexToRgb(value);

    if (color) {
      MY_COLOR = color;
      preview.style.backgroundColor = value;
    }
  }

  violet.onclick = function() {
    var value = "#633e97";
    var color = hexToRgb(value);

    if (color) {
      MY_COLOR = color;
      preview.style.backgroundColor = value;
    }
  }

  purple.onclick = function() {
    var value = "#A333C8";
    var color = hexToRgb(value);

    if (color) {
      MY_COLOR = color;
      preview.style.backgroundColor = value;
    }
  }

  brown.onclick = function() {
    var value = "#A5673F";
    var color = hexToRgb(value);

    if (color) {
      MY_COLOR = color;
      preview.style.backgroundColor = value;
    }
  }

  grey.onclick = function() {
    var value = "#3b3a39";
    var color = hexToRgb(value);

    if (color) {
      MY_COLOR = color;
      preview.style.backgroundColor = value;
    }
  }

  preview.style.backgroundColor = "#3f6e87";
  input.value = "#3f6e87";

  input.onchange = function (e) {
    var value = e.target.value;
    var color = hexToRgb(value);

    if (color) {
      MY_COLOR = color;
      preview.style.backgroundColor = value;
    }
  };
}

// Canvas methods.
function sendPixel (x, y, r, g, b) {
  CONNECTION.send([x, y, r, g, b].join(" "));
}

function setPixel (x, y, r, g, b) {
  if (!this.id) {
    this.id = CONTEXT.createImageData(1, 1);
    this.idd = this.id.data;
    this.idd[3] = 255;
  }

  this.idd[0] = r;
  this.idd[1] = g;
  this.idd[2] = b;

  CONTEXT.putImageData(this.id, x, y);
}

// Util
function hexToRgb (hex) {
  var shorthandRegex = /^#?([a-f\d])([a-f\d])([a-f\d])$/i;

  hex = hex.replace(shorthandRegex, function(m, r, g, b) {
    return r + r + g + g + b + b;
  });

  var result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex);

  return result ? {
    r: parseInt(result[1], 16),
    g: parseInt(result[2], 16),
    b: parseInt(result[3], 16)
  } : null;
}
`
)
