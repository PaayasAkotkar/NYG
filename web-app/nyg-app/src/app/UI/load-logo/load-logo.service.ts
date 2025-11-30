import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { Injectable } from '@angular/core';
import { LoadLogoComponent } from './load-logo.component';
import { ComponentPortal } from '@angular/cdk/portal';

@Injectable({
  providedIn: 'root'
})
export class LoadLogoService {

  private ____OR: OverlayRef | null = null

  constructor(private overaly: Overlay) { }

  Setup() {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerHorizontally().top(),
      hasBackdrop: false,
      panelClass: "LOGO",
      backdropClass: "BackDropXXX",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(LoadLogoComponent, null,)
    const __ref = this.____OR.attach(portal)
    // this.____OR.backdropClick().subscribe(() => {
    //   this.closeSetup()
    // })
    return this.____OR
  }
  closeSetup() {
    if (this.____OR && this.____OR.hasAttached()) {
      this.____OR.dispose()
      this.____OR = null
    }
  }
}
