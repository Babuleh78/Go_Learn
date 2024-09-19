const messages = document.getElementById('messages');
const messageInput = document.getElementById('messageInput');
const sendMessageButton = document.getElementById('sendMessage');

const socket = new WebSocket('ws://localhost:8080/ws');

socket.onmessage = function(event) {
    const message = JSON.parse(event.data);
    const messageElement = document.createElement('div');
    messageElement.textContent = message.text;
    messages.appendChild(messageElement);
};

sendMessageButton.onclick = function() {
    const message = { text: messageInput.value };
    socket.send(JSON.stringify(message));
    messageInput.value = '';
};
