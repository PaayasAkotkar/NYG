import { Component, Input, Output, EventEmitter } from '@angular/core';
import { RouterModule } from '@angular/router';
import { RemoveClassDirective } from "../power-up/power-up-pattern/remove-class.directive";

@Component({
  selector: 'app-home-page-nav',
  imports: [RouterModule, RemoveClassDirective],
  templateUrl: './home-page-nav.component.html',
  styleUrls: ['./touchup/animation.scss', './touchup/color.scss', './touchup/homepage-nav.scss'],

})
export class HomePageNavComponent {
  @Input() NYGHeader: string = ""
  @Input() NYGLink: string = ""
  @Input() OCSmsg: string = "new"
  @Input() OCSlatest: boolean = true
  @Input() NYGdoIt: boolean = true
  @Output() NYGListen = new EventEmitter<void>()
  Listen() {
    this.NYGListen.emit()
  }
}
