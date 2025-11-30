import { Component, effect, Input, OnInit, ViewEncapsulation, WritableSignal } from '@angular/core';
import { MatGridListModule } from '@angular/material/grid-list';
import { LobbyResource, Rounds, SheetUpdate } from '../lobby/resources/lobbyResource';
import { RoomService } from '../roomService/room.service';

@Component({
  selector: 'app-cheatsheet',
  imports: [MatGridListModule],
  templateUrl: './cheatsheet.component.html',
  styleUrls: ['./touchup/cheatsheet.scss',
    './touchup/color.scss'
  ],
  encapsulation: ViewEncapsulation.Emulated,
})
export class CheatsheetComponent implements OnInit {
  Sets = ['S1', 'S2', 'S3']

  constructor() { }

  ngOnInit(): void { }

  @Input({ required: false }) NYGplayersNickNames = ['RED', 'BLU']

  // nickname and stats
  @Input({ required: false }) NYGupdateSheet: Record<string, Record<number, string>> = {}
  @Input({ required: false }) NYGredScore = 1
  @Input({ required: false }) NYGblueScore = 1


}
