import { Component, Input, Output, EventEmitter, ViewEncapsulation } from '@angular/core';
import { FormBuilder } from '@angular/forms';
@Component({
  selector: 'ocs-button',
  imports: [],
  templateUrl: './ocs-button.component.html',
  styleUrls: ['./touchup/upgrade-part1.scss', './touchup/color.scss'],
  encapsulation: ViewEncapsulation.Emulated
})
export class OCSbuttonV1 {

  @Input({ required: true }) NYGheader: any = "X"
  @Output() NYGEVENT = new EventEmitter<any>()

  Click() {
    this.NYGEVENT.emit(this.NYGheader)
  }
  constructor(private FB: FormBuilder) { }
}
