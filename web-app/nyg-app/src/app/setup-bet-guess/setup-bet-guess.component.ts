import { Component, Input, OnDestroy, OnInit, WritableSignal } from '@angular/core';
import { RoomService } from '../roomService/room.service';
import { BetGuessPowerService } from '../UI/power-up/power-bet/bet-picker/power-bet-picker.service';
import { RecBetItems } from '../lobby/resources/lobbyResource';

@Component({
  selector: 'app-setup-bet-guess',
  imports: [],
  templateUrl: './setup-bet-guess.component.html',
})
export class SetupBetGuessComponent implements OnInit, OnDestroy {
  constructor(private bg: BetGuessPowerService) { }
  @Input({ required: true }) NYGmyRoom: string
  @Input({ required: true }) NYGbetItem: WritableSignal<RecBetItems>
  @Input({ required: true }) NYGClash: WritableSignal<boolean>

  ngOnInit(): void {
    this.bg.Setup(this.NYGmyRoom, this.NYGbetItem, this.NYGClash)
  }
  ngOnDestroy(): void {
    this.bg.closeSetup()
  }
}
