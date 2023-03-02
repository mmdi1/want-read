export class Ws {
    link: WebSocket
    constructor(ws: string) {
        this.link = new WebSocket(ws);
        this.link.onopen = () => {
        }
        this.link.onclose = () => {

        }
    }
    send(data: any) {
        this.link.send(data)
    }
}