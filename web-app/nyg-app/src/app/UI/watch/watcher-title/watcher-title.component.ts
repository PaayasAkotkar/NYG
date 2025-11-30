import { Component, Input, OnInit } from '@angular/core';
import { OnTeamColorDirective } from "../../game-tickets/on-team-color.directive";

@Component({
  selector: 'app-watcher-title',
  imports: [OnTeamColorDirective],
  templateUrl: './watcher-title.component.html',
  styleUrls: ['./touchup/title.scss', './touchup/color.scss']
})
export class WatcherTitleComponent implements OnInit {
  @Input({ required: true }) NYGheader: string = ""
  @Input({ required: true }) NYGisRedTeam: boolean
  team: string = ""
  ngOnInit(): void {
    if (this.NYGisRedTeam) {
      this.team = "RED"
    } else {
      this.team = "BLUE"
    }
  }
}
