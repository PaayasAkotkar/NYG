import { GlobalPositionStrategy, Overlay, OverlayConfig, OverlayRef } from "@angular/cdk/overlay";
import { ComponentPortal } from "@angular/cdk/portal";
import { Injectable } from "@angular/core";
import { FrameComponent } from "../../frame/frame.component";
import { CheatsheetComponent } from "../../cheatsheet/cheatsheet.component";
import { MaterialComponent } from "../../material/material.component";

@Injectable({
  providedIn: 'root',
})
export class LftOverService {
  private ____OR: OverlayRef | null = null
  constructor(private overaly: Overlay) { }
  Setup() {
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global()
        .centerHorizontally().centerVertically(),
      hasBackdrop: true,
      disposeOnNavigation: true,
      panelClass: "lft-class",
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(MaterialComponent, null,)
    const __ref = this.____OR.attach(portal)
    this.____OR.backdropClick().subscribe(() => {
      this.closeSetup()
      console.log("working")
    })
    return this.____OR
  }
  closeSetup() {
    if (this.____OR) {
      this.____OR.detach()
      this.____OR = null
    }
  }
}