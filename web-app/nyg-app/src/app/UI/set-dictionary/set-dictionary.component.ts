import { Component, Input, EventEmitter, OnChanges, OnInit, Output, signal, SimpleChanges, WritableSignal } from '@angular/core';
import { RoomService } from '../../roomService/room.service';
import { IMyView, PickPatternComponent } from "../pick-pattern/pick-pattern.component";
import { FormControl, FormGroup } from '@angular/forms';
import { OCSFieldFormat } from '../../helpers/patterns';
import { ISpectate } from '../../roomService/resources/resource';

@Component({
  selector: 'set-dictionary',
  imports: [PickPatternComponent],
  templateUrl: './set-dictionary.component.html',
})
export class SetDictionaryComponent implements OnInit, OnChanges {
  @Input({ required: false }) OCScollection: WritableSignal<OCSFieldFormat[]> = signal([
  ])
  @Input() NYGMode: string = ''
  @Input({ required: true }) NYGClash: WritableSignal<boolean>
  @Input({ required: true }) NYGchances: WritableSignal<number>
  @Input({ required: false }) NYGisRef: boolean = false
  @Input({ required: false }) NYGmyTeam: string
  @Input({ required: false }) NYGfriend: WritableSignal<boolean>
  @Input({ required: true }) NYGmyroom: string
  @Input({ required: false }) NYGproceed: WritableSignal<boolean>
  @Input({ required: true }) NYGfinalBoss: WritableSignal<boolean>
  @Input({ required: true }) NYGlastDance: WritableSignal<boolean>
  @Input({ required: true }) NYGtimeout: boolean = false
  @Input({ required: true }) NYGround: number = 0
  @Input({ required: true }) NYGset: number = 0
  @Output() NYGstop: EventEmitter<boolean> = new EventEmitter<boolean>() // this basically sets to default value
  getSetDictionaryToken: string

  challengeDictionaryForm = new FormGroup(
    {
      ChallengeDictionaryToken: new FormControl(' ')
    }
  )

  constructor(private room: RoomService) { }

  ngOnInit(): void { }

  ngOnChanges(changes: SimpleChanges): void {
    if (changes['NYGtimeout']) {
      if (!this.NYGisRef) {

        const triggerBot = this.NYGtimeout == true
        if (triggerBot) {
          this.Bot()
          this.NYGstop.emit(false)
        }
      }
    }
  }
  Bot() {
    var rand = this.OCScollection()[0].OCStemplate
    this.room.DictionaryTimeOut(this.NYGmyroom, this.NYGClash(), rand, this.NYGround, this.NYGset)

  }
  TrackPosition: number = 0

  myView(token: IMyView) {
    this.TrackPosition = token.position
    var _k: ISpectate = { scrollHappend: token.scrollHappend, scrollPos: token.position, onClick: token.selectedToken, word: "" }
    this.room.Spectate(this.NYGmyroom, this.NYGClash(), this.NYGfinalBoss(), _k)
  }

  GetToken(token: string) {
    this.getSetDictionaryToken = token
    var _k: ISpectate = { scrollHappend: false, scrollPos: -1, onClick: token, word: "" }

    this.room.Spectate(this.NYGmyroom, this.NYGClash(), this.NYGfinalBoss(), _k)

  }

  Dictionary() {
    // if (!this.NYGClash()) {
    //   this.room.SetDictionary(this.getSetDictionaryToken, this.NYGmyroom, false, 0, false, false)
    // } else {
    //   this.room.SetDictionary(this.getSetDictionaryToken, this.NYGmyroom, true, this.NYGchances(), this.NYGfinalBoss(), this.NYGlastDance())
    // }
    if (this.getSetDictionaryToken != "" && !this.NYGisRef) {
      this.room.SetDictionary(this.getSetDictionaryToken, this.NYGmyroom, this.NYGClash(), this.NYGchances(), this.NYGfinalBoss(), this.NYGlastDance())
    }
  }

}
