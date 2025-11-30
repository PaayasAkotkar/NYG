import { ChangeDetectorRef, Component, Inject, Input, NgZone, OnChanges, OnInit, PLATFORM_ID, SimpleChanges } from '@angular/core';
import { DailyEventsComponent, IDailyEvents, XX } from "../UI/daily-events/daily-events.component";
import { BehaviorSubject, filter, ReplaySubject } from 'rxjs';
import { DailyEventsService, Reward } from './daily-events.service';
import { isPlatformBrowser } from '@angular/common';
import { NYGqService } from '../graphql/nygq.service';
import { SessionService } from '../storage/session/session.service';

@Component({
  selector: 'app-setup-daily-events',
  imports: [DailyEventsComponent],
  templateUrl: './setup-daily-events.component.html',
  styleUrl: './setup-daily-events.component.scss'
})
export class SetupDailyEventsComponent implements OnInit, OnChanges {
  constructor(private run: NgZone, private cdr: ChangeDetectorRef, private se: SessionService, private q: NYGqService, private d: DailyEventsService,) { }
  @Input({ required: true }) NYGclashEvents: IDailyEvents[] = []
  @Input({ required: true }) NYGlocifyEvents: IDailyEvents[] = []

  @Input({ required: true }) Reward: string = "default"
  monitorRewardClash: ReplaySubject<Reward> = new ReplaySubject<Reward>()
  monitorRewardLocify: ReplaySubject<Reward> = new ReplaySubject<Reward>()

  ngOnChanges(changes: SimpleChanges): void { }

  getClashUpdate(k: XX) {
    if (this.NYGclashEvents[k.inx].toComplete == this.NYGclashEvents[k.inx].completed) {
      var j = k.reward.includes('C')
      this.monitorRewardClash.next({ isCoin: j, reward: k.reward, index: k.inx })
    }
  }

  getLocifyUpdate(k: XX) {
    if (this.NYGlocifyEvents[k.inx].toComplete == this.NYGlocifyEvents[k.inx].completed) {
      var j = k.reward.includes('C')
      this.monitorRewardLocify.next({ isCoin: j, reward: k.reward, index: k.inx })
    }
  }
  ngOnInit(): void { }
}
