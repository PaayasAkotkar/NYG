import { GlobalPositionStrategy, Overlay, OverlayConfig, OverlayRef } from "@angular/cdk/overlay";
import { ComponentPortal } from "@angular/cdk/portal";
import { Injectable } from "@angular/core";
import { MenuComponent } from "../../menu/menu.component";
import { ViewMenuComponent } from "../../view-menu/view-menu.component";
import { FieldComponent } from "../../field/field/field.component";

@Injectable({
  providedIn: 'root',
})
export class SportsService {
  private ____OR: OverlayRef | null = null
  constructor(private overaly: Overlay) { }
  Setup() {
    this.closeSetup()
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global()
        .centerHorizontally().centerVertically(),
      hasBackdrop: true,
      panelClass: "sports-class",
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(FieldComponent, null,)
    const __ref = this.____OR.attach(portal)
    return this.____OR
  }
  closeSetup() {
    if (this.____OR) {
      this.____OR.dispose()
      this.____OR = null
    }
  }
}