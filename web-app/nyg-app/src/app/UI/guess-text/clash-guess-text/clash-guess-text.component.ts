import { Component, EventEmitter, Input, Output, signal, WritableSignal } from '@angular/core';
import { FormGroup, ReactiveFormsModule } from '@angular/forms';
import { onUnfreeze } from '../../../animations/freezeLeave';

@Component({
  selector: 'app-clash-guess-text',
  imports: [ReactiveFormsModule],
  animations: [onUnfreeze],
  templateUrl: './clash-guess-text.component.html',
  styleUrls: ['./touchup/animation.scss', './touchup/color.scss', './touchup/guess.scss']
})
export class ClashGuessTextComponent {

  @Input({ required: true }) OCSformGroup: FormGroup = new FormGroup({})
  @Input({ required: true }) OCSformControl: string = ""
  @Input({ required: true }) NYGfriend: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGmyroom: string = ""
  @Input({ required: true }) NYGset: number = 1
  @Input({ required: true }) NYGround: number = 1
  @Input({ required: true }) NYGunFreeze: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGuseCovertPower: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGunderTast: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGmyTeam: string = "BLUE"

  @Output() NYGguess = new EventEmitter<boolean>()
  @Output() NYGtext = new EventEmitter<string>()

  change(token: string) {
    this.NYGtext.emit(token)
  }
  Guess() {
    this.NYGguess.emit(true)
  }
}
