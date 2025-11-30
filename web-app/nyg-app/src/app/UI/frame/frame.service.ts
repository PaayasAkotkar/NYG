import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { Injectable, signal, WritableSignal } from '@angular/core';
import { OCSFrameComponent } from './frame.component';
import { OCSFieldFormat } from '../../helpers/patterns';

@Injectable({
  providedIn: 'root'
})
export class FrameService {

  private ____OR: OverlayRef | null = null
  constructor(private overaly: Overlay) { }
  Setup(frame: WritableSignal<OCSFieldFormat[]>,) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerVertically().centerHorizontally(),
      hasBackdrop: true,
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(OCSFrameComponent, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.OCSproceed = signal(true)
    __ref.instance.OCSlist = frame
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
