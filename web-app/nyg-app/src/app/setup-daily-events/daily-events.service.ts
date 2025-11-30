import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { ComponentRef, Injectable } from '@angular/core';
import { SetupDailyEventsComponent } from './setup-daily-events.component';
import { ComponentPortal } from '@angular/cdk/portal';
import { IDailyEvents } from '../UI/daily-events/daily-events.component';
import { ReplaySubject } from 'rxjs';

export interface Reward {
  isCoin: boolean
  reward: string
  index: number
}

@Injectable({
  providedIn: 'root'
})
export class DailyEventsService {

  private ____OR: OverlayRef | null = null
  private monitorLocifyReward: ReplaySubject<Reward> = new ReplaySubject<Reward>()
  private monitorClashReward: ReplaySubject<Reward> = new ReplaySubject<Reward>()
  constructor(private overaly: Overlay) { }
  Setup(Clashevents: IDailyEvents[], Locifyevents: IDailyEvents[],) {
    this.closeSetup(); // stop rendering

    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerVertically().centerHorizontally(),
      hasBackdrop: true,
      panelClass: "DAILYEVENT",
      backdropClass: "DBack",
    })


    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(SetupDailyEventsComponent, null,)
    const __ref = this.____OR.attach(portal)


    __ref.instance.NYGclashEvents = Clashevents
    __ref.instance.NYGlocifyEvents = Locifyevents

    __ref.instance.monitorRewardLocify.subscribe({
      next: (token) => {
        this.monitorLocifyReward.next(token)
      }
    })

    __ref.instance.monitorRewardClash.subscribe({
      next: (token) => {
        this.monitorClashReward.next(token)
      }
    })

    this.____OR.backdropClick().subscribe(() => {
      this.closeSetup()
    })
    __ref.changeDetectorRef.detectChanges()

    return true
  }

  getLocifyRewards() {
    return this.monitorLocifyReward.asObservable()
  }

  getClashRewards() {
    return this.monitorClashReward.asObservable()
  }

  closeSetup() {
    if (this.____OR) {
      this.____OR.dispose()
      this.____OR = null
    }
  }
}
