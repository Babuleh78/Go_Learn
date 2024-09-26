const conn = new WebSocket(`ws://${window.location.host}/ws`);
        const messages = document.getElementById("messages");
        const messageInput = document.getElementById("messageInput");
        const sendMessageButton = document.getElementById("sendMessage");

        conn.onmessage = function(event) {
            const msg = JSON.parse(event.data);
            const li = document.createElement("li");
            li.textContent = msg.text;
            messages.appendChild(li);
        };

sendMessageButton.onclick = function() {
    const msg = { text: messageInput.value };
    conn.send(JSON.stringify(msg));
    messageInput.value = ""; // очищаем поле ввода
};