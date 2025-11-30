import { Injectable, NgZone } from '@angular/core';
import { SessionService } from '../session/session.service';
import { v4 } from 'uuid';

// the service implements the stuffs like get cookie, init storage
@Injectable({
  providedIn: 'root'
})
export class MiscellaenousService {

  constructor(private _session: SessionService, private run: NgZone) { }
  private ID: string = ""

  InitStorage() {
    this.run.run(() => {
      var ______ok = this._session.getItem('session')
      if (______ok) {
        this.ID = ______ok
      } else {
        this.ID = v4()
        this._session.setItem("session", this.ID)
      }
    })
  }
}
