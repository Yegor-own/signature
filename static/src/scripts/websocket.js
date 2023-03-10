function connectWebsocket() {
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws")

        conn.onmessage = function(ev) {
            data = ev.data
            console.log(data);

            msg = document.createElement('p').innerText(data)

            document.getElementById("#log").appendChild(msg)
            
        }

        document.getElementById("send").onsubmit = function() {
            if (!conn) {
                return false
            }
            if (!document.getElementById("message").value) {
                return false
            }
            conn.send(document.getElementById("message").value)
            return flase
        }
    }
}




window.onload = function() {
    connectWebsocket()
}