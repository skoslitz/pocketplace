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
        text-align: center;
        margin-top: 50px;
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
        width: 30px;
        height: 30px;
        background-color: #3f6e87;
        top: 10px;
        margin-left: 910px;
      }

      #redColor {
        position: fixed;
        width: 30px;
        height: 30px;
        background-color: #8e225c;
        top: 45px;
        margin-left: 910px;
      }

      #yellowColor {
        position: fixed;
        width: 30px;
        height: 30px;
        background-color: #e1c94b;
        top: 80px;
        margin-left: 910px;
      }

      #orangeColor {
        position: fixed;
        width: 30px;
        height: 30px;
        background-color: #e06645;
        top: 115px;
        margin-left: 910px;
      }

      #tealColor {
        position: fixed;
        width: 30px;
        height: 30px;
        background-color: #00937f;
        top: 150px;
        margin-left: 910px;
      }

      #pinkColor {
        position: fixed;
        width: 30px;
        height: 30px;
        background-color: #e55b94;
        top: 185px;
        margin-left: 910px;
      }

      #oliveColor {
        position: fixed;
        width: 30px;
        height: 30px;
        background-color: #9c8c42;
        top: 220px;
        margin-left: 910px;
      }

      #greenColor {
        position: fixed;
        width: 30px;
        height: 30px;
        background-color: #6e873f;
        top: 255px;
        margin-left: 910px;
      }

      #violetColor {
        position: fixed;
        width: 30px;
        height: 30px;
        background-color: #633e97;
        top: 290px;
        margin-left: 910px;
      }

      #purpleColor {
        position: fixed;
        width: 30px;
        height: 30px;
        background-color: #A333C8;
        top: 325px;
        margin-left: 910px;
      }

      #brownColor {
        position: fixed;
        width: 30px;
        height: 30px;
        background-color: #A5673F;
        top: 360px;
        margin-left: 910px;
      }

      #greyColor {
        position: fixed;
        width: 30px;
        height: 30px;
        background-color: #3b3a39;
        top: 395px;
        margin-left: 910px;
      }

      #picker input {
        box-sizing: border-box;
        padding-right: 20px;
        width: 100%;
        display: none;
      }

      #picker div {
        box-sizing: border-box;
        display: inline-block;
        height: 25px;
        position: absolute;
        right: 1px;
        top: 1px;
        width: 25px;
      }

      #picker-preview {
        border-style: solid;
        border-color: black;
        border-width: 2px;
      }

      #textarea {
        margin-left: 1000px;
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
    <div id="blueColor"></div>
    <div id="redColor"></div>
    <div id="yellowColor"></div>
    <div id="orangeColor"></div>
    <div id="tealColor"></div>
    <div id="pinkColor"></div>
    <div id="oliveColor"></div>
    <div id="greenColor"></div>
    <div id="violetColor"></div>
    <div id="purpleColor"></div>
    <div id="brownColor"></div>
    <div id="greyColor"></div>

    <div id="picker">
      <input id="picker-input" type="text" name="color" placeholder="color">
      <div id="picker-preview"></div>
    </div>
    <div id="textarea">
      <h2>NIMIRUM PIXELATED</h2>
    </div>

    <div id="cooldown"></div>
  </body>
</html>
`
)
