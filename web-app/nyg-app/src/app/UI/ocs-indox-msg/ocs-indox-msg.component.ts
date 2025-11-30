import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-ocs-indox-msg',
  imports: [],
  templateUrl: './ocs-indox-msg.component.html',
  styleUrls: ['./touchup/box.scss', './touchup/color.scss',
    './touchup/animation.scss'
  ],
})
export class OcsIndoxMsgComponent {
  @Input({ required: true }) OCStitle: string = "NYG"
  @Input({ required: true }) OCSsubtitle: string = "welcome to nyg hfh fghfhf h gfhhgfh fghhfghfgh fg"

}
