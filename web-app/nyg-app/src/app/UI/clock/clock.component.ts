import { ChangeDetectionStrategy, ChangeDetectorRef, Component, effect, Inject, Input, OnChanges, OnDestroy, OnInit, PLATFORM_ID, signal, SimpleChanges, ViewEncapsulation, WritableSignal } from '@angular/core';
import { TimerService } from '../../timer/timer.service';
import { BehaviorSubject, interval, Observable, scan, startWith, Subject, Subscription, switchMap, takeUntil, takeWhile, tap } from 'rxjs';
import { isPlatformBrowser } from '@angular/common';
import { ClockTimerService } from './clock-timer.service';

@Component({
  selector: 'app-clock',
  imports: [],
  templateUrl: './clock.component.html',
  styleUrls: ['./touchup/animations.scss', './touchup/clock.scss', './touchup/color.scss'],
  encapsulation: ViewEncapsulation.Emulated
})
export class ClockComponent implements OnDestroy, OnInit {
  @Input({ required: false }) NYGdiscussionSec: number = 10
  @Input({ required: false }) Start: WritableSignal<boolean> = signal(true)
  displaySec: string = ""
  displayMin: string = ""
  NYGdiscussionMin: number = 2
  Channel: Subscription = new Subscription()
  StopChannel: Subject<void> = new Subject<void>()

  ngOnDestroy(): void {
    this.UnSub()
    this.MarkComplete()
    this.Start.set(true)
    this._timer.StopTimer()
    this._timer.UnSub()
    this._timer.MarkDone()
  }
  MarkComplete() {
    this.StopChannel.next()
    this.StopChannel.complete()
  }

  UnSub() {
    if (!this.Channel.closed)
      this.Channel.unsubscribe()
  }
  ngOnInit(): void {
    if (this.Start()) {
      var min = this.NYGdiscussionMin
      var sec = this.NYGdiscussionSec
      var speed = 1000
      this.StopChannel = new Subject<void>()

      this._timer.Initialize(this.NYGdiscussionMin, this.NYGdiscussionSec, speed)

      this._timer.Start(min, sec, speed)
      this.Channel = this._timer.Watch().pipe(takeUntil(this.StopChannel)).subscribe({
        next: (token) => {
          this.NYGdiscussionSec = token.sec
          this.NYGdiscussionMin = token.min
          this.displaySec = this.NYGdiscussionSec.toString()
        }

      })
    }
  }

  constructor(private _timer: ClockTimerService) { }

}
