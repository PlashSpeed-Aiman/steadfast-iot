import "./style.css";
import javascriptLogo from "./javascript.svg";
import viteLogo from "/vite.svg";
import { setupCounter } from "./counter.js";
import $ from "jquery";

let socket = undefined;

$("#submit").on("click", function (event) {
  event.preventDefault();
  const message = $("#message").val();
  console.log(message);
  socket.send(message);
});
$("#stats").on("click", async function (event) {
  $("#statsResult").hide();
  event.preventDefault();
  await fetch("http://192.168.1.6:8080/stats").then(
    (response) => {
      response.text().then((text) => {
        $("#statsResult").html(
          `<div style="opacity:1">
          ${text}
          </div>`
        ).fadeIn(1000);
      
    }).catch((error) => {
      console.log(error);
      $("#statsResult").html(
        `${error}`
      );
    });;
  }).catch((error) => {
    console.log(error);
    $
    $("#statsResult").html(
      `${error}`
    );
  });;
});

$("#disconnect").on("click", function (event) {
  event.preventDefault();
  socket.close();
  $("#status").html("<h2>Connection Closed</h2>");
});
$("#connect").on("click", async function () {
  const delay = (ms) => new Promise((resolve) => setTimeout(resolve, ms));
  if (socket !== undefined) {
    socket.close();
    $("#status").html("<h2>Connection Closed</h2>");
    await delay(2000);
  }
  $("#status").html("<h2>Attempting Connection...</h2>");
  await delay(2000);
  socket = new WebSocket("ws://192.168.1.6:8080/ws");

  socket.onopen = () => {
    const message = $("#message").html;
    console.log("Successfully Connected");
    message("<h2>Successfully Connected</h2>")
    socket.send("Hi From the Client!");
    message("<h2>Connected</h2>");
  };

  socket.onclose = (event) => {
    console.log("Socket Closed Connection: ", event);
    socket.send("Client Closed!");
    $("#status").html("Closed");
  };
  socket.onmessage = (msg) => {
    console.log(msg);
    $("#messages").html(`<h2>${msg.data}</h2>`);
  };
  socket.onerror = (error) => {
    console.log("Socket Error: ", error);
    $("#status").html(`<h2>${error}</h2>`);
  };
});
