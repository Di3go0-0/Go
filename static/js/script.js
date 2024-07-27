function changeMessage() {
  const messageElement = document.getElementById("message");
  const currentMessage = messageElement.innerHTML;

  const alternateMessage = [
    "Hello, World!",
    "Hello, Universe!",
    "Hello, Again!",
  ];

  let newMessage;

  do {
    newMessage =
      alternateMessage[Math.floor(Math.random() * alternateMessage.length)];
  } while (newMessage === currentMessage);

  messageElement.textContent = newMessage;
}
