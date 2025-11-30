import { Component, effect, EventEmitter, ElementRef, Input, NgZone, Renderer2, signal, ViewChild, WritableSignal, Output } from '@angular/core';
import { ClockTimerService } from '../../clock/clock-timer.service';
import { concatMap, delay, repeat, Subject, Subscription, takeUntil, tap, timer } from 'rxjs';
import { RemoveClassDirective } from "../../power-up/power-up-pattern/remove-class.directive";
import { OCSFieldFormat } from '../../../helpers/patterns';

@Component({
  selector: 'app-node-clock',
  imports: [RemoveClassDirective],
  templateUrl: './node-clock.component.html',
  styleUrls: ['./touchup/clock.scss', './touchup/color.scss', './touchup/animation.scss']
})
export class NodeClockComponent {
  @Input({ required: true }) NYGdiscussionSec: number = 12
  @Input({ required: true }) NYGStart: WritableSignal<boolean> = signal(false)
  @Output() NYGtimeout: EventEmitter<boolean> = new EventEmitter<boolean>(false)

  Slot: boolean = false
  NYGdiscussionMin: number = 2
  Channel: Subscription = new Subscription()
  StopChannel: Subject<void> = new Subject<void>()
  DisplayMin: string = ""
  DisplaySec: string = ""
  replaceWith: string = "animateLeft"
  replace: string = "animateRight"
  delay: boolean = false

  ngOnDestroy(): void {
    this.UnSub()
    this.MarkComplete()
    this.NYGStart.set(false)
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
    var swap = () => {
      var temp = this.replace
      this.replace = this.replaceWith
      this.replaceWith = temp
    }
    // slot=false to right
    // division into 4:
    // animation-duration=2.45s
    // 2450/4= 612.5
    // now to summarize the proportion for 500:
    // than 1500->2000...
    // sum= 1500+2000+2500+3000=9000
    // proportion=2450/9000=appr0.2722
    // than proportion*delay=1500*0.27
    this.Zone.runOutsideAngular(() => {
      const k = 0.2722
      var f = 1500, s = 2000, t = 2500, f = 3000, fth = 3500
      timer(0).pipe(
        // actual delay
        concatMap(() => timer(f * k).pipe(tap(() => {
          this.delay = true
          swap()
        }))),

        concatMap(() => timer(s * k).pipe(tap(() => {
          swap()
        }))),

        concatMap(() => timer(s * k * 2).pipe(tap(() => {
          swap()
          this.delay = true
        }))),

        repeat({
          delay: () => {
            swap()
            this.Slot = !this.Slot
            return timer(f * k * 2)
          }
        }),
      ).subscribe()

    })

  }

  constructor(private Zone: NgZone, private rend: Renderer2, private _timer: ClockTimerService) {

    effect(() => {
      if (this.NYGStart()) {
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
            if (token.sec < 10) {
              this.DisplaySec = '0' + this.NYGdiscussionSec.toString()
            } else {
              this.DisplaySec = this.NYGdiscussionSec.toString()
            }
            if (token.sec == 0 && token.min == 0) {
              console.log(' time up 0')
              this.NYGtimeout.emit(true)
            }
            this.DisplayMin = this.NYGdiscussionMin.toString()
          }
        })
      }

    })
  }



}
