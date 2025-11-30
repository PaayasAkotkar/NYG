import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { Injectable } from '@angular/core';
import { OcsInputComponent } from './ocs-input.component';
import { FormControl, FormGroup } from '@angular/forms';

@Injectable({
  providedIn: 'root'
})
export class OcsInputService {

  private ____OR: OverlayRef | null = null
  constructor(private overaly: Overlay) { }
  Setup(OCSheader: string, OCSformControl: string, formGroup: FormGroup) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerVertically().centerHorizontally(),
      hasBackdrop: true,
      panelClass: "VIEW",
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(OcsInputComponent, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.NYGformControl_ = formGroup.get(OCSformControl) as FormControl
    __ref.instance.NYGheader = OCSformControl
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
