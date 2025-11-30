import { ChangeDetectorRef, Component, Input, NgZone, signal, ViewEncapsulation, WritableSignal } from '@angular/core';
import { Subject, Subscription } from 'rxjs';
import { LobbyResource } from '../../../lobby/resources/lobbyResource';
import { RoomService } from '../../../roomService/room.service';

@Component({
  selector: 'Refscorboard-25x',
  imports: [],
  encapsulation: ViewEncapsulation.Emulated,
  templateUrl: './Ref25xScoreboard.component.html',
  styleUrls: [
    './touchup/score-board.scss',
    './touchup/container.scss',
    './touchup/time.scss',
    './touchup/color.scss',
    './touchup/team.scss',
    './touchup/animation.scss'
  ],
})
export class Ref25xScoreboardComponent {
  Team = [['R', 'E', 'D'], ['B', 'L', 'U']]
  Red_ = ['R', 'E', 'D']
  Blue_ = ['B', 'L', 'U']
  Rounds = [1, 2, 3]
  NYGredScore: number = 3
  NYGblueScore: number = 3
  OCSgameTime: number = 10
  NYGset: number = 3
  NYGround: number = 6

  constructor() {
  }
  ngOnInit(): void {
  }
}