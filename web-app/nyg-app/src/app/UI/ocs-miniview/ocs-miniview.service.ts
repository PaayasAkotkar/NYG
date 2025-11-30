import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { Injectable } from '@angular/core';
import { OcsMiniviewComponent } from './ocs-miniview.component';
import { IProfile_, JoinedProfile } from '../../lobby/resources/lobbyResource';

@Injectable({
  providedIn: 'root'
})
export class OcsMiniviewService {

  private ____OR: OverlayRef | null = null
  constructor(private overaly: Overlay) { }
  Setup(profile: IProfile_) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerVertically().centerHorizontally(),
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(OcsMiniviewComponent, null,)
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
