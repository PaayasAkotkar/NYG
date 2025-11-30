import { GlobalPositionStrategy, Overlay, OverlayConfig, OverlayRef } from "@angular/cdk/overlay";
import { ComponentPortal } from "@angular/cdk/portal";
import { Injectable } from "@angular/core";
import { MenuComponent } from "../menu/menu.component";

@Injectable({
    providedIn: 'root',
})
export class MenuService {
    private ____OR: OverlayRef | null
    constructor(private overaly: Overlay) { }
    Setup() {
        const config = new OverlayConfig({
            positionStrategy: this.overaly.position().global()
                .centerHorizontally().centerVertically().bottom(),
            hasBackdrop: true,
            panelClass: "menu-panel-class",
            backdropClass: "BackDSHEET",
        })
        this.____OR = this.overaly.create(config)
        const portal = new ComponentPortal(MenuComponent, null,)
        const __ref = this.____OR.attach(portal)
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