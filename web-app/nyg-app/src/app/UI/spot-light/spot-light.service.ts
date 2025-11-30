import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { Injectable, NgZone, OnDestroy } from '@angular/core';
import { Subject, Subscription, take, takeUntil, timer } from 'rxjs';
import { SpotLightComponent } from './spot-light.component';

@Injectable({
  providedIn: 'root'
})
export class SpotLightService {

  private ____OR: OverlayRef | null

  constructor(private overaly: Overlay, private run: NgZone) { }

  Setup() {
    if (this.____OR?.hasAttached()) {
      this.closeSetup()
    }
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerHorizontally().centerVertically(),
      panelClass: "GAMESTART",
      hasBackdrop: true,
      backdropClass: "STARTUPTIMER",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(SpotLightComponent, null,)
    const __ref = this.____OR.attach(portal)
    return this.____OR
  }

  closeSetup() {
    if (this.____OR) {
      this.____OR.dispose()
      this.____OR = null
    }
  }
}
