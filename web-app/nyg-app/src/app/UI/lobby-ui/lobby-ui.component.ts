import { Component, ViewEncapsulation } from '@angular/core';
import { RoomService } from '../../roomService/room.service';
import { LobbyResource } from '../../lobby/resources/lobbyResource';
import { PlayMenuComponent } from "../play-menu/play-menu.component";

@Component({
  selector: 'app-lobby-ui',
  imports: [PlayMenuComponent],
  templateUrl: './lobby-ui.component.html',
  encapsulation: ViewEncapsulation.Emulated
})
export class LobbyUiComponent {
  constructor(private room: RoomService) {
  }

  toJ(e: boolean) {

  }
  toC(e: boolean) {

  }
}
