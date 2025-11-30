import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-fire',
  imports: [],
  templateUrl: './fire.component.html',
  styleUrls: ['./touchup/fire.scss', './touchup/color.scss', './touchup/animation.scss']
})
export class FireComponent {
  @Input() NYGheader: string | number = 3
}
