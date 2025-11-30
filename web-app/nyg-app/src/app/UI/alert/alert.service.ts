import { Injectable } from '@angular/core';
import { AlertComponent } from './alert.component';
import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';

@Injectable({
  providedIn: 'root'
})
export class AlertService {
  private ____OR: OverlayRef | null
  constructor(private overaly: Overlay) { }
  Setup(msg: string, title: string) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerHorizontally().centerVertically().top(),
      hasBackdrop: false,
      panelClass: "POWERNEXUS",
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(AlertComponent, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.OCStitle = title
    __ref.instance.OCSmessage = msg
    this.____OR.backdropClick().subscribe(() => {
      this.closeSetup()
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
