import { BehaviorSubject } from "rxjs"
import { v4 } from "uuid"




export class PartyChatFields {
    protected ____PC = { roomName: "", id: "", token: "" }
    protected Conn: WebSocket
    protected getId: string = ""
    protected genId: string = v4()
    protected URL = `http://localhost:9000/ws/${this.genId}/party-chat`
    protected receiveTokens: BehaviorSubject<string> = new BehaviorSubject<string>("")
}