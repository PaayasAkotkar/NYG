import { BehaviorSubject, catchError, concatMap, delayWhen, distinctUntilChanged, interval, map, Observable, of, Subject, switchMap, take, takeUntil, tap, timer } from "rxjs";
import { PartyChatFields } from "../vouchers/voucher";
import { HttpClient, HttpHeaderResponse, HttpHeaders } from "@angular/common/http";
import { CompletionTriggerKind } from "typescript";
import { OnDestroy } from "@angular/core";


interface ChatReq {
    roomName: string
    nickName: string
    id: string
    token: string
    chat: boolean
}
interface RegisterRoomReq {
    roomName: string
    id: string
    register: boolean
}


export class PartyChatMethods extends PartyChatFields {
    constructor() {
        super()
    }
    Connect(RoomName: string) {
        var ___id = this.genId

        const URL = `ws://localhost:9000/ws/${___id}/party-chat`
        if (!this.Conn || this.Conn.readyState === WebSocket.CLOSED) {
            this.Conn = new WebSocket(URL)
        } else {
            console.log("error creating socket")
        }
        this.Conn.onopen = () => {
            console.log("_____SOCKET_CONNECTED_____")
            // var tokens = ["Specification", "Hosts",]
            // for (let i = 0; i < tokens.length; i++) {
            //     this.Conn.send(tokens[i])
            // }
            var token: RegisterRoomReq =
                { roomName: RoomName, id: ___id, register: true }
            this.Conn.send(JSON.stringify(token))
        }
        this.Conn.onmessage = (token) => {
            this.receiveTokens.next(token.data)
        }
        this.Conn.onerror = (e) => {
            console.log("ERROR___SCHEMA_____ WEBSOCKET")
        }
        this.Conn.onclose = () => {
            console.log("socket close")
        }
    }
    RoomChat(): Observable<string> {
        return this.receiveTokens.asObservable()
    }
    Chat(token: string, roomId: string, nickName: string) {
        var t: ChatReq = { nickName: nickName, roomName: roomId, id: this.genId, token: token, chat: true }
        this.____PC = t
        const sendToken = JSON.stringify(t)
        this.Conn.send(sendToken)
    }


}