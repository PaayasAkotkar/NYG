import { Component } from '@angular/core';
import { RoomsListComponent } from "../rooms-list/rooms-list.component";
import { OnEnter } from './animation/animation';
import { Router } from '@angular/router';

@Component({
  selector: 'app-locify-lobby',
  imports: [RoomsListComponent],
  animations: [OnEnter],
  templateUrl: './locify-lobby.component.html',
  styleUrls: ['./touchup/locify.scss', './touchup/color.scss', './touchup/animations.scss'],

})
export class LocifyLobbyComponent {
  show: boolean = false
  click() {
    this.show = !this.show
  }
  toMake() {
    this._r.navigateByUrl('/Search')
  }
  constructor(private _r: Router) { }
}
