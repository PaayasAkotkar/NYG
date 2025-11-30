import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { Injectable } from '@angular/core';
import { ViewPatternComponent } from './view-pattern.component';

@Injectable({
  providedIn: 'root'
})
export class SetScoreboardService {


  private ____OR: OverlayRef | null = null
  constructor(private overaly: Overlay) { }
  Setup() {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerVertically().centerHorizontally(),
      hasBackdrop: true,
      panelClass: "VIEW",
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(ViewPatternComponent, null,)
    const __ref = this.____OR.attach(portal)
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
