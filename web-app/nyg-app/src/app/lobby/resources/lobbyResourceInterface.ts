import { RoomService } from "../../roomService/room.service";

export interface LobbyMethods {
    create(name: string, limit: number, friend: boolean): void
    join(name: string, friend: boolean): void
    leave(name: string): void

}

