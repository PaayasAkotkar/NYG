import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { Injectable, WritableSignal } from '@angular/core';
import { ComponentPortal } from '@angular/cdk/portal';
import { PowerFreezeComponent } from './power-freeze.component';

@Injectable({
  providedIn: 'root'
})
export class FreezePowerService {

  private ____OR: OverlayRef | null
  constructor(private overaly: Overlay) { }
  Setup(room: string, freezeAccess: number, _unFreezeControl: WritableSignal<boolean>, stop: WritableSignal<boolean>, unfreeze: WritableSignal<boolean>, clash: WritableSignal<boolean>) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerHorizontally().top(),
      hasBackdrop: false,
      panelClass: "POWERFREEZE",
      backdropClass: "freezeBackdrop",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(PowerFreezeComponent, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.NYGmyRoom = room
    __ref.instance.NYGunfreezeControl = _unFreezeControl
    __ref.instance.NYGunfreeze = unfreeze
    __ref.instance.NYGclash = clash
    __ref.instance.stop = stop
    __ref.instance.NYGfreezeAccess = freezeAccess
    return this.____OR
  }
  closeSetup() {
    if (this.____OR) {
      this.____OR.dispose()
      this.____OR = null
    }
  }
}
