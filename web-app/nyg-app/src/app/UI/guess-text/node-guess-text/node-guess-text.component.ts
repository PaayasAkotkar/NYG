import { Component, EventEmitter, Input, Output, signal, WritableSignal } from '@angular/core';
import { FormGroup, ReactiveFormsModule } from '@angular/forms';

@Component({
  selector: 'app-node-guess-text',
  imports: [ReactiveFormsModule],
  templateUrl: './node-guess-text.component.html',
  styleUrls: ['./touchup/guess.scss', './touchup/animation.scss', './touchup/color.scss']
})
export class NodeGuessTextComponent {

  @Input({ required: false }) OCSformGroup: FormGroup
  @Input({ required: false }) OCSformControl: string
  @Input({ required: false }) NYGfriend: WritableSignal<boolean> = signal(false)
  @Input({ required: false }) NYGmyroom: string
  @Input({ required: false }) NYGset: number
  @Input({ required: false }) NYGround: number
  @Input({ required: false }) NYGunFreeze: WritableSignal<boolean> = signal(false)
  @Input({ required: false }) NYGuseCovertPower: WritableSignal<boolean> = signal(true)
  @Input({ required: false }) NYGunderTast: WritableSignal<boolean> = signal(false)
  @Input({ required: false }) NYGmyTeam: string = "BLUE"

  @Output() NYGguess = new EventEmitter<boolean>()
  @Output() NYGtext = new EventEmitter<string>()

  change(token: string) {
    this.NYGtext.emit(token)
  }
  Guess() {
    this.NYGguess.emit(true)
  }
}
