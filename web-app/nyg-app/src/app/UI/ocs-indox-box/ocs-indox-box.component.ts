import { Component, EventEmitter, Input, Output } from '@angular/core';


@Component({
  selector: 'app-ocs-indox-box',
  imports: [],
  templateUrl: './ocs-indox-box.component.html',
  styleUrls: ['./touchup/box.scss', './touchup/color.scss',
    './touchup/animation.scss'
  ],
})
export class OcsIndoxBoxComponent {
  @Input() OCSlatest: boolean = false
  @Input() OCSmsg: string = ""
  @Output() OCSlisten: EventEmitter<void> = new EventEmitter<void>()
  @Output() OCSview: EventEmitter<boolean> = new EventEmitter<boolean>()
  listen() {
    this.OCSlisten.emit()
  }
}
