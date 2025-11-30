import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { UpdateLivelyResource } from './resource/update-livelyResource';
import { SessionService } from '../storage/session/session.service';

@Injectable({
  providedIn: 'root'
})
export class UpdateLivelyService extends UpdateLivelyResource {

  constructor(private __htpp: HttpClient, private sess: SessionService) {
    var id = ""
    var found = sess.getItem('session')
    if (found) {
      id = found
    }
    super(id, __htpp)
  }
}
