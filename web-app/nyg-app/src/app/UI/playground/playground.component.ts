import { Component, Input, EventEmitter, OnDestroy, OnInit, Output, ViewEncapsulation, WritableSignal, signal, AfterViewInit, ViewChild, ElementRef, Renderer2 } from '@angular/core';
import { GuessTextComponent } from "../guess-text/guess-text.component";
import { ScoreboardComponent } from "../scoreboard/scoreboard.component";
import { ColorPalletes as theme } from '../../helpers/hand';

import { OnStyleDirective } from "../../directives/on-style.directive";
@Component({
  selector: 'app-playground',
  imports: [GuessTextComponent, ScoreboardComponent, OnStyleDirective],
  templateUrl: './playground.component.html',
  styleUrls: ['./touchup/playground.scss', './touchup/color.scss'],
  encapsulation: ViewEncapsulation.Emulated
})
export class PlaygroundComponent implements OnInit, OnDestroy, AfterViewInit {
  @Input({ required: true }) NYGblueScore: number
  @Input({ required: true }) NYGredScore: number
  @Input({ required: true }) NYGfriend: WritableSignal<boolean>
  @Input({ required: true }) NYGmyroom: string
  @Input({ required: true }) NYGset: number
  @Input({ required: true }) NYGround: number
  @Input({ required: true }) NYGpowerActivated: WritableSignal<boolean>
  @Input({ required: true }) NYGrewindRestart: WritableSignal<boolean>
  @Input({ required: true }) OCSgameTime: number = 10
  @Input({ required: true }) NYGuseRewindPower: WritableSignal<boolean>
  @Input({ required: true }) NYGuseCovertPower: WritableSignal<boolean>
  @Input({ required: true }) NYGunFreeze: WritableSignal<boolean> = signal(true)
  @Input({ required: true }) NYGunderTast: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGroundOver: WritableSignal<boolean>
  @Input({ required: true }) NYGClash: WritableSignal<boolean>
  @Input({ required: true }) NYGchances: WritableSignal<number>
  @Input({ required: true }) NYGTime: number
  @Input({ required: true }) NYGmyTeam: string = "BLUE"
  @Input({ required: true }) NYGonFire: number = 0
  @Input({ required: true }) NYGfinalBoss: WritableSignal<boolean>
  @Input({ required: true }) NYGlastDance: WritableSignal<boolean>
  @Input({ required: true }) NYGguessTheme: string = ""
  @Input({ required: true }) NYGboardTheme: string = ""
  @Input({ required: true }) NYGtimeout: boolean
  @Input({ required: true }) OCSstartClock: WritableSignal<boolean>
  @Output() NYGtimerStopper = new EventEmitter<boolean>()
  isGuessDone: WritableSignal<boolean> = signal(false)
  bg: string = ""
  GuessDone(done: boolean) {
    this.isGuessDone.set(done)
    this.NYGtimerStopper.emit(done)
  }
  linearbg: boolean = false
  fromColor: string = ''
  toColor: string = ''
  ngOnDestroy(): void {
    this.NYGtimerStopper.emit(false)
  }
  ngAfterViewInit(): void { }
  ngOnInit(): void {
    switch (this.NYGboardTheme) {
      case ('text11x'):
        this.linearbg = true
        this.fromColor = theme.$nyg_black_palletes.umber_black
        this.toColor = theme.$nyg_black_palletes.potrait3_black


        break
      case ('text25x'):
        this.linearbg = true
        this.fromColor = theme.$nyg_brown_palletes.ad8_brown
        this.toColor = theme.$nyg_blue_palletes.ad4_blue

        break
      case 'clash-mysterium':
        this.bg = theme.$nyg_brown_palletes.ad8_brown
        break
      case 'clash-node':
        this.bg = theme.$nyg_blue_palletes.ad16_blue
        break
      case 'clash-fastlane':
        this.bg = theme.$nyg_green_palletes.ad19_green
        break
      case 'clash-pshycic':
        this.bg = theme.$nyg_voilet_palletes.ad9_voilet
        break
      case 'clash-gym':
        this.bg = theme.$nyg_black_palletes.ad8_black
        break
      case 'locify':
        this.bg = theme.$nyg_orange_palletes.ad9_orange
        break
      default:

    }
  }
  constructor(private renderer: Renderer2) { }
}