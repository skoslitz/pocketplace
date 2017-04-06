package main

const (
	indexHtml = `<html>
  <head>
    <title>pocketplace</title>
    <style>
      body {
        margin: 0;
      }

      #canvas {
        transform-origin: 0 0;
        image-rendering: pixelated;
        cursor: crosshair;
        display: none;
      }

      #loading {
        display: block;
      }

      #picker {
        position: fixed;
        width: 100px;
        top: 10px;
        left: 50%;
        margin-left: -50px;
        display: none;
      }

      #blueColor {
        position: fixed;
        top: 10px;
        left: 10%;
        margin-left: -10px;
      }

      #redColor {
        position: fixed;
        top: 10px;
        left: 20%;
        margin-left: -10px;
      }

      #yellowColor {
        position: fixed;
        top: 10px;
        left: 25%;
        margin-left: -10px;
      }

      #picker input {
        box-sizing: border-box;
        padding-right: 20px;
        width: 100%;
      }

      #picker div {
        box-sizing: border-box;
        display: inline-block;
        height: 19px;
        position: absolute;
        right: 1px;
        top: 1px;
        width: 19px;
      }

      #cooldown {
        background-color: black;
        color: white;
        display: none;
        font-size: 2em;
        position: fixed;
        right: 0;
        top: 0;
      }
    </style>
    <script src="/options.js"></script>
    <script src="/frontend.js"></script>
  </head>
  <body>
    <div id="loading">Loading ...</div>
    <canvas id="canvas"></canvas>
    <button id="blueColor">BLAU</button>
    <button id="redColor">ROT</button>
    <button id="yellowColor">GELB</button>
    <div id="picker">
      <input id="picker-input" type="text" name="color" placeholder="color">
      <div id="picker-preview"></div>
    </div>
    <div id="cooldown"></div>
  </body>
</html>
`
)
