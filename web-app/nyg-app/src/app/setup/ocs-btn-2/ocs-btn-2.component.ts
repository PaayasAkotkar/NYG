import { Component, Input, Output, EventEmitter, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'ocs-btnV2',
  imports: [],
  templateUrl: './ocs-btn-2.component.html',
  styleUrls: ['./touchup/animation.scss', './touchup/btn.scss', './touchup/color.scss',],
  encapsulation: ViewEncapsulation.Emulated,
})
export class OcsBtn2Component {
  @Input() NYGheader: string = ""
  @Input() NYGpar: string = ""
  @Output() NYGEVENT = new EventEmitter<string>()
  Click() {
    this.NYGEVENT.emit(this.NYGpar)
  }
}
