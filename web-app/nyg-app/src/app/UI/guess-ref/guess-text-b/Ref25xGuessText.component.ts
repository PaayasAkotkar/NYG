import { Component, Input, Output, EventEmitter, WritableSignal, signal } from '@angular/core';
import { LobbyResource } from '../../../lobby/resources/lobbyResource';
import { RoomService } from '../../../roomService/room.service';
import { FormGroup, ReactiveFormsModule } from '@angular/forms';
import { OnTeamColorDirective } from './on-team-color.directive';
import { OnFreezeDirective } from '../../power-up/power-freeze/on-freeze.directive';
import { PowerCovertComponent } from "../../power-up/power-covert/power-covert.component";

@Component({
  selector: 'RefguessText-25x',
  imports: [OnTeamColorDirective, PowerCovertComponent],
  templateUrl: './Ref25xGuessText.component.html',
  styleUrls: ['./touchup/guess.scss', './touchup/color.scss'
    , './touchup/animation.scss'
  ],
})
export class Ref25xGuessTextComponent {
  constructor(private room: RoomService) {
  }

  @Input({ required: false }) NYGmyTeam: string = "BLUE"
  @Input() NYGcovertRef: boolean = false
  dv: string = "fuck"
}
