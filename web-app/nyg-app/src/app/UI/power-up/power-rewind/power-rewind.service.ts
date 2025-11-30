import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { Injectable } from '@angular/core';
import { PowerRewindComponent } from './power-rewind.component';
import { Subscription, take, timer } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class PowerRewindService {

  private ____OR: OverlayRef | null
  public duration: number = 1
  private channel: Subscription = new Subscription()
  constructor(private overaly: Overlay) { }
  Setup() {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerHorizontally().centerVertically().top(),
      hasBackdrop: false,
      panelClass: "POWERNEXUS",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(PowerRewindComponent, null,)
    const __ref = this.____OR.attach(portal)
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
