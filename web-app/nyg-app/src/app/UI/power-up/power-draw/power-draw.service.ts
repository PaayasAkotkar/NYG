import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { Injectable } from '@angular/core';
import { ComponentPortal } from '@angular/cdk/portal';
import { PowerDrawComponent } from './power-draw.component';
import { Subscription, take, timer } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class DrawPowerService {
  private ____OR: OverlayRef | null
  constructor(private overaly: Overlay) { }
  Setup(room: string, myteam: string) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerHorizontally().centerVertically().top(),
      hasBackdrop: false,
      panelClass: "POWERNEXUS",
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(PowerDrawComponent, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.NYGmyRoom = room
    __ref.instance.NYGmyteam = myteam
    // only invoke for 3sec

    return this.____OR
  }
  closeSetup() {
    if (this.____OR) {
      this.____OR.dispose()
      this.____OR = null
    }
  }
}
