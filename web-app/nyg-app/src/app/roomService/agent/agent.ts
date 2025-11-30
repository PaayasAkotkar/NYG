import { NgZone } from "@angular/core";
import { LocalService } from "../../storage/local/local.service";
import { SessionService } from "../../storage/session/session.service";
import { RoomResource } from "../resources/resource";
import { CookieService } from "../../storage/cookie/cookie.service";

export class RoomAgent extends RoomResource {
    constructor(private __cookie: CookieService, private __session: SessionService, private __local: LocalService, private _run: NgZone) {
        super(__cookie, __session, __local, _run)
    }
}