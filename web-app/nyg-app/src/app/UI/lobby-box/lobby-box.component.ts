import { Component, Input, OnInit, signal, WritableSignal } from '@angular/core';
import { IProfile_, JoinedProfile } from '../../lobby/resources/lobbyResource';
import { OcsMiniviewService } from '../ocs-miniview/ocs-miniview.service';
import { BehaviorSubject } from 'rxjs';

@Component({
  selector: 'app-lobby-box',
  imports: [],
  templateUrl: './lobby-box.component.html',
  styleUrls: ['./touchup/lb.scss', './touchup/color.scss'],
})
export class LobbyBoxComponent implements OnInit {
  @Input({ required: true }) OCSprofiles: JoinedProfile[]
  show: boolean = false

  ViewProfile(profile: IProfile_) {
    this.c.Setup(profile)
  }
  TimeOut() {
    this.c.closeSetup()
  }
  ngOnInit(): void { }
  constructor(private c: OcsMiniviewService) { }
}
