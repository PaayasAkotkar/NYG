import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-ocs-nme',
  imports: [],
  templateUrl: './ocs-nme.component.html',
  styleUrls: ['./touchup/ocs-nme.scss', './touchup/color.scss'],

})
export class OcsNmeComponent {
  // character limit is only 6
  @Input() OCSname: string = "NYG"
}
