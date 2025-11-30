import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { Injectable, WritableSignal } from '@angular/core';
import { ComponentPortal } from '@angular/cdk/portal';
import { PowerNexusComponent } from './power-nexus.component';

@Injectable({
  providedIn: 'root'
})
export class NexusPowerService {

  private ____OR: OverlayRef | null
  constructor(private overaly: Overlay) { }
  Setup(Nexus: WritableSignal<string>) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerHorizontally().centerVertically().top(),
      hasBackdrop: false,
      panelClass: "POWERNEXUS",
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(PowerNexusComponent, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.Nexus = Nexus
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
