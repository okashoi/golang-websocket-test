const socket = new WebSocket('ws://localhost:8080/');

socket.addEventListener('open', function (event) {
    const message = {msg: 'hello world!'};
    socket.send(JSON.stringify(message));
});

socket.addEventListener('message', function (event) {
    console.log('Message from server', event.data);
});
