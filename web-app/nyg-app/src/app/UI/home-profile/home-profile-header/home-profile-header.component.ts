import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-home-profile-header',
  imports: [],
  templateUrl: './home-profile-header.component.html',
  styleUrls: ['./touchup/color.scss', './touchup/ocs-nme.scss', './touchup/animation.scss'],
})
export class HomeProfileHeaderComponent {
  @Input() OCSheader: string = ""
}
