import { Component, Input, OnChanges, SimpleChanges, } from '@angular/core';
import { RemoveClassDirective } from "../power-up/power-up-pattern/remove-class.directive";

@Component({
  selector: 'app-home-page-label',
  imports: [RemoveClassDirective],
  templateUrl: './home-page-label.component.html',
  styleUrls: ['./touchup/animation.scss', './touchup/color.scss', './touchup/homepage-label.scss'],
})
export class HomePageLabelComponent {
  @Input() NYGTitle: string = "SPURS"
  @Input() NYGSubtitle: string = "500"
  @Input() NYGdoIt: boolean = true

}
