import { Component, CSP_NONCE, EventEmitter, Input, Output, ViewEncapsulation } from '@angular/core';
import { ControlContainer, FormGroupDirective, ReactiveFormsModule } from '@angular/forms';
import { OCSFieldFormat } from '../../helpers/patterns';

@Component({
  selector: 'app-ocs-select',
  imports: [ReactiveFormsModule],
  templateUrl: './ocs-select.component.html',
  styleUrls: ['./touchup/color.scss', './touchup/ocs-select.scss',
    './touchup/animation.scss'
  ],
  viewProviders: [{ provide: ControlContainer, useExisting: FormGroupDirective }],
  encapsulation: ViewEncapsulation.Emulated,

})
export class OcsSelectComponent {
  @Output() OCSselectValue = new EventEmitter<string>()
  @Input() NYGformControlName: string = ""
  @Input() NYGlabel: string = "MODE"
  @Input() NYGisError: boolean = true
  @Input() NYGerrorMessage: string
  @Input() NYGsetToolTip: boolean = false
  @Input() NYGtoolTip: string = ""
  @Input() NYGheader: string = "TYPE"
  @Input() NYGoptions: OCSFieldFormat[] = []

  Click(token: string) {
    this.OCSselectValue.emit(token)
  }

}
