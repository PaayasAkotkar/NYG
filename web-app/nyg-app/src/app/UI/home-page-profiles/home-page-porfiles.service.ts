import { Injectable } from '@angular/core';
import { IHomePageProfile } from '../home-profile/home-profile.component';
import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { HomePageProfilesComponent } from './home-page-profiles.component';

@Injectable({
  providedIn: 'root'
})
export class HomePagePorfilesService {

  private ____OR: OverlayRef | null = null
  constructor(private overaly: Overlay) { }
  Setup(clashProfile: IHomePageProfile, onevoneProfile: IHomePageProfile, twovtwoProfile: IHomePageProfile) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerVertically().centerHorizontally(),
      hasBackdrop: true,
      backdropClass: "PDBack",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(HomePageProfilesComponent, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.NYGclashProfile = clashProfile
    __ref.instance.NYGonevoneProfile = onevoneProfile
    __ref.instance.NYGtwovtwoProfile = twovtwoProfile

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
