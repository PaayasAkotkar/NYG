import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-ocs-card',
  imports: [],
  templateUrl: './ocs-card.component.html',
  styleUrls: ['./touchup/card.scss', './touchup/color.scss',
    './touchup/animation.scss'
  ],
})
export class OcsCardComponent {
  @Input() NYGsetHeadline: string = ""
}
