import { Component, Input, Output, EventEmitter } from '@angular/core';
import { ControlContainer, FormBuilder, FormControl, FormGroup, FormGroupDirective, ReactiveFormsModule } from '@angular/forms';
import { OcsBtnsComponent } from "../../UI/ocs-btns/ocs-btns.component";
import { OcsInputServiceV2 } from './ocs-input-service-v2.service';

@Component({
  selector: 'app-ocs-input-2',
  imports: [ReactiveFormsModule, OcsBtnsComponent],
  templateUrl: './ocs-input-2.component.html',
  styleUrls: ['./touchup/ocs-input.scss', './touchup/color.scss'],
  viewProviders: [{ provide: ControlContainer, useExisting: FormGroupDirective }],
})
export class OcsInput2Component {
  @Input() NYGformGroup: FormGroup = new FormBuilder().group({
    Token: ['']
  })
  @Input() NYGheader: string = "NICKNAME"
  @Output() NYGEVENT = new EventEmitter<string>()
  token: string = ""
  Token() {
    this.token = this.NYGformGroup.get('Token')?.value
    this.NYGEVENT.emit(this.token)
    this.i.closeSetup()
  }
  constructor(private i: OcsInputServiceV2) { }
}
