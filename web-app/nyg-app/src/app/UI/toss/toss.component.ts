import { ChangeDetectorRef, EventEmitter, Component, Input, OnChanges, OnInit, Output, signal, SimpleChanges, ViewEncapsulation, WritableSignal } from '@angular/core';
import { LobbyResource } from '../../lobby/resources/lobbyResource';
import { RoomService } from '../../roomService/room.service';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';

import { HAND } from '../../helpers/hand';
import { OCSFieldFormat } from '../../helpers/patterns';
import { IMyView, PickPatternComponent } from "../pick-pattern/pick-pattern.component";
import { ISpectate } from '../../roomService/resources/resource';

@Component({
  selector: 'app-toss',
  imports: [ReactiveFormsModule, PickPatternComponent],
  templateUrl: './toss.component.html',
  encapsulation: ViewEncapsulation.Emulated
})
export class TossComponent implements OnInit, OnChanges {

  constructor(private cdr: ChangeDetectorRef, private room: RoomService, private _FB: FormBuilder) { }

  flip: WritableSignal<OCSFieldFormat[]> = signal(
    [
      { OCStemplate: 'TAIL', OCSformControl: false, OCSToolTip: "" },
      { OCStemplate: 'HEAD', OCSformControl: false, OCSToolTip: "" },
    ])

  @Input({ required: true }) NYGproceed: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGmyTeam: string = ""
  @Input({ required: true }) NYGClash: boolean = false
  @Input({ required: true }) NYGchances: number = 0
  @Input({ required: true }) NYGMode: string = ''
  @Input({ required: true }) NYGfinalBoss: boolean = false
  @Input({ required: true }) NYGlastDance: boolean = false
  @Input() isFriend: WritableSignal<boolean> = signal(false)
  @Input() NYGroomname: string = ""
  @Input({ required: true }) NYGtimeout: boolean = false
  @Input({ required: true }) NYGround: number = 0
  @Input({ required: true }) NYGset: number = 0
  @Output() NYGstop: EventEmitter<boolean> = new EventEmitter<boolean>() // this basically sets to default value

  /**
   * @todo make sure to set the default value non required control
   */
  tossForm = new FormGroup(
    {
      TossToken: new FormControl('')
    }
  )

  getTossToken = String(this.tossForm.get('tossToken')?.value)

  myView(token: IMyView) {
    var _k: ISpectate = { scrollHappend: token.scrollHappend, scrollPos: token.position, onClick: token.selectedToken, word: "" }
    this.room.Spectate(this.NYGroomname, this.NYGClash, this.NYGfinalBoss, _k)
  }

  ngOnChanges(changes: SimpleChanges): void {
    if (changes['NYGtimeout']) {
      const triggerBot = this.NYGtimeout

      console.log('triggered', changes['NYGtimeout'].currentValue)
      if (triggerBot === true) {
        console.log('bot send: ', this.NYGtimeout)
        this.Bot()
        this.NYGstop.emit(false)
      }
    }
    this.cdr.detectChanges()

  }

  Bot() {
    var b = this.flip()[0].OCStemplate
    this.room.TossTimeOut(this.NYGroomname, this.NYGClash, b, this.NYGround, this.NYGset)
  }
  ngOnInit(): void {
    HAND.FillControls(this.tossForm, this.flip(), this._FB)
  }

  OCScoinSide: string = ""

  TossBody(side: string) {
    this.OCScoinSide = side
    var _k: ISpectate = { scrollHappend: false, scrollPos: -1, onClick: side, word: "" }

    this.room.Spectate(this.NYGroomname, this.NYGClash, this.NYGfinalBoss, _k)
  }

  Toss() {

    // if (!this.isFriend && !this.NYGClash) {
    //   this.room.TossSession(this.OCScoinSide, false, this.NYGroomname, false, 0, false, false)
    // } else if (this.NYGClash) {
    //   this.room.TossSession(this.OCScoinSide, false, this.NYGroomname, this.NYGClash, this.NYGchances, this.NYGfinalBoss, this.NYGlastDance)
    // }
    // else {
    //   this.room.TossSession(this.OCScoinSide, true, this.NYGroomname, false, 0, false, false)
    // }
    if (this.OCScoinSide != "") {
      this.room.TossSession(this.OCScoinSide, this.isFriend(), this.NYGroomname, this.NYGClash, this.NYGchances, this.NYGfinalBoss, this.NYGlastDance)
    }
  }

}