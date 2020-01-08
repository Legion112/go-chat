/**
 *  @namespace
 *  @property {WebSocket} webSocket
 */
const SocketService = {
    webSocket: null,
    init(webSocket){
        this.webSocket = webSocket;
    },
    sendDate(data){
        this.webSocket.send(JSON.stringify(data));
    }
}
export default SocketService