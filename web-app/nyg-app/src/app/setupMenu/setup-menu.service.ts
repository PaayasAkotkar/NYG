import { GlobalPositionStrategy, Overlay, OverlayConfig, OverlayRef } from "@angular/cdk/overlay";
import { ComponentPortal } from "@angular/cdk/portal";
import { HostListener, Injectable } from "@angular/core";
import { MenuComponent } from "../menu/menu.component";
import { ViewMenuComponent } from "../view-menu/view-menu.component";

@Injectable({
  providedIn: 'root',
})
export class SetupMenuService {
  private ____OR: OverlayRef | null = null
  constructor(private overaly: Overlay) { }
  Setup() {
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().bottom().centerHorizontally(),
      hasBackdrop: false,
      panelClass: "MENU",
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(ViewMenuComponent, null,)
    const __ref = this.____OR.attach(portal)

    return this.____OR
  }
  closeSetup() {
    if (this.____OR) {
      this.____OR.detach()
      this.____OR = null
    }
  }
  Create() {
    if (!this.HIDE) {
      this.Setup()
    } else {
      this.closeSetup()
    }
  }
  HIDE: boolean = false
  @HostListener('document:keydown', ['$event'])
  HideMenu(event: KeyboardEvent) {
    if (event.key == 'Control') {
      this.HIDE = !this.HIDE
    }
  }
}