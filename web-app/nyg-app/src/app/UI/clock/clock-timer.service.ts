import { isPlatformBrowser, Time } from '@angular/common';
import { Inject, Injectable, OnDestroy, PLATFORM_ID } from '@angular/core';
import { BehaviorSubject, interval, Observable, ReplaySubject, scan, startWith, Subject, switchMap, takeUntil, takeWhile, tap, timer } from 'rxjs';

interface Timer {
  min: number
  sec: number
  speed: number
}

@Injectable({
  providedIn: 'root'
})
export class ClockTimerService {

  constructor(@Inject(PLATFORM_ID) private platformID: Object) { }

  private StartChannel: BehaviorSubject<Timer>
  private recieveToken: ReplaySubject<Timer>
  private StopChannel: Subject<void>
  private defaultMin: number = 0
  private defaultSec: number = 0
  private Stop: boolean = false

  MarkDone() {
    this.StopChannel.next()
    this.StopChannel.complete()
  }
  UnSub() {
    if (!this.StartChannel.closed)
      this.StartChannel.unsubscribe()
  }
  StopTimer() {
    this.Stop = true
  }
  Watch(): Observable<Timer> {
    return this.recieveToken.asObservable()
  }

  Initialize(min: number, sec: number, speed: number) {
    this.defaultMin = min
    this.defaultSec = sec
    this.StopChannel = new Subject<void>()
    this.StartChannel = new BehaviorSubject<Timer>({ min: this.defaultMin, sec: this.defaultSec, speed: speed })
    this.recieveToken = new ReplaySubject<Timer>()
    this.Stop = false
  }

  Start(min: number, sec: number, _speed: number) {

    if (!isPlatformBrowser(this.platformID) && !this.Stop) { return }
    var speed = _speed

    this.defaultMin = min
    this.defaultSec = sec

    this.StartChannel = new BehaviorSubject<Timer>({ min: min, sec: sec, speed: speed })

    this.StartChannel.pipe(
      switchMap(() => {
        return timer(0, _speed).pipe(
          startWith(sec),
          scan((acc: Timer, _: number) => {
            let { min, sec, speed } = acc;
            if (sec === 0 && min > 0) {
              min -= 1
              sec = 59
            } else if (min == 0 && sec == 0) {
              min = this.defaultMin
              sec = this.defaultSec
            } else {
              sec -= 1
            }
            speed = _speed
            return { min, sec, speed }

          }, { min, sec, speed }),
          takeWhile(count => count.min >= 0 && count.min <= 10 && count.sec >= 0),
          tap((token) => {

            this.recieveToken.next(token)
          }),
        )

      }),
      takeWhile(token => token.min >= 0 && token.sec >= 0),
      takeUntil(this.StopChannel),
    )
      .subscribe({
        next(value) {
        },
      })
  }
}