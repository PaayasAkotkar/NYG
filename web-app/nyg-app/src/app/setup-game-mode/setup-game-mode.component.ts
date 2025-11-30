import { Component, Input, ViewEncapsulation } from '@angular/core';
import { SportsService } from './services/setup-sports.service';
import { EntertainmentService } from './services/setup-entertainment.service';
@Component({
  selector: 'app-setup-game-mode',
  imports: [],
  templateUrl: './setup-game-mode.component.html',
  styles: `
  .setup{
    display:flex;
    width:100%;
    height:100%;
    justify-content:center;
    align-items:center;
  }
  .bl{
    width:100%;
    height:100%;
    display:block
  }
  `,
  encapsulation: ViewEncapsulation.Emulated
})
export class SetupGameModeComponent {
  @Input() view: string = ""
  viewSports() {
    this.sports.Setup()
  }
  viewEntertainment() {
    this._entertainement.Setup()
  }
  constructor(private sports: SportsService, private _entertainement: EntertainmentService) { }

}
