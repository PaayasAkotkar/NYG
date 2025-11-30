import { Component, Output, EventEmitter, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DeckLayoutComponent } from "../deck-layout/deck-layout.component";
import { OcsXxBtnComponent } from "../ocs-xx-btn/ocs-xx-btn.component";
import { Location } from '@angular/common';
import { RoomService } from '../../roomService/room.service';

@Component({
  selector: 'app-clash-lobby',
  imports: [OcsXxBtnComponent],
  templateUrl: './clash-lobby.component.html',
  styleUrls: ['./touchup/clash.scss', './touchup/color.scss', './touchup/animations.scss'],
})
export class ClashLobbyComponent implements OnInit {
  @Output() OCSListen = new EventEmitter<void>()
  cancel() {
    this.OCSListen.emit()
    this._r.navigateByUrl('/Clash')
  }

  constructor(private room: RoomService, private _r: Router, private location: Location) { }
  ngOnInit(): void {

  }
}
