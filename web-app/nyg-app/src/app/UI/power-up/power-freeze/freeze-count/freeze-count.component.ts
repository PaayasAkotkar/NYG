import { Component, OnInit, EventEmitter, Output, signal, WritableSignal, Input, OnDestroy, effect, NgZone } from '@angular/core';
import { TimerService } from '../../../../timer/timer.service';
import { BehaviorSubject, interval, scan, startWith, Subject, Subscription, switchMap, takeUntil, takeWhile, tap } from 'rxjs';
import { FreezeCountService } from './freeze-count.service';
import { RoomService } from '../../../../roomService/room.service';
import { FreezeTimerService } from './freeze-timer.service';

@Component({
  selector: 'app-freeze-count',
  imports: [],
  templateUrl: './freeze-count.component.html',
  styleUrl: './freeze-count.component.scss'
})
export class FreezeCountComponent implements OnInit, OnDestroy {

  freezeAccess: WritableSignal<number> = signal(15)
  @Input({ required: true }) freezeFor: number = 15
  channel = new Subscription()
  stop = new Subject<void>()
  @Output() Watcher = new EventEmitter<boolean>()
  @Input({ required: false }) NYGStopSignal: WritableSignal<boolean> = signal(false)
  @Input({ required: false }) NYGroomName: string = ""
  @Input({ required: true }) NYGclash: WritableSignal<boolean>
  ngOnDestroy(): void {
    this.stop.next()
    this.stop.complete()
    this.fc.closeSetup()
    this.timer.StopTimer()
  }

  ngOnInit(): void {
    this.Zone.runOutsideAngular(() => {

      if (!this.NYGStopSignal()) {
        var min = 0
        var sec = this.freezeFor
        var speed = 1000
        this.timer.Initialize(0, sec, speed)
        this.timer.Start(min, sec, speed)
        this.timer.Watch().pipe(takeUntil(this.stop)).subscribe({
          next: (token) => {
            this.freezeAccess.set(token.sec)
            if (token.sec == 0) {
              this.room.Unfreeze(this.NYGroomName, this.NYGclash())
              this.fc.closeSetup()
              this.Watcher.emit(true)
            }
          }
        })
      } else {
        this.room.Unfreeze(this.NYGroomName, this.NYGclash())
        this.fc.closeSetup()
      }

    })
  }

  constructor(private timer: FreezeTimerService, private fc: FreezeCountService, private room: RoomService, private Zone: NgZone) {
  }

}
