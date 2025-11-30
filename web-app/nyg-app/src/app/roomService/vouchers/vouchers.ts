import { BehaviorSubject, } from "rxjs";
import { v4, } from "uuid";
import { SessionService } from "../../storage/session/session.service";
import { LocalService } from "../../storage/local/local.service";

// restriction one making the room named as private-room
export class RoomFields {
    protected receiveTokens: BehaviorSubject<any> = new BehaviorSubject<any>('')
    protected roomName: string = "private-room"
    public id = v4()
    public getId = ""
    // protected URL: string = `ws://localhost:3000/ws/${this.id}/${this.roomName}`
    protected URL: string
    protected Conn: WebSocket
    protected signUPDetails = {
        joinRoom: false, to: "", leaveRoom: false, createRoom: false,
        roomName: "", roomCapacity: 0,
        book: "", category: "", field: "",
    }

    constructor(private session: SessionService, private local: LocalService) {

    }
}

