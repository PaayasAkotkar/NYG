import { ChangeDetectorRef, Component, Input, NgZone, ViewEncapsulation, WritableSignal } from '@angular/core';
import { LobbyResource } from '../../../lobby/resources/lobbyResource';
import { Subject, Subscription } from 'rxjs';
import { RoomService } from '../../../roomService/room.service';

@Component({
  selector: 'scoreboard-11x',
  imports: [],
  templateUrl: './11xScoreboard.component.html',
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
export class Scoreboard11xComponent {
  Team = [['R', 'E', 'D'], ['B', 'L', 'U']]
  Red_ = ['R', 'E', 'D']
  Blue_ = ['B', 'L', 'U']
  PlayerAgainst = ['QUEEN', 'KING']
  Rounds = [1, 2, 3]

  @Input({ required: true }) NYGredScore: number
  @Input({ required: true }) NYGblueScore: number
  @Input({ required: true }) NYGset: number
  @Input({ required: true }) NYGround: number
  @Input({ required: true }) NYGroundOver: WritableSignal<boolean>
  @Input({ required: true }) OCSgameTime: number
  @Input({ required: true }) NYGrewindRestart: WritableSignal<boolean>
  @Input({ required: true }) NYGunFreeze: WritableSignal<boolean>
  @Input({ required: true }) NYGuseRewindPower: WritableSignal<boolean>

  ngOnInit(): void { }
}
