import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-watcher-scoreboard',
  imports: [],
  templateUrl: './watcher-scoreboard.component.html',
  styleUrls: ['./touchup/scoreboard.scss', './touchup/color.scss']
})
export class WatcherScoreboardComponent {
  @Input({ required: true }) NYGredTeamScore: string | number = ""
  @Input({ required: true }) NYGblueTeamScore: string | number = ""

}
