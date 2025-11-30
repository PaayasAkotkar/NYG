import { Component, Input, signal, WritableSignal } from '@angular/core';
import { Tags2Component } from "../tags2/tags2.component";
import { TagsComponent } from "../tags/tags.component";
import { ListComponent } from "../list/list.component";
import { BarComponent } from "../bar/bar.component";

import { OCSFieldFormat } from '../../helpers/patterns';
import { Pairing, SpectateInfo } from '../../lobby/resources/lobbyResource';
import { BehaviorSubject } from 'rxjs';
import { NYGlistService } from '../list/nyglist.service';
import { PowerTagService } from '../power-up/power-tag/power-tag.service';
import { FrameService } from '../frame/frame.service';
import { RoomService } from '../../roomService/room.service';
import { CheatService } from '../../cheatsheet/cheat.service';
import { PowerService } from '../power-up/power.service';
import { ColorPalletes as theme } from '../../helpers/hand';
import { CountComponent } from "../count/count.component";
import { LocifyClockComponent } from '../clocks/locify-clock/locify-clock.component';
import { MysteriumClockComponent } from '../clocks/mysterium-clock/mysterium-clock.component';
import { GymClockComponent } from '../clocks/gym-clock/gym-clock.component';
import { FastlaneClockComponent } from '../clocks/fastlane-clock/fastlane-clock.component';
import { PhysicsClockComponent } from '../clocks/physics-clock/physics-clock.component';
import { NodeClockComponent } from '../clocks/node-clock/node-clock.component';
import { FireComponent } from '../fire/fire.component';
import { SetDictionaryComponent } from "../set-dictionary/set-dictionary.component";
import { OnModeDirective } from "../../directives/on-mode.directive";

@Component({
  selector: 'app-ref-setup-board',
  imports: [LocifyClockComponent, MysteriumClockComponent, GymClockComponent, FastlaneClockComponent, PhysicsClockComponent, NodeClockComponent, FireComponent, Tags2Component, TagsComponent, ListComponent, BarComponent, CountComponent, SetDictionaryComponent, OnModeDirective],
  templateUrl: './ref-setup-board.component.html',
  styleUrls: ['./touchup/setup-board.scss', './touchup/color.scss'],

})
export class RefSetupBoardComponent {
  fromColor: string = ''
  toColor: string = ''
  bg: string = ''
  NYGspectate: boolean = false
  NYGlock: WritableSignal<boolean> = signal(true)
  NYGround: number = 1
  NYGonFire: number = 0
  NYGClashProfile: Record<string, Pairing> = {
    "P": { teamname: "RED", chances: 10 },
    "A": { teamname: "RED", chances: 10 },
    "Ya": { teamname: "BLUE", chances: 10 },
    "S": { teamname: "BLUE", chances: 10 },

  }
  CollectKeys: string[] = [
    // 'king', 'queen', 'prince', 'princess'
  ]

  watch: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false)

  constructor(private ls: NYGlistService, private tg: PowerTagService, private left: FrameService, private room: RoomService, private sco: CheatService, private po: PowerService) {
  }

  viewStats() { }

  viewLeftOver() {
  }
  viewList() {
  }
  viewPowers() {
  }

  nHeader = ['POWER', 'SCOREBOARD'].sort()
  ClashEntries: Record<string, Pairing> = {
  }

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
    this.update()
  }
  NYGno = signal(false)
  NYGyes = signal(true)

  NYGChances: WritableSignal<number> = signal(10)
  NYGuseTag: WritableSignal<boolean> = signal(false)
  NYGstartClock: WritableSignal<boolean> = signal(true)
  @Input() isNYGClash: WritableSignal<boolean> = signal(false)
  NYGrating: number = 67

  @Input() NYGlist: WritableSignal<OCSFieldFormat[]> = signal([
    { OCStemplate: "2001 I AM HERE", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2009 welcome to PUNE", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2010 Elementary School", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2011 got into CRICKET", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2012 went to Playground", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2013 too much into EA CRICKET 2007", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2014 too much into WWE impact pc game", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2015 new Mid School", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2016 too much into NBA 2k14 pc", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2017 trip to BANGLORE", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2018 welcome to NAGPUR", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2019 graducated secondary school", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2020 welcome back to PUNE", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2021 reunion with old FRIENDS", OCSformControl: false, OCSToolTip: "" },

  ])

  NYGLproceed: WritableSignal<boolean> = signal(false)
  @Input() NYGMode: WritableSignal<string> = signal('node')
  NYGroomMode: string = "2v2"
  NYGgameClock: number = 10
  NYGdecisionTime: number = 20

  Conns: WritableSignal<number> = signal(4)
  AS: number = 0
  HS: number = 0

  NYGisPlaying: string = "NYG"
  NYGmyTeam: string = "RED"
  NYGmyroom: string = "JOURNEY"

  NYGblock: WritableSignal<boolean> = signal(false)
  update() {
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
}
