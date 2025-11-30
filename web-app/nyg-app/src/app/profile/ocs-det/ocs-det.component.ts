import { Component, input, Input } from '@angular/core';

@Component({
  selector: 'app-ocs-det',
  imports: [],
  templateUrl: './ocs-det.component.html',
  styleUrls: ['./touchup/ocs-det.scss', './touchup/color.scss'],

})
export class OcsDetComponent {
  @Input() OCStitle: string = "GAMEs PLAYED"
  @Input() OCSsubTitle: any = "99"
  @Input() OCSuseImage: boolean = false
  @Input() OCSImgSrc: string = ""
}
