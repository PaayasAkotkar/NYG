import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-ocs-header-v2',
  imports: [],
  templateUrl: './ocs-header-v2.component.html',
  styleUrls: ['./touchup/ocs-headerv2.scss', './touchup/color.scss']
})
export class OcsHeaderV2Component {
  @Input() OCSheader: any = "HEADER"
}
