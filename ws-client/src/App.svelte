<script lang="ts">
  // This line connects to the server
  let socket = new WebSocket("ws://localhost:8888/ws");

  let message = "";

  // This is called when websocket is opened
  socket.onopen = function (e) {
    console.log("[open] Connection established");
    console.log("Sending to server");
    socket.send("My name is John");
  };

  socket.onmessage = function (event) {
    console.log(`[message] Data received from server: ${event.data}`);
    message = event.data;
  };

  socket.onclose = function (event) {
    if (event.wasClean) {
      console.log(
        `[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`
      );
    } else {
      // e.g. server process killed or network down
      // event.code is usually 1006 in this case
      console.log("[close] Connection died");
    }
  };

  socket.onerror = function (error) {
    alert(`[error]`);
  };
</script>

<main>
  <div>{message}</div>
</main>
