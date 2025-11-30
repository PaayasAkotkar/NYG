import { Component, EventEmitter, input, Input, OnDestroy, OnInit, Output, signal, ViewEncapsulation, WritableSignal } from '@angular/core';
import { LobbyResource } from '../../lobby/resources/lobbyResource';
import { RoomService } from '../../roomService/room.service';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { GuessText11xComponent } from "./guess-text-a/guess-text-11x.component";
import { GuessText25xComponent } from "./guess-text-b/guess-text-25x.component";
import { IMyView } from '../pick-pattern/pick-pattern.component';
import { ISpectate } from '../../roomService/resources/resource';
import { ClashGuessTextComponent } from "./clash-guess-text/clash-guess-text.component";

@Component({
  selector: 'app-guess-text',
  imports: [ReactiveFormsModule, GuessText11xComponent, GuessText25xComponent, ClashGuessTextComponent],
  templateUrl: './guess-text.component.html',
  styleUrls: ['./touchup/guess.scss', './touchup/color.scss'],
  encapsulation: ViewEncapsulation.Emulated
})
export class GuessTextComponent implements OnInit {

  constructor(private room: RoomService) { }

  ngOnInit(): void { }

  @Input({ required: true }) OCStheme: string = "text11x"
  @Input({ required: true }) NYGfriend: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGmyroom: string
  @Input({ required: true }) NYGset: number
  @Input({ required: true }) NYGround: number
  @Input({ required: true }) NYGmyTeam: string = "RED"
  @Input({ required: true }) NYGunFreeze: WritableSignal<boolean> = signal(true)
  @Input({ required: true }) NYGuseCovertPower: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGunderTast: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGroundOver: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGredScore: number = 1
  @Input({ required: true }) NYGblueScore: number = 1
  @Input({ required: true }) NYGclash: WritableSignal<boolean>
  @Input({ required: true }) NYGchances: WritableSignal<number>
  @Input({ required: false }) NYGPressedTime: number = 0
  @Input({ required: false }) NYGonFire: number = 0
  @Input({ required: false }) NYGfinalBoss: WritableSignal<boolean>
  @Input({ required: false }) NYGlastDance: WritableSignal<boolean>
  @Input({ required: true }) NYGtimeout: boolean

  // stops the timer
  @Output() NYGguessDone = new EventEmitter<boolean>()

  // guess token
  gameForm = new FormGroup(
    {
      GuessToken: new FormControl('')
    }
  )

  getGuessToken: string = String(this.gameForm.get('GuessToken')?.value)
  GuessToken = ""

  myView(text: string) {
    var token: ISpectate = { scrollHappend: false, word: text, onClick: "", scrollPos: -1 }
    this.room.TextUpdate(this.NYGmyroom, this.NYGclash(), this.NYGfinalBoss(), token)
  }

  Guess() {
    this.NYGroundOver.set(true)
    this.NYGguessDone.emit(true)
    this.GuessToken = String(this.gameForm.get('GuessToken')?.value)
    // if (!this.NYGclash()) {
    //   this.NYGround += 1
    //   this.room.GuessToken(this.GuessToken, this.NYGredScore,
    //     this.NYGblueScore, this.NYGmyroom, this.NYGround,
    //     this.NYGset, 0, false, -1, 0, false, false) // dont know about firend=false and to=false
    // } else {
    //   this.NYGround += 1 // update the round before sending it for new token
    //   this.room.GuessToken(this.GuessToken,
    //     this.NYGredScore, this.NYGblueScore, this.NYGmyroom,
    //     this.NYGround,
    //     this.NYGset, this.NYGchances(), true,
    //     this.NYGPressedTime, this.NYGonFire, this.NYGfinalBoss(), this.NYGlastDance())
    // }

    this.NYGround += 1 // update the round before sending it for new token
    this.room.GuessToken(this.GuessToken,
      this.NYGredScore, this.NYGblueScore, this.NYGmyroom,
      this.NYGround,
      this.NYGset, this.NYGchances(), this.NYGclash(),
      this.NYGPressedTime, this.NYGonFire, this.NYGfinalBoss(), this.NYGlastDance())

  }

}
