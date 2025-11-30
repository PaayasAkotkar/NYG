import { Component, Input } from '@angular/core';


@Component({
  selector: 'app-arbitar',
  imports: [],
  templateUrl: './arbitar.component.html',
  styleUrls: ['./touchup/arbitar.scss', './touchup/color.scss', './touchup/animations.scss'],

})
export class ArbitarComponent {
  @Input() NYGtitle: string = "tag"
  @Input() NYGsubTitle: string = "ye dgfdg fdgdf gdg g dgd gdf gs"

}
