import { Component, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-view-pattern3',
  imports: [],
  templateUrl: './view-pattern3.component.html',
  styleUrls: ['./touchup/color.scss', './touchup/view-pattern3.scss'],
})
export class ViewPattern3Component {
  @Input() OCStitle: string = "GAMEs PLAYED"
  @Output() OCSEVENT = new EventEmitter<any>()
  TEST() {
    this.OCSEVENT.emit()
  }
}
