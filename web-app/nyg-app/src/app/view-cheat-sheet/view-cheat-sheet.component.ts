import { Component, Input, ViewEncapsulation } from '@angular/core';
import { ViewCSheetService } from './view-Csheet.service';
import { ViewCheatsService } from './view-cheats.service';
import { LobbyResource } from '../lobby/resources/lobbyResource';
import { RoomService } from '../roomService/room.service';
import { Tags2Component } from "../UI/tags2/tags2.component";

@Component({
  selector: 'app-view-cheat-sheet',
  imports: [Tags2Component],
  templateUrl: './view-cheat-sheet.component.html',
  styleUrls: ['./touchup/view-cheat-sheet.scss'],
  encapsulation: ViewEncapsulation.Emulated,
})
export class ViewCheatSheetComponent {
  Spell = ['C', 'H', 'E', 'A', 'T', 'S', 'H', 'E', 'E', 'T']
  show: boolean = false

  @Input({ required: false }) NYGplayersNickname: string[] = ['KING', 'QUEEN']

  toggle() {
    this.show = !this.show
    if (this.show) {
      this.show = false
    } else {
      this.service.closeSetup()
    }
  }
  constructor(private service: ViewCheatsService, private room: RoomService) {

  }
}