import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { Injectable } from '@angular/core';
import { OcsManualComponent } from './ocs-manual.component';
import { ComponentPortal } from '@angular/cdk/portal';
import { NavigationStart, Router } from '@angular/router';
import { filter } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class OcsManualService {
  private ____OR: OverlayRef | null = null

  constructor(private overaly: Overlay, private router: Router) { }
  Setup() {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerVertically().centerHorizontally(),
      hasBackdrop: true,
      backdropClass: "VIEWBC",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(OcsManualComponent, null,)
    const __ref = this.____OR.attach(portal)
    this.____OR.backdropClick().subscribe(() => {
      this.closeSetup()
    })
    this.router.events
      .pipe(

        filter(event => event instanceof NavigationStart),
      )
      .subscribe(() => {
        this.closeSetup()
      });
    return this.____OR
  }
  closeSetup() {
    if (this.____OR) {
      this.____OR.dispose()
      this.____OR = null
    }
  }
}
