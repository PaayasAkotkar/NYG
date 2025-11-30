import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { Injectable, WritableSignal } from '@angular/core';
import { ComponentPortal } from '@angular/cdk/portal';
import { BetGuessComponent } from './bet-guess.component';
import { Subscription, take, timer } from 'rxjs';
import { RecBetItems } from '../../../../lobby/resources/lobbyResource';
@Injectable({
  providedIn: 'root'
})
export class BetGuessPowerService {
  private channel = new Subscription()
  private ____OR: OverlayRef | null
  constructor(private overaly: Overlay) { }
  Setup(room: string, collect: WritableSignal<RecBetItems>, clash: WritableSignal<boolean>) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerHorizontally().centerVertically().top(),
      hasBackdrop: false,
      panelClass: "POWERNEXUS",
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(BetGuessComponent, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.NYGmyRoom = room
    __ref.instance.NYGcollection = collect
    __ref.instance.NYGclash = clash
    return this.____OR
  }
  closeSetup() {
    if (this.____OR) {
      this.____OR.dispose()
      this.____OR = null
    }
  }
}
