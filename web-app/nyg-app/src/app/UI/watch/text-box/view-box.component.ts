import { Component, Input, OnInit, signal, WritableSignal } from '@angular/core';

import { WatcherTitleComponent } from "../watcher-title/watcher-title.component";
import { BannerComponent } from "../banner/banner.component";
import { WatcherScoreboardComponent } from "../watcher-scoreboard/watcher-scoreboard.component";
import { ClockComponent } from "../../clock/clock.component";

@Component({
  selector: 'app-watcher-view-box',
  imports: [WatcherTitleComponent, BannerComponent, WatcherScoreboardComponent, ClockComponent],
  templateUrl: './view-box.component.html',
  styleUrls: ['./touchup/view-box.scss', './touchup/color.scss', './touchup/animation.scss']

})
export class WatcherViewBoxComponent implements OnInit {
  @Input({ required: false }) NYGteamRedText: WritableSignal<string> = signal('writing')
  @Input({ required: false }) NYGteamBlueText: WritableSignal<string> = signal('writing fsdfsfsfs fsd f sdfsdfsf sfs dfsd fsd fsdf sf sdfs fs fsdf fsfs fs f sf sf sfsf sf sf dsf sf sf sf sf dsf dsfdsf ds fsf sf dsf sfd')
  @Input({ required: false }) NYGteamRedName: WritableSignal<string> = signal("KINGKING")
  @Input({ required: false }) NYGteamBlueName: WritableSignal<string> = signal("QUEENQUE")
  @Input({ required: false }) NYGredTeamScore: string | number = 0
  @Input({ required: false }) NYGblueTeamScore: string | number = 0

  ngOnInit(): void { }

}
