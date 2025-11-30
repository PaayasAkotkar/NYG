import { Component, Input, Output, EventEmitter, ViewEncapsulation, OnInit, WritableSignal, signal, input } from '@angular/core';
import { ControlContainer, FormControl, FormGroupDirective, ReactiveFormsModule } from '@angular/forms';


@Component({
  selector: 'app-ocs-input',
  imports: [ReactiveFormsModule],
  templateUrl: './ocs-input.component.html',
  styleUrls: ['./touchup/ocs-input.scss', './touchup/color.scss',
    './touchup/animation.scss'
  ],
  viewProviders: [{ provide: ControlContainer, useExisting: FormGroupDirective }],
  encapsulation: ViewEncapsulation.Emulated
})
export class OcsInputComponent implements OnInit {
  @Input() NYGheader: string = "NICKNAME"
  @Input() NYGformControl!: string
  @Input() NYGformControl_!: FormControl
  @Input() NYGType: string = ''
  @Input() NYGsetToolTip: boolean = false
  @Input() NYGtoolTip: string = ""
  @Input() NYGerrorMessage: string
  @Input() NYGisError: boolean = true
  @Output() NYGEVENT = new EventEmitter<any>()
  test() {
    this.NYGEVENT.emit()
  }

  ngOnInit(): void {
    var a = /[^a-zA-Z\s-']/g
  }
}
