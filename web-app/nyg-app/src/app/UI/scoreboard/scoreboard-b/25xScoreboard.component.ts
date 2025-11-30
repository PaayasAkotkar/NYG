import { ChangeDetectorRef, Component, Input, NgZone, signal, ViewEncapsulation, WritableSignal } from '@angular/core';
import { Subject, Subscription } from 'rxjs';
import { LobbyResource } from '../../../lobby/resources/lobbyResource';
import { RoomService } from '../../../roomService/room.service';

@Component({
  selector: 'scoreboard-25x',
  imports: [],
  encapsulation: ViewEncapsulation.Emulated,
  templateUrl: './25xScoreboard.component.html',
  styleUrls: [
    './touchup/score-board.scss',
    './touchup/container.scss',
    './touchup/time.scss',
    './touchup/color.scss',
    './touchup/team.scss',
    './touchup/animation.scss'
  ],
})
export class Scoreboard25xComponent {
  Team = [['R', 'E', 'D'], ['B', 'L', 'U']]
  Red_ = ['R', 'E', 'D']
  Blue_ = ['B', 'L', 'U']
  PlayerAgainst = ['QUEEN', 'KING']
  Rounds = [1, 2, 3]
  @Input({ required: false }) NYGredScore: number = 1
  @Input({ required: false }) NYGblueScore: number = 1
  @Input({ required: false }) OCSgameTime: number = 10
  @Input({ required: false }) NYGset: number = 1
  @Input({ required: false }) NYGround: number = 1
  @Input({ required: false }) NYGroundOver: WritableSignal<boolean> = signal(false)
  @Input({ required: false }) NYGrewindRestart: WritableSignal<boolean> = signal(false)
  @Input({ required: false }) NYGunFreeze: WritableSignal<boolean> = signal(true)
  @Input({ required: false }) NYGuseRewindPower: WritableSignal<boolean> = signal(false)
  constructor() {
  }
  ngOnInit(): void {
  }
}