import { Injectable, NgZone } from '@angular/core';
import { LoginResources } from './resources/loginResources';
import { HttpClient } from '@angular/common/http';
import { SessionService } from '../storage/session/session.service';
import { timer } from 'rxjs';
import { UpdateLivelyService } from '../update-lively/update-lively.service';

@Injectable({
  providedIn: 'root'
})
export class LoginService extends LoginResources {

  constructor(private _R: NgZone, private _http: HttpClient, private ses: SessionService) {

    var id: string = ""
    var found = ses.getItem('session')
    if (found) {
      id = found
      ses.setItem('login', 'yes')
    }
    super(_R, _http, id,)
    UpdateLivelyService
  }
}
