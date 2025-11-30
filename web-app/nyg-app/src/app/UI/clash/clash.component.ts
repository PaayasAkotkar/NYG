import { Component, EventEmitter, Output } from '@angular/core';

@Component({
  selector: 'app-clash',
  imports: [],
  templateUrl: './clash.component.html',
  styleUrls: ['./touchup/clash.scss', './touchup/color.scss', './touchup/animations.scss'],

})
export class ClashComponent {
  @Output() OCSListen: EventEmitter<void> = new EventEmitter<void>()
  Click() {
    this.OCSListen.emit()
  }
}
