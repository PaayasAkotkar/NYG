import { Component, Input, Output, EventEmitter, WritableSignal, signal, OnInit } from '@angular/core';
import { LobbyResource } from '../../../lobby/resources/lobbyResource';
import { RoomService } from '../../../roomService/room.service';
import { FormGroup, ReactiveFormsModule } from '@angular/forms';
import { OnTeamColorDirective } from './on-team-color.directive';
import { OnFreezeDirective } from "../../power-up/power-freeze/on-freeze.directive";


@Component({
  selector: 'guessText-25x',
  imports: [ReactiveFormsModule, OnTeamColorDirective, OnFreezeDirective],
  templateUrl: './guess-text-25x.component.html',
  styleUrls: ['./touchup/guess.scss', './touchup/color.scss'
    , './touchup/animation.scss'
  ],
})
export class GuessText25xComponent implements OnInit {

  @Input({ required: true }) OCSformGroup: FormGroup
  @Input({ required: true }) OCSformControl: string
  @Input({ required: true }) NYGfriend: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGmyroom: string
  @Input({ required: true }) NYGset: number
  @Input({ required: true }) NYGround: number
  @Input({ required: true }) NYGunFreeze: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGuseCovertPower: WritableSignal<boolean> = signal(true)
  @Input({ required: true }) NYGunderTast: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGmyTeam: string = "BLUE"

  @Output() NYGguess = new EventEmitter<boolean>()
  @Output() NYGtext = new EventEmitter<string>()

  ngOnInit(): void { }

  change(token: string) { this.NYGtext.emit(token) }

  Guess() { this.NYGguess.emit(true) }
}
