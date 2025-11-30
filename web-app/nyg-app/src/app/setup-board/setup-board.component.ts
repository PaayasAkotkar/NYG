import { Component, EventEmitter, effect, Input, OnInit, Output, signal, ViewEncapsulation, WritableSignal, OnChanges, SimpleChanges } from '@angular/core';
import { MatGridListModule } from '@angular/material/grid-list';
import { Pairing, Rounds, SheetUpdate, SpectateInfo } from '../lobby/resources/lobbyResource';
import { RoomService } from '../roomService/room.service';
import { MatCardModule } from '@angular/material/card';
import { CountComponent } from "../UI/count/count.component";
import { TagsComponent } from "../UI/tags/tags.component";

import { Tags2Component } from "../UI/tags2/tags2.component";
import { CheatService } from '../cheatsheet/cheat.service';
import { PowerService } from '../UI/power-up/power.service';
import { ListComponent } from "../UI/list/list.component";
import { OCSFieldFormat } from '../helpers/patterns';
import { FrameService, } from '../UI/frame/frame.service';
import { PowerTagService } from '../UI/power-up/power-tag/power-tag.service';
import { BarComponent } from "../UI/bar/bar.component";
import { MysteriumClockComponent } from "../UI/clocks/mysterium-clock/mysterium-clock.component";
import { GymClockComponent } from "../UI/clocks/gym-clock/gym-clock.component";
import { OnModeDirective } from '../directives/on-mode.directive';
import { ColorPalletes as theme } from '../helpers/hand';
import { FastlaneClockComponent } from "../UI/clocks/fastlane-clock/fastlane-clock.component";
import { PhysicsClockComponent } from "../UI/clocks/physics-clock/physics-clock.component";
import { NodeClockComponent } from "../UI/clocks/node-clock/node-clock.component";
import { LocifyClockComponent } from "../UI/clocks/locify-clock/locify-clock.component";
import { FireComponent } from "../UI/fire/fire.component";
import { NYGlistService } from '../UI/list/nyglist.service';
import { BehaviorSubject } from 'rxjs';

@Component({
  selector: 'app-setup-board',
  imports: [OnModeDirective, MatGridListModule, MatCardModule, MatGridListModule, CountComponent, TagsComponent, Tags2Component, ListComponent, BarComponent, MysteriumClockComponent, GymClockComponent, FastlaneClockComponent, PhysicsClockComponent, NodeClockComponent, LocifyClockComponent, FireComponent],
  templateUrl: './setup-board.component.html',
  styleUrls: ['./touchup/setup-board.scss', './touchup/color.scss'],
  encapsulation: ViewEncapsulation.Emulated

})
export class SetupBoardComponent implements OnInit {
  fromColor: string = ''
  toColor: string = ''
  bg: string = ''
  @Input({ required: true }) NYGspectate: boolean = false
  @Input({ required: true }) NYGlock: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGround: number = 0
  @Input({ required: true }) NYGonFire: number = 0
  @Input({ required: true }) NYGClashProfile: Record<string, Pairing> = {}
  CollectKeys: string[] = [
    // 'king', 'queen', 'prince', 'princess'
  ]
  @Input({ required: true }) NYGspectateInfo: SpectateInfo
  @Output() NYGtimeOut: EventEmitter<boolean> = new EventEmitter<boolean>()
  timeOut(e: boolean) {
    console.log('timeout!!! ⏲️')
    this.NYGtimeOut.emit(e)
  }

  watch: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false)
  constructor(private ls: NYGlistService, private tg: PowerTagService, private left: FrameService, private room: RoomService, private sco: CheatService, private po: PowerService) {
    // effect(() => {
    //   const tag = this.NYGuseTag()
    //   if (tag) {
    //     this.tg.Setup()
    //   } else {
    //     this.tg.closeSetup()
    //   }
    // })
  }

  viewStats() {
    this.sco.Setup(
      this.NYGplayersNickName, this.NYGCupdate, this.HS, this.AS
    )
  }
  @Input({ required: true }) NYGrating: number = 0
  @Input({ required: true }) NYGFrames: WritableSignal<OCSFieldFormat[]>

  viewLeftOver() {
    this.left.Setup(this.NYGFrames)
  }
  viewList() {
    this.ls.Setup(this.NYGlist, "LIST", this.NYGMode(), this.NYGLproceed)
  }
  viewPowers() {
    this.po.Setup(this.NYGmyroom, this.isNYGClash, this.NYGnexus, this.NYGrewind, this.NYGfreeze, this.NYGdraw, this.NYGtag, this.NYGbet, this.NYGcovert, this.NYGblock, this.NYGincludeR, this.NYGincludeD, this.NYGincludeT, this.NYGincludeN, this.NYGincludeF, this.NYGincludeC, this.NYGincludeB,
      this.NYGcanUsePower
    )
  }

  nHeader = ['POWER', 'SCOREBOARD'].sort()
  ClashEntries: Record<string, Pairing> = {}
  InitClashProfile() {
    // sort in descending order
    const _ref = Object.entries(this.NYGClashProfile)
    _ref.sort(([, a], [, b]) => a.chances - b.chances)
    for (let nicknames of Object.keys(this.NYGClashProfile)) {
      this.CollectKeys.push(nicknames)
    }
    // empty
    for (let key of Object.keys(this.NYGClashProfile)) {
      delete (this.NYGClashProfile[key])
    }
    for (let [key, val] of _ref) {
      this.NYGClashProfile[key] = val
    }
  }
  ngOnInit(): void {

    this.InitClashProfile()
    switch (this.NYGMode()) {
      case 'mysterium':
        this.bg = theme.$nyg_brown_palletes.ad8_brown
        break
      case 'node':
        this.bg = theme.$nyg_blue_palletes.ad16_blue
        break
      case 'fastlane':
        this.bg = theme.$nyg_green_palletes.ad19_green
        break
      case 'pshycic':
        this.bg = theme.$nyg_voilet_palletes.ad9_voilet
        break
      case 'gym':
        this.bg = theme.$nyg_black_palletes.ad8_black
        break
      case 'locify':
        this.bg = theme.$nyg_orange_palletes.ad9_orange
        break
      default:
        this.NYGMode.set('')
        this.bg = 'none'

    }
  }

  @Input({ required: true }) NYGChances: WritableSignal<number>
  @Input({ required: true }) NYGuseTag: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) isNYGClash: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGstartClock: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGlist: WritableSignal<OCSFieldFormat[]> = signal([])
  @Input({ required: true }) NYGLproceed: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGMode: WritableSignal<string> = signal('')
  @Input({ required: true }) NYGroomMode: string = "2v2"
  @Input({ required: true }) NYGgameClock: number = 10
  @Input({ required: true }) NYGdecisionTime: number = 20

  @Input({ required: true }) Conns: WritableSignal<number>
  @Input({ required: true }) AS: number
  @Input({ required: true }) HS: number
  @Input({ required: true }) NYGCupdate: Record<string, Record<number, string>>

  @Input({ required: true }) NYGisPlaying: string = "NYG"
  @Input({ required: true }) NYGplayersNickName: string[]
  @Input({ required: true }) NYGmyTeam: string
  @Input({ required: true }) NYGcanUsePower: WritableSignal<boolean>
  @Input({ required: true }) NYGmyroom: string
  @Input({ required: true }) NYGincludeR: WritableSignal<boolean>
  @Input({ required: true }) NYGincludeD: WritableSignal<boolean>
  @Input({ required: true }) NYGincludeT: WritableSignal<boolean>
  @Input({ required: true }) NYGincludeN: WritableSignal<boolean>
  @Input({ required: true }) NYGincludeF: WritableSignal<boolean>
  @Input({ required: true }) NYGincludeB: WritableSignal<boolean>
  @Input({ required: true }) NYGincludeC: WritableSignal<boolean>

  @Input({ required: true }) NYGcovert: WritableSignal<boolean>
  @Input({ required: true }) NYGbet: WritableSignal<boolean>
  @Input({ required: true }) NYGnexus: WritableSignal<boolean>
  @Input({ required: true }) NYGrewind: WritableSignal<boolean>
  @Input({ required: true }) NYGfreeze: WritableSignal<boolean>
  @Input({ required: true }) NYGtag: WritableSignal<boolean>
  @Input({ required: true }) NYGdraw: WritableSignal<boolean>

  @Input({ required: true }) NYGblock: WritableSignal<boolean>

}
