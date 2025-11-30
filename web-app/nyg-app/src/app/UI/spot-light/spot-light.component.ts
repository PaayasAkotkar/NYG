import { Component, Input, NgZone, OnDestroy, OnInit, signal, WritableSignal } from '@angular/core';
import { TimerService } from '../../timer/timer.service';
import { BehaviorSubject, interval, scan, startWith, Subject, Subscription, switchMap, takeUntil, takeWhile, tap, timer } from 'rxjs';
import { SpotLightService } from './spot-light.service';

@Component({
  selector: 'app-spot-light',
  imports: [],
  templateUrl: './spot-light.component.html',
  styleUrls: ['./touchup/spot-light.scss',
    './touchup/color.scss', './touchup/animations.scss'
  ]
})
export class SpotLightComponent implements OnInit, OnDestroy {
  OCSTime: number = 5
  channel = new Subscription()
  stop = new Subject<void>()
  StartsIn: number = this.OCSTime
  constructor(private s: SpotLightService, private zone: NgZone, private timer: TimerService) { }
  ngOnInit(): void {
    this.zone.runOutsideAngular(() => {
      this.zone.run(() => {
        timer(1).pipe(takeUntil(this.stop)).subscribe(() => {
          this.StartFast(this.OCSTime, this.OCSTime, 300)
        })
      })
    })

    this.channel = this.Start().pipe(takeUntil(this.stop)).subscribe({
      next: (token) => {
        this.zone.run(() => {
          this.StartsIn = token
        })
      }
    })
  }
  StopFastChannel = new Subject<void>()
  recieveToken = new BehaviorSubject<number>(1)
  StartFastChannel = new BehaviorSubject<number>(0)
  StartFast(min: number, max: number, speed: number) {

    this.StartFastChannel.pipe(
      switchMap(() => {
        return interval(speed).pipe(
          startWith(min),
          // [5,10]=> acc==10 acc+1 if and only if it has reached the final time 
          scan(acc => acc -= 1, min),
          takeWhile(count => count >= 0 && count <= max),
          tap((token) => {
            min = token
            this.recieveToken.next(token)

          }),
        )
      }),
      takeWhile(token => token > 0),
      takeUntil(this.StopFastChannel),
    )
      .subscribe({
        next(value) {

        },
      })

  }
  Start() {
    return this.recieveToken.asObservable()
  }

  ngOnDestroy(): void {
  }
}
