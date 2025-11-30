import { Component, Output, EventEmitter, Input, WritableSignal } from '@angular/core';


@Component({
  selector: 'app-view-pattern',
  imports: [],
  templateUrl: './view-pattern.component.html',
  styleUrl: './view-pattern.component.scss'
})
export class ViewPatternComponent {
  PageNavigation: number = 1
  @Input({ required: true }) NYGredScore: number = 1
  @Input({ required: true }) NYGblueScore: number = 1
  @Input({ required: true }) OCSgameTime: number = 10
  @Input({ required: true }) NYGset: number = 1
  @Input({ required: true }) NYGround: number = 1
  @Input({ required: true }) NYGroundOver: WritableSignal<boolean>
  @Input({ required: true }) NYGrewindRestart: WritableSignal<boolean>
  @Input({ required: true }) NYGunFreeze: WritableSignal<boolean>
  @Input({ required: true }) NYGuseRewindPower: WritableSignal<boolean>

  Navigate(e: number) {
    this.PageNavigation = parseInt(String(e))
  }
  @Output() OCSEVENT = new EventEmitter<any>()
  TEST() {
    this.OCSEVENT.emit()
  }
}
