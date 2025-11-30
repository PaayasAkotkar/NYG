import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-ocs-header-v3',
  imports: [],
  templateUrl: './ocs-header-v3.component.html',
  styleUrls: ['./touchup/color.scss', './touchup/ocs-header.scss']
})
export class OcsHeaderV3Component {
  @Input({ required: true }) OCSheader: string = "NYG"
}
