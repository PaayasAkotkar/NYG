import { Component, Input, WritableSignal } from '@angular/core';
import { BannerComponent } from "./banner/banner.component";
import { ScoresheetComponent } from "./scoresheet/scoresheet.component";
import { WatcherTimerComponent } from "./timer/timer.component";
import { WatcherScoreboardComponent } from "./watcher-scoreboard/watcher-scoreboard.component";
import { WatcherTitleComponent } from "./watcher-title/watcher-title.component";
import { WatcherViewBoxComponent } from './text-box/view-box.component';

@Component({
  selector: 'app-watch',
  imports: [WatcherViewBoxComponent,],
  templateUrl: './watch.component.html',
  styleUrls: ['./touchup/watch.scss', './touchup/color.scss']
})
export class WatchComponent {
  @Input({ required: true }) NYGteamRedText: WritableSignal<string>
  @Input({ required: true }) NYGteamBlueText: WritableSignal<string>
  @Input({ required: true }) NYGteamRedName: WritableSignal<string>
  @Input({ required: true }) NYGteamBlueName: WritableSignal<string>
  @Input({ required: true }) NYGredTeamScore: string | number
  @Input({ required: true }) NYGblueTeamScore: string | number

}
