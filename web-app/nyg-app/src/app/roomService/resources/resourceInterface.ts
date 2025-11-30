import { Observable } from "rxjs"

export interface RoomMethods {
    joinRoom(to: string, nickname: string, friend: boolean, code: string): void
    Rooms(): Observable<any>
    // LeaveRoom(): void
    CreateRoom(roomName: string, nickname: string, limit: number, friend: boolean, code: string, setToss: boolean): void
}

interface RoomProfile {
    id: string
    join: boolean
}

