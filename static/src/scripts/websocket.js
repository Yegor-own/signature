function sendMessage() {
    console.log("send message initialised");
    let newMsg = document.getElementById("message")
    // connectWebsocket()
    if (newMsg != null && newMsg.value != "") {
        console.log("send message" + newMsg.value)
        // console.log("send message")
        conn.send(newMsg.value)
    }
    return false
}
    

// let conn
window.onload = function() {
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws")


        conn.onopen = function() {
            console.log("connection opened");
        }

        conn.onmessage = function(ev) {
            data = ev.data
            console.log("get message" + data);

            // msg = document.createElement("p").innerText(data)

            log = document.getElementById("log")
            document.getElementById("log").innerHTML = log.value + '\n' + data
            
        }
    }
    document.getElementById("send-message").onsubmit = sendMessage
}