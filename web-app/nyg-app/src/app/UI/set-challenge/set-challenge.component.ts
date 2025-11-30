import { Component, EventEmitter, input, Input, OnChanges, OnDestroy, OnInit, Output, signal, SimpleChanges, ViewEncapsulation, WritableSignal } from '@angular/core';
import { RoomService } from '../../roomService/room.service';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { IMyView, PickPatternComponent } from '../pick-pattern/pick-pattern.component';
import { OCSFieldFormat } from '../../helpers/patterns';
import { ISpectate } from '../../roomService/resources/resource';
import { in_out } from '../create-room/animation/cr-animation';

@Component({
  selector: 'set-challenge',
  imports: [PickPatternComponent, ReactiveFormsModule, CommonModule],
  templateUrl: './set-challenge.component.html',
  encapsulation: ViewEncapsulation.Emulated
})
export class SetChallengeComponent implements OnDestroy, OnInit, OnChanges {
  @Input({ required: false }) NYGmyTeam: string
  @Input({ required: false }) NYGfriend: WritableSignal<boolean>
  @Input({ required: false }) NYGmyroom: string
  @Input({ required: false }) NYGproceed: WritableSignal<boolean>
  @Input({ required: true }) NYGClash: WritableSignal<boolean>
  @Input({ required: true }) NYGchances: WritableSignal<number>
  @Input({ required: true }) NYGround: number
  @Input({ required: true }) NYGMode: string = ''
  @Input({ required: true }) NYGfinalBoss: WritableSignal<boolean>
  @Input({ required: true }) NYGlastDance: WritableSignal<boolean>
  @Input({ required: true }) NYGtimeout: boolean = false
  @Input({ required: true }) NYGset: number = 0
  @Output() NYGstop: EventEmitter<boolean> = new EventEmitter<boolean>(false) // this basically sets to default value

  scrollHappend: boolean = false

  // challenge exchange token
  challengeForm = new FormGroup(
    {
      ChallengeToken: new FormControl('')
    }
  )

  myView(token: IMyView) {
    var _k: ISpectate = { scrollHappend: token.scrollHappend, scrollPos: token.position, onClick: token.selectedToken, word: "" }

    this.room.Spectate(this.NYGmyroom, this.NYGClash(), this.NYGfinalBoss(), _k)
  }

  ngOnChanges(changes: SimpleChanges): void {
    if (changes['NYGtimeout']) {
      const triggerBot = this.NYGtimeout == true
      if (triggerBot) {
        this.Bot()
        this.NYGstop.emit(false)
      }
    }
  }

  Bot() {
    var rand = this.NYGlist()[0].OCStemplate
    this.room.ChallengeTimeOut(this.NYGmyroom, this.NYGClash(), rand, this.NYGround, this.NYGset)
  }

  @Input({ required: false }) NYGlist: WritableSignal<OCSFieldFormat[]> = signal([
  ])

  ngOnInit(): void {
  }
  constructor(private room: RoomService) {

  }

  getChallengeToken: string = ""

  ngOnDestroy(): void { }

  GetToken(token: string) {
    this.getChallengeToken = token
    var _k: ISpectate = { scrollHappend: false, scrollPos: -1, onClick: token, word: "" }
    this.room.Spectate(this.NYGmyroom, this.NYGClash(), this.NYGfinalBoss(), _k)
  }

  Challenge() {
    // if (this.getChallengeToken != "" && !this.NYGClash()) {
    //   this.room.challengeToken(this.getChallengeToken.toLowerCase(), this.NYGmyroom, false, 0, false, false, this.NYGround)
    // } else {
    //   this.room.challengeToken(this.getChallengeToken.toLowerCase(), this.NYGmyroom, true, this.NYGchances(), this.NYGfinalBoss(), this.NYGlastDance(), this.NYGround)
    // }
    if (this.getChallengeToken != "") {
      this.room.challengeToken(this.getChallengeToken.toLowerCase(), this.NYGmyroom, this.NYGClash(), this.NYGchances(), this.NYGfinalBoss(), this.NYGlastDance(), this.NYGround)
    }
  }

}
