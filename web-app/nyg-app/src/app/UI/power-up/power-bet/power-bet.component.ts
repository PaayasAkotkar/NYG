import { Component, Input, OnInit, Output, EventEmitter, WritableSignal, signal, OnChanges, SimpleChanges } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { RoomService } from '../../../roomService/room.service';
import { OCSFieldFormat } from '../../../helpers/patterns';
import { MatCardModule } from '@angular/material/card';
import { PickPatternComponent } from "../../pick-pattern/pick-pattern.component";
@Component({
  selector: 'bet-on',
  imports: [MatCardModule, ReactiveFormsModule, PickPatternComponent],
  templateUrl: './power-bet.component.html',
  styleUrls: ['./touchup/bet.scss', './touchup/color.scss', './touchup/animations.scss'],
})

export class PowerBetComponent implements OnInit, OnChanges {

  @Input() NYGcollection: WritableSignal<OCSFieldFormat[]> = signal([
  ])

  @Input({}) NYGproceed: WritableSignal<boolean>
  @Input({ required: false }) NYGmyTeam: string
  @Input({ required: false }) NYGmyRoom: string = ""
  @Input({ required: false }) NYGMode: string = ""
  @Input({ required: false }) NYGclash: WritableSignal<boolean>
  @Input({ required: false }) NYGround: number
  @Input({ required: false }) NYGset: number
  @Output() NYGstop: EventEmitter<boolean> = new EventEmitter<boolean>()
  @Input({ required: false }) NYGtimeout: boolean
  NYGbetForm: FormGroup = this.FB.group({})
  jump: string[][] = []
  index: number[] = []
  @Input() NYGref: boolean = false
  constructor(private room: RoomService, private FB: FormBuilder) { }
  ngOnChanges(changes: SimpleChanges): void {
    if (changes['NYGtimeout']) {
      if (this.NYGtimeout == true) {
        this.Bot()
        this.NYGstop.emit(false)
      }
    }
  }
  ngOnInit(): void {
  }

  getBetToken: string
  Bot() {
    if (!this.NYGref) {
      var b = this.NYGcollection()[0].OCStemplate
      this.room.BetOnTimeOut(this.NYGmyRoom, b, this.NYGclash(), this.NYGround, this.NYGset)
    }
  }
  GetBetToken(token: string) {
    this.getBetToken = token
  }

  Bet() {
    this.room.BetOn(this.NYGmyRoom, this.getBetToken, this.NYGclash())
  }

}
