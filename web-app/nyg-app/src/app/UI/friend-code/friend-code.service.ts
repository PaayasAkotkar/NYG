import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { Injectable } from '@angular/core';
import { ComponentPortal } from '@angular/cdk/portal';
import { FriendCodeComponent } from './friend-code.component';

@Injectable({
  providedIn: 'root'
})
export class FriendCodeService {

  private ____OR: OverlayRef | null
  constructor(private overaly: Overlay) { }
  Setup(code: string) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().bottom().right(),
      hasBackdrop: false,
      panelClass: "FRIENDCODE",
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(FriendCodeComponent, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.NYGfriendCode = code

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
