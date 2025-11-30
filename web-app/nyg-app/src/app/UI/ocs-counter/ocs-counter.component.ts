import { Component, Input, Output, EventEmitter, OnInit, Signal, signal } from '@angular/core';
import { OCSbuttonV1 } from "../upgrade-menu/upgrade-part1/ocs-button.component";


import { OcsHeaderV2Component } from "../ocs-header-v2/ocs-header-v2.component";

@Component({
  selector: 'app-ocs-counter',
  imports: [OCSbuttonV1, OcsHeaderV2Component],
  templateUrl: './ocs-counter.component.html',
  styleUrls: ['./touchup/ocs-counter.scss', './touchup/color.scss']
})
export class OcsCounterComponent implements OnInit {
  @Input() OCSheader: any = 1
  @Output() OCSAdd = new EventEmitter<any>()
  @Output() OCSSub = new EventEmitter<any>()

  Add() {
    this.OCSAdd.emit()
  }
  Sub() {
    this.OCSSub.emit()
  }

  ngOnInit(): void {
  }
}
