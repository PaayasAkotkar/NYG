import { ChangeDetectionStrategy, EventEmitter, Component, Input, Output, input } from '@angular/core';
import { OnProgressDirective } from './on-progress.directive';
import { PBAnimation } from './animation/animation';
import { DailyEventsService, Reward } from '../../setup-daily-events/daily-events.service';
import { SetupDailyEventsComponent } from '../../setup-daily-events/setup-daily-events.component';
export interface IDailyEvents {
  completed: number
  description: string
  eventName: string
  isClash: boolean
  progress: number
  reward: string
  toComplete: number
}
export interface XX {
  inx: number,
  reward: string
}
@Component({
  selector: 'app-daily-events',
  imports: [OnProgressDirective],
  animations: [PBAnimation],
  changeDetection: ChangeDetectionStrategy.OnPush,
  templateUrl: './daily-events.component.html',
  styleUrls: ['./touchup/animation.scss', './touchup/color.scss', './touchup/dailyevents.scss'],
})

export class DailyEventsComponent {
  @Input({ required: true }) NYGevents: IDailyEvents
  @Output() OCSlisten: EventEmitter<XX> = new EventEmitter<XX>()
  @Input({ required: true }) monitor: any = 0

  constructor() { }

  listen(token: string) {
    this.OCSlisten.emit({ reward: token, inx: this.monitor })
  }
}