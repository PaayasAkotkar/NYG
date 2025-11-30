import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { Injectable, WritableSignal } from '@angular/core';
import { HAND } from '../../helpers/hand';
import { ListComponent } from './list.component';
import { OCSFieldFormat } from '../../helpers/patterns';
import { ComponentPortal } from '@angular/cdk/portal';

@Injectable({
  providedIn: 'root'
})
export class NYGlistService {
  private ____OR: OverlayRef | null = null
  constructor(private overlay: Overlay,) { }
  Setup(list: WritableSignal<OCSFieldFormat[]>, header: string, nygMode: string, proceed: WritableSignal<boolean>) {
    this.closeSetup();
    const config = new OverlayConfig({
      positionStrategy: this.overlay.position().global().centerHorizontally().centerVertically(),
      hasBackdrop: true,
      backdropClass: "LBack",
    })
    this.____OR = this.overlay.create(config)
    const portal = new ComponentPortal(ListComponent, null,)
    const __ref = this.____OR.attach(portal)
    this.____OR.backdropClick().subscribe(() => {
      this.closeSetup()
    })

  }

  private closeSetup() {
    if (this.____OR) {
      this.____OR.dispose()
      this.____OR = null
    }
  }
}
