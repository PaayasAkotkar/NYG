import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-banner',
  imports: [],
  templateUrl: './banner.component.html',
  styleUrls: ['touchup/banner.scss', './touchup/color.scss']

})
export class BannerComponent {
  @Input({ required: true }) NYGredTeamPlayer: string = ""
  @Input({ required: true }) NYGblueTeamPlayer: string = ""

}
