import { Component, Input, input } from '@angular/core';

@Component({
  selector: 'app-alert',
  imports: [],
  templateUrl: './alert.component.html',
  styleUrls: ['./touchup/alert.scss', './touchup/animations.scss', './touchup/color.scss']
})
export class AlertComponent {
  @Input() OCSmessage: string = "WELCOME"
  @Input() OCStitle: string = "NYG"
}
