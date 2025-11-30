import { Component, Input, OnInit, ViewEncapsulation } from '@angular/core';
import { MatGridListModule } from '@angular/material/grid-list';
import { LobbyResource } from '../lobby/resources/lobbyResource';
import { RoomService } from '../roomService/room.service';
import { RoundKeys } from '../lobby/vouchers/lobbyVouchers';

@Component({
  selector: 'app-frame',
  imports: [MatGridListModule],
  templateUrl: './frame.component.html',
  styleUrls: ['./touchup/frame.scss', './touchup/color.scss'],
  encapsulation: ViewEncapsulation.Emulated

})
export class FrameComponent implements OnInit {
  Rounds: string[] = ['R1', 'R2', 'R3', 'R4', 'R5', 'R6', 'TEAM', 'SCORE']
  Rounds2: string[] = ['R1', 'R2', 'R3', 'R4', 'R5', 'R6']
  constructor(private room: RoomService) {

  }

  ngOnInit(): void {
  }
  @Input() AS: number
  @Input() HS: number
  class = [
    'nyg-item1',
    'nyg-item2',
    'nyg-item3',
    'nyg-item4',
    'nyg-item5',
    'nyg-item6',
  ]
  @Input() NYGupdateBlue: Map<string, string | undefined>
  @Input() NYGupdateRed: Map<string, string | undefined>
}
