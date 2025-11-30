import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { Injectable, signal, WritableSignal } from '@angular/core';
import { ComponentPortal } from '@angular/cdk/portal';
import { PowerBetComponent } from './power-bet.component';
import { OCSFieldFormat } from '../../../helpers/patterns';
import { Subscriber, Subscription, take, timer } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class BetPowerService {
  private channel = new Subscription()
  private ____OR: OverlayRef | null
  constructor(private overaly: Overlay) { }
  Setup(room: string, collection: OCSFieldFormat[], myTeam: string, clash: WritableSignal<boolean>) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerHorizontally().centerVertically().top(),
      hasBackdrop: false,
      panelClass: "POWERNEXUS",
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(PowerBetComponent, null,)
    const __ref = this.____OR.attach(portal)
    var s: WritableSignal<OCSFieldFormat[]> = signal(collection)
    __ref.instance.NYGmyRoom = room
    __ref.instance.NYGcollection = s
    __ref.instance.NYGmyTeam = myTeam
    __ref.instance.NYGclash = clash
    // only invoke for 3sec
    this.channel = timer(3000).pipe(take(1)).subscribe(
      {
        next: () => {
          this.closeSetup()
        }
      }
    )
    return this.____OR
  }
  closeSetup() {
    if (this.____OR) {
      this.____OR.dispose()
      this.____OR = null
    }
  }
}
