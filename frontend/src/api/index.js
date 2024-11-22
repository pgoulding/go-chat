var socket =  new WebSocket("ws://localhost:4000")

let connect = cb => {
    console.log("connecting...")

    socket.onopen = () => {
        console.log("Succesfully Connected!")
    }

    socket.onmessage = msg => {
        console.log(msg)
    }

    socket.onclose = event => {
        console.log("Socket closed connection: ", event)
    }

    socket.onerror = err => {
        console.error("Socket Error: ", err)
    }
}

let sendMsg = msg => {
    console.log("Sending Message: ", msg)
    socket.send(msg)
}

export {connect, sendMsg}