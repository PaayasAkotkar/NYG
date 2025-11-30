import { Subscription } from "rxjs";
import { RoomService } from "../../../roomService/room.service";
import { ChangeDetectorRef, signal, WritableSignal } from "@angular/core";

export interface RoomsPattern {
    RoomName: string
    Category: string
    Type: string
    Book: string
}

export class RoomListAgent {
    watch = new Subscription()
    RoomsList: WritableSignal<RoomsPattern[]> = signal([])
    _ref: RoomsPattern[] = []
    _proceed: boolean = false

    constructor(private room: RoomService, private cdr: ChangeDetectorRef) {
        console.log("connefwf")
        this.room.connect()
        this.watch = this.room.Rooms().subscribe({
            next: (token) => {
                const RoomList = /RoomLists:(?=[\s\S\w])/
                const getRooms = RoomList.test(token)
                console.log("token: ", token)
                switch (true) {
                    case getRooms:
                        var parse = "RoomLists:"
                        var forate: RoomsPattern = { RoomName: "", Category: "", Type: "", Book: "" }
                        var _get = this.afterColonJson(token, parse, forate)
                        this._ref.push(_get)
                        this.RoomsList.set(_get)
                        this._ref = []
                        if (this.RoomsList().length > 0)
                            this._proceed = true
                        this.cdr.detectChanges()
                        console.log("len: ", this.RoomsList.length)
                        console.log('faff: ', this.RoomsList()[0])
                        console.log("rooms fetched: ", this.RoomsList())
                        break
                }
            }
        })
    }

    afterColonJson(token: string, parse: string, _interface: any) {
        token = token.substring(parse.length)
        _interface = JSON.parse(token)
        return _interface
    }
}