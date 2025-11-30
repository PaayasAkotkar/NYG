import { GlobalPositionStrategy, Overlay, OverlayConfig, OverlayRef } from "@angular/cdk/overlay";
import { ComponentPortal } from "@angular/cdk/portal";
import { Injectable } from "@angular/core";
import { MenuComponent } from "../menu/menu.component";
import { SetupGameModeComponent } from "../setup-game-mode/setup-game-mode.component";

@Injectable({
  providedIn: 'root',
})
export class GameModeService {
  private ____OR: OverlayRef | null
  constructor(private overaly: Overlay) { }
  Setup(view: string = "") {
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global()
        .centerHorizontally().centerVertically(),
      hasBackdrop: true,
      panelClass: "mode-panel-class",
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(SetupGameModeComponent, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.view = view
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