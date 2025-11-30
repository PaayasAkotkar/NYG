import { ChangeDetectorRef, Component, Input, NgZone, OnDestroy, OnInit, Output, ViewEncapsulation, EventEmitter, WritableSignal, signal } from '@angular/core';
import { LobbyResource } from '../../lobby/resources/lobbyResource';
import { RoomService } from '../../roomService/room.service';
import { interval, Subject, Subscription, takeUntil, timer } from 'rxjs';
import { Scoreboard25xComponent } from "./scoreboard-b/25xScoreboard.component";
import { Scoreboard11xComponent } from './scoreboard-a/11xScoreboard.component';
import { GymClockComponent } from "../clocks/gym-clock/gym-clock.component";
import { MysteriumClockComponent } from "../clocks/mysterium-clock/mysterium-clock.component";
import { FastlaneClockComponent } from "../clocks/fastlane-clock/fastlane-clock.component";
import { PhysicsClockComponent } from "../clocks/physics-clock/physics-clock.component";
import { NodeClockComponent } from "../clocks/node-clock/node-clock.component";

@Component({
  selector: 'app-scoreboard',
  imports: [Scoreboard25xComponent, Scoreboard11xComponent, GymClockComponent, MysteriumClockComponent, FastlaneClockComponent, PhysicsClockComponent, NodeClockComponent],
  templateUrl: './scoreboard.component.html',
  styleUrls: [
    './touchup/score-board.scss',
    './touchup/container.scss',
    './touchup/time.scss',
    './touchup/color.scss',
    './touchup/team.scss',
  ],
  encapsulation: ViewEncapsulation.Emulated
})

export class ScoreboardComponent implements OnInit {
  Team = [['R', 'E', 'D'], ['B', 'L', 'U']]
  Red_ = ['R', 'E', 'D']
  Blue_ = ['B', 'L', 'U']
  PlayerAgainst = ['QUEEN', 'KING']
  Rounds = [1, 2, 3]

  @Input({ required: true }) OCStheme: string = "text25x"
  @Input({ required: false }) NYGredScore: number
  @Input({ required: false }) NYGblueScore: number
  @Input({ required: false }) NYGset: number
  @Input({ required: false }) NYGround: number
  @Input({ required: false }) NYGroundOver: WritableSignal<boolean> = signal(false)
  @Input({ required: false }) OCSgameTime: number
  @Input({ required: false }) NYGrewindRestart: WritableSignal<boolean> = signal(false)
  @Input({ required: false }) NYGunFreeze: WritableSignal<boolean> = signal(false)
  @Input({ required: false }) NYGuseRewindPower: WritableSignal<boolean> = signal(false)
  @Input({ required: false }) NYGisGuessDone: WritableSignal<boolean> = signal(false)
  @Output() NYGTime: EventEmitter<WritableSignal<number>> = new EventEmitter<WritableSignal<number>>()
  @Output() NYGtimeout: EventEmitter<boolean> = new EventEmitter<boolean>()
  @Input({ required: true }) NYGstartClock: WritableSignal<boolean>
  ngOnInit(): void { }

  Listen(time: number) { this.NYGTime.emit() }
}