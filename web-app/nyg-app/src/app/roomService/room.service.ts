import { Inject, Injectable, NgZone } from '@angular/core';
import { RoomAgent } from './agent/agent';
import { SessionService } from '../storage/session/session.service';
import { LocalService } from '../storage/local/local.service';
import { CookieService } from '../storage/cookie/cookie.service';

@Injectable({
  providedIn: 'root'
})
export class RoomService extends RoomAgent {
  constructor(private ___cookie: CookieService, private ___session: SessionService, private ___local: LocalService, private __run: NgZone) {
    super(___cookie, ___session, ___local, __run)
  }
}
