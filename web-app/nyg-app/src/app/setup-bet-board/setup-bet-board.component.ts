import { Component, Input, OnDestroy, OnInit, WritableSignal } from '@angular/core';
import { BetPowerService } from '../UI/power-up/power-bet/power-bet.service';
import { RoomService } from '../roomService/room.service';
import { FormControl, FormGroup } from '@angular/forms';
import { CheatService } from '../cheatsheet/cheat.service';

@Component({
  selector: 'app-setup-bet-board',
  imports: [],
  templateUrl: './setup-bet-board.component.html',
})
export class SetupBetBoardComponent implements OnInit, OnDestroy {
  @Input({ required: false }) NYGmyRoom: string = "NYG"
  @Input({ required: false }) NYGcoll: any[]
  @Input({ required: false }) NYGisChallengeSet: WritableSignal<boolean>
  @Input({ required: false }) NYGmyTeam: string
  @Input({ required: false }) NYGclash: WritableSignal<boolean>
  betForm = new FormGroup({
    BetToken: new FormControl('')
  })

  ngOnInit(): void {
    this.bs.Setup(this.NYGmyRoom, this.NYGcoll, this.NYGmyTeam, this.NYGclash)
  }
  ngOnDestroy(): void {
    this.bs.closeSetup()
  }
  constructor(private sco: CheatService, private bs: BetPowerService, private room: RoomService) { }
}
