import { Component, ViewEncapsulation } from '@angular/core';
import { MenuService } from './view-menu.service';

@Component({
  selector: 'app-view-menu',
  imports: [],
  templateUrl: './view-menu.component.html',
  styleUrls: ['./touchup/view-menu.component.scss'],
  encapsulation: ViewEncapsulation.Emulated

})
export class ViewMenuComponent {
  show: boolean = false
  constructor(private service: MenuService) { }
  toggle() {
    this.show = !this.show
    if (this.show) {
      this.show = false
      this.service.Setup()
    } else {
      this.service.closeSetup()
    }
  }
}
