import { Component, Input, ViewEncapsulation } from '@angular/core';
import { LftOverService } from './services/setup-lftover.service';
import { ViewCSheetService } from '../view-cheat-sheet/view-Csheet.service';
import { RoomService } from '../roomService/room.service';
import { LobbyResource } from '../lobby/resources/lobbyResource';

@Component({
  selector: 'app-setup-sheet',
  imports: [],
  templateUrl: './setup-sheet.component.html',
  styleUrls: ['./touchup/setup-sheet.scss'],
  encapsulation: ViewEncapsulation.Emulated
})
export class SetupSheetComponent {

  score: string = 'O'
  Noscore: string = 'X'
  Players: string[] = ['KING', 'QUEEN', 'PRINCE', 'PRINCESS']
  Cshow: boolean = false
  Lshow: boolean = false
  @Input() NYGgetUpdate: Map<string, Map<string, string>>
  @Input() NYGplayersnickname: string[]
  // LftOver=left over

  viewSheet() {
    this.Cshow = !this.Cshow
    if (this.Cshow) {
      this.Cshow = false
      this.service.Setup(this.NYGgetUpdate, this.NYGplayersnickname)
    } else {
      this.service.closeSetup()
    }
  }

  viewLftOver() {
    this.Lshow = !this.Lshow
    if (this.Lshow) {
      this.Lshow = false
      this.lft.Setup()
    } else {
      this.service.closeSetup()
    }
  }

  constructor(private room: RoomService, private service: ViewCSheetService, private lft: LftOverService) {
  }
}
