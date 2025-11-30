import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { Injectable, WritableSignal } from '@angular/core';
import { PowerUpComponent } from './power-up.component';

@Injectable({
  providedIn: 'root'
})
export class PowerService {

  private ____OR: OverlayRef | null
  constructor(private overaly: Overlay) { }
  Setup(
    roomName: string,
    OCSclash: WritableSignal<boolean>,
    OCSnexus: WritableSignal<boolean>,
    OCSrewind: WritableSignal<boolean>,
    OCSfreeze: WritableSignal<boolean>,
    OCSdraw: WritableSignal<boolean>,
    OCStag: WritableSignal<boolean>,
    OCSbet: WritableSignal<boolean>,
    OCScovert: WritableSignal<boolean>,

    OCSblock: WritableSignal<boolean>,

    OCSincludeR: WritableSignal<boolean>,
    OCSincludeD: WritableSignal<boolean>,
    OCSincludeT: WritableSignal<boolean>,
    OCSincludeN: WritableSignal<boolean>,
    OCSincludeF: WritableSignal<boolean>,
    OCSincludeC: WritableSignal<boolean>,
    OCSincludeB: WritableSignal<boolean>,
    canUsePower: WritableSignal<boolean>,
  ) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerHorizontally().centerVertically(),
      hasBackdrop: true,
      backdropClass: "CL",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(PowerUpComponent, null,)
    const __ref = this.____OR.attach(portal)

    __ref.instance.NYGroomname = roomName
    __ref.instance.NYGincludeB = OCSincludeB
    __ref.instance.NYGincludeN = OCSincludeN
    __ref.instance.NYGincludeD = OCSincludeD
    __ref.instance.NYGincludeT = OCSincludeT
    __ref.instance.NYGincludeR = OCSincludeR
    __ref.instance.NYGincludeC = OCSincludeC
    __ref.instance.NYGincludeT = OCSincludeT
    __ref.instance.NYGincludeF = OCSincludeF
    __ref.instance.NYGcanUsePower = canUsePower
    __ref.instance.NYGblock = OCSblock
    __ref.instance.NYGnexus = OCSnexus
    __ref.instance.NYGbet = OCSbet
    __ref.instance.NYGdraw = OCSdraw
    __ref.instance.NYGtag = OCStag
    __ref.instance.NYGcovert = OCScovert
    __ref.instance.NYGrewind = OCSrewind
    __ref.instance.NYGdraw = OCSdraw
    __ref.instance.NYGfreeze = OCSfreeze
    __ref.instance.NYGclash = OCSclash
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
