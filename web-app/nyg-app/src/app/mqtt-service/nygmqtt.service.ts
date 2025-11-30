import { Injectable } from '@angular/core';
import { MqttMethods } from './resources/mqttResource';
import { HttpClient } from '@angular/common/http';
import { SessionService } from '../storage/session/session.service';

@Injectable({
  providedIn: 'root'
})
export class NYGmqttService extends MqttMethods {

  constructor(private _http: HttpClient, private sess: SessionService) {
    super(_http, sess)
  }
}
