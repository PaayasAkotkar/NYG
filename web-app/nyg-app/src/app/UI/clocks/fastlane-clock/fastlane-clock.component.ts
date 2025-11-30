import { Component, effect, EventEmitter, Input, Output, signal, WritableSignal } from '@angular/core';
import { ClockTimerService } from '../../clock/clock-timer.service';
import { Subject, Subscription, takeUntil } from 'rxjs';
import { OCSFieldFormat } from '../../../helpers/patterns';
import { RoomService } from '../../../roomService/room.service';

@Component({
  selector: 'app-fastlane-clock',
  imports: [],
  templateUrl: './fastlane-clock.component.html',
  styleUrls: ['./touchup/clock.scss', './touchup/color.scss', './touchup/animation.scss']
})
export class FastlaneClockComponent {
  @Input({ required: true }) NYGdiscussionSec: number = 12
  @Input({ required: true }) NYGStart: WritableSignal<boolean> = signal(true)
  @Output() NYGtimeout: EventEmitter<boolean> = new EventEmitter<boolean>(false)
  NYGdiscussionMin: number = 2

  Channel: Subscription = new Subscription()
  StopChannel: Subject<void> = new Subject<void>()
  DisplayMin: string = ""
  DisplaySec: string = ""

  ngOnDestroy(): void {
    this.UnSub()
    this.MarkComplete()
    this.NYGStart.set(true)
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
  ngOnInit(): void { }

  constructor(private room: RoomService, private _timer: ClockTimerService) {
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
