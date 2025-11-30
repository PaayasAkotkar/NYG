import { Inject, Injectable, NgZone, PLATFORM_ID } from '@angular/core';
import { ProfileMethods } from './resource/profileResource';
import { HttpClient } from '@angular/common/http';
import { SessionService } from '../storage/session/session.service';

@Injectable({
  providedIn: 'root'
})
export class ProfileService extends ProfileMethods {

  constructor(private ses: SessionService, _http: HttpClient, _n: NgZone, @Inject(PLATFORM_ID) _PlatformID: Object) {
    var found = ses.getItem('session')
    var id_: string = ""
    if (found) {
      id_ = found
    }
    super(id_, _http, _n, _PlatformID)
  }
}
