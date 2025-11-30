import { ChangeDetectorRef, Component, Input, NgZone, ViewEncapsulation, WritableSignal } from '@angular/core';
import { LobbyResource } from '../../../lobby/resources/lobbyResource';
import { Subject, Subscription } from 'rxjs';
import { RoomService } from '../../../roomService/room.service';

@Component({
  selector: 'Refscoreboard-11x',
  imports: [],
  templateUrl: './Ref11xScoreboard.component.html',
  styleUrls: [
    './touchup/score-board.scss',
    './touchup/container.scss',
    './touchup/time.scss',
    './touchup/color.scss',
    './touchup/team.scss',
    './touchup/animation.scss'
  ],
  encapsulation: ViewEncapsulation.Emulated
})
export class Ref11xScoreboardComponent {
  Team = [['R', 'E', 'D'], ['B', 'L', 'U']]
  Red_ = ['R', 'E', 'D']
  Blue_ = ['B', 'L', 'U']
  Rounds = [1, 2, 3]

  NYGredScore: number = 3
  NYGblueScore: number = 3
  OCSgameTime: number = 10
  NYGset: number = 3
  NYGround: number = 6

  ngOnInit(): void { }
}
