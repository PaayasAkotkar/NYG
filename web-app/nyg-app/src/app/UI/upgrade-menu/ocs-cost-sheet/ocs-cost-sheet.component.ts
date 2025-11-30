import { Component, Input, OnChanges, EventEmitter, OnInit, SimpleChanges, Output } from '@angular/core';
import { OCSbuttonV1, } from "../upgrade-part1/ocs-button.component";

@Component({
  selector: 'app-ocs-cost-sheet',
  imports: [OCSbuttonV1],
  templateUrl: './ocs-cost-sheet.component.html',
  styleUrls: ['./touchup/cost-sheet.scss', './touchup/color.scss'],
})
export class OcsCostSheetComponent implements OnChanges {
  @Input({ required: true }) OCScostSheet = '100S'
  //@Input() OCScost: number = 100
  @Input({ required: true }) OCSheader = 'COST'
  @Input({ required: true }) OCSbuttonHeader = 'U'

  @Output() OCSListen = new EventEmitter<void>()

  Listen() {
    this.OCSListen.emit()
  }
  ngOnChanges(changes: SimpleChanges): void { }
}
