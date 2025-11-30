import { Component, Output, EventEmitter } from '@angular/core';
import { OnEnter } from '../locify-lobby/animation/animation';

@Component({
  selector: 'app-locify',
  imports: [],
  templateUrl: './locify.component.html',
  styleUrls: ['./touchup/locify.scss', './touchup/color.scss', './touchup/animations.scss'],

})
export class LocifyComponent {
  @Output() OCSListen = new EventEmitter<void>()
  Click() {
    this.OCSListen.emit()
  }
}
