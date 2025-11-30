import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { Injectable } from '@angular/core';
import { PowerTagComponent } from './power-tag.component';
import { Subscription, timer } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class PowerTagService {


  private ____OR: OverlayRef | null
  private channel: Subscription = new Subscription()

  constructor(private overaly: Overlay) { }

  Setup() {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerHorizontally().centerVertically().top(),
      hasBackdrop: false,
      panelClass: "POWERTAG",
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(PowerTagComponent, null,)
    const __ref = this.____OR.attach(portal)
    timer(2860).subscribe({
      next: () => {
        this.closeSetup()
      }
    })
    return this.____OR
  }

  closeSetup() {
    if (this.____OR) {
      this.____OR.dispose()
      this.____OR = null
    }
  }
}
