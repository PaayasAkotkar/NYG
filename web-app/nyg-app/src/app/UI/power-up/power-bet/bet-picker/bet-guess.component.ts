import { Component, effect, input, Input, OnInit, signal, WritableSignal } from '@angular/core';
import { RoomService } from '../../../../roomService/room.service';
import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { HAND } from '../../../../helpers/hand';
import { RecBetItems } from '../../../../lobby/resources/lobbyResource';

@Component({
  selector: 'guess-bet',
  imports: [ReactiveFormsModule],
  templateUrl: './bet-guess.component.html',
  styleUrls: ['./touchup/animations.scss', './touchup/bet.scss', './touchup/color.scss'],
})
export class BetGuessComponent implements OnInit {
  @Input() NYGcollection: WritableSignal<RecBetItems> = signal({ firstCup: "", secondCup: "", thirdCup: "" })
  @Input() NYGmyTeam: string
  @Input() NYGmyRoom: string = ""
  @Input({ required: true }) NYGclash: WritableSignal<boolean>
  NYGref = input(false)
  jump: string[][] = []
  betGuessForm: FormGroup = this.FB.group({})

  constructor(private room: RoomService, private FB: FormBuilder) {
    effect(() => {
      var z = [this.NYGcollection().firstCup, this.NYGcollection().secondCup, this.NYGcollection().thirdCup]
      const __ref = HAND.ConvDtoTwoD(z)
      this.jump = HAND.sentencesSeparation(__ref, false)
    })
  }

  getBetGuessToken: string

  GetBetGuessToken(token: string) {
    this.getBetGuessToken = token
  }

  Guess() {
    if (!this.NYGref) {
      this.room.BetOn(this.NYGmyRoom, this.getBetGuessToken, this.NYGclash())
    }
  }

  ngOnInit(): void {
  }

}
