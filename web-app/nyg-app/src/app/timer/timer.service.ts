import { isPlatformBrowser } from '@angular/common';
import { ChangeDetectorRef, Inject, Injectable, NgZone, OnChanges, OnDestroy, PLATFORM_ID, SimpleChanges } from '@angular/core';
import { BehaviorSubject, endWith, filter, iif, interval, map, Observable, of, scan, startWith, Subject, Subscription, switchMap, takeUntil, takeWhile, tap } from 'rxjs';

interface RewindTime {
  speed: number
}

@Injectable({
  providedIn: 'root'
})
export class TimerService implements OnDestroy {

  private StartFastChannel: BehaviorSubject<number> = new BehaviorSubject<number>(0)
  private StartRewindChannel: BehaviorSubject<RewindTime>

  private recieveRewindToken = new BehaviorSubject<number>(0)
  private recieveToken = new BehaviorSubject<number>(0)
  private StopFastChannel: Subject<void> = new Subject<void>()
  private StopRewindChannel: Subject<void> = new Subject<void>()

  private stopSignal: boolean = false

  constructor(@Inject(PLATFORM_ID) private pID: object,) { }

  SignalStop(signal: boolean) {
    this.stopSignal = signal
  }

  UnSubAll() {
    if (!this.StartFastChannel.closed)
      this.StartFastChannel.unsubscribe()
  }

  ngOnDestroy(): void {
    this.MarkComplete()
    this.UnSubAll()
  }

  MarkComplete() {

    this.StopFastChannel.next()
    this.StopFastChannel.complete()

    this.StopRewindChannel.next()
    this.StopRewindChannel.complete()
  }

  // StartFast creates a timer [min, max]=> [10,10] is the key
  // where the clock starts in 10-1=>9-1=>8...
  StartFast(min: number, max: number, speed: number) {
    if (isPlatformBrowser(this.pID) && !this.stopSignal) {
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
        takeWhile(token => token >= 0),
        takeUntil(this.StopFastChannel),
      )
        .subscribe({
          next(value) {
          },
        })

    }
  }

  /***
   * 
   * RewindTime returns live updates of value from x....y and than y...x
   * note: final time > intial time
   */
  RewindTime(initialTime: number, finalTime: number, seeUpdatingDecrementing: number, seeUpdatingIncrement: number) {
    if (isPlatformBrowser(this.pID) && !this.stopSignal) {

      let speed = seeUpdatingDecrementing // initial speed

      // important for speed
      this.StartRewindChannel = new BehaviorSubject<RewindTime>({ speed: speed })

      this.StartRewindChannel.pipe(
        switchMap(({ speed }) => {
          return interval(speed).pipe(
            startWith(initialTime),
            // [5,10]=> acc==10 acc+1 if and only if it has reached the final time 
            scan(acc => acc === finalTime ? acc -= 1 : finalTime === 0 ? acc -= 1 : acc += 1, initialTime),
            takeWhile(count => count >= 0 && count <= 10),
            tap((token) => {
              initialTime = token
              this.recieveRewindToken.next(token)
              // pivot 
              if (token == finalTime) {
                finalTime = 0
                this.StartRewindChannel.next({ speed: seeUpdatingIncrement })
              }


            }),
          )

        }),
        takeWhile(token => token > 0),
        takeUntil(this.StopRewindChannel),
      )
        .subscribe({
          next(value) {

          },
        })

    }

  }

  Start(): Observable<number> {
    return this.recieveToken.asObservable()
  }

  ReStart(): Observable<number> {
    return this.recieveRewindToken.asObservable()
  }

}