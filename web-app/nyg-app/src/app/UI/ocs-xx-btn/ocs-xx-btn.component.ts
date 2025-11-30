import { Component, Output, EventEmitter, Input } from '@angular/core';

@Component({
  selector: 'app-ocs-xx-btn',
  imports: [],
  templateUrl: './ocs-xx-btn.component.html',
  styleUrls: ['./touchup/ocs-btn.scss', './touchup/color.scss',
    './touchup/animation.scss'
  ],
})
export class OcsXxBtnComponent {
  @Output() OCSListen = new EventEmitter<void>()
  @Input() OCSHeader: string = ""
  listen() {
    this.OCSListen.emit()
  }
}
