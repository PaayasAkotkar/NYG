import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { Injectable } from '@angular/core';
import { IDailyEvents } from '../daily-events/daily-events.component';
import { ComponentPortal } from '@angular/cdk/portal';
import { HomeProfileComponent, IHomePageProfile } from './home-profile.component';

@Injectable({
  providedIn: 'root'
})
export class HomeProfileService {

  private ____OR: OverlayRef | null = null
  constructor(private overaly: Overlay) { }
  Setup(profile: IHomePageProfile) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerVertically().centerHorizontally(),
      hasBackdrop: true,
      backdropClass: "PDBack",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(HomeProfileComponent, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.NYGProfile = profile
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
