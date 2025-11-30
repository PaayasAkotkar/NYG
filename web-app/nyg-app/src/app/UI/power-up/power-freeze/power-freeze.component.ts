import { Component, effect, Input, NgZone, OnDestroy, OnInit, signal, WritableSignal } from '@angular/core';
import { LobbyResource } from '../../../lobby/resources/lobbyResource';
import { RoomService } from '../../../roomService/room.service';
import { TimerService } from '../../../timer/timer.service';
import { single, Subscription } from 'rxjs';
import { FreezeCountService } from './freeze-count/freeze-count.service';
import { FreezePowerService } from './power-freeze.service';

@Component({
  selector: 'app-power-freeze',
  imports: [],
  templateUrl: './power-freeze.component.html',
  styleUrls: ['./touchup/freeze.scss', './touchup/animation.scss', './touchup/color.scss'],
})
export class PowerFreezeComponent implements OnInit, OnDestroy {
  constructor(private fc: FreezeCountService, private fcs: FreezePowerService, private __zone: NgZone, private room: RoomService, private timer: TimerService) {
    effect(() => {
      if (this.NYGunfreeze())
        this.Close()
    })
  }
  @Input() NYGref: boolean = false
  @Input({ required: false }) NYGmyRoom: string = ""
  @Input({ required: false }) NYGunfreezeControl: WritableSignal<boolean> = signal(false)
  @Input({ required: false }) NYGunfreeze: WritableSignal<boolean> = signal(false)
  @Input({ required: false }) NYGclash: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGfreezeAccess: number = 0

  @Input({ required: false }) stop: WritableSignal<boolean> = signal(false)
  Close() {
    this.fcs.closeSetup()
    this.fc.closeSetup()

  }
  ngOnDestroy(): void {
    this.fcs.closeSetup()
    this.fc.closeSetup()
  }

  UnFreeze() {
    if (!this.NYGref) {
      this.NYGunfreezeControl.set(false)
      this.room.Unfreeze(this.NYGmyRoom, this.NYGclash())
    }
  }

  ngOnInit(): void {
    if (!this.NYGref) {
      this.fc.Setup(this.NYGfreezeAccess, this.stop, this.NYGclash, this.NYGmyRoom)
      this.fc.Listen().subscribe((token) => {
        if (token == true)
          this.UnFreeze()
      })
    }
  }
}
