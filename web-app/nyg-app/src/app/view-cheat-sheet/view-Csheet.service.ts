import { GlobalPositionStrategy, Overlay, OverlayConfig, OverlayRef } from "@angular/cdk/overlay";
import { ComponentPortal } from "@angular/cdk/portal";
import { Injectable, Injector } from "@angular/core";
import { StatboardComponent } from "../statboard/statboard.component";
import { CheatsheetComponent } from "../cheatsheet/cheatsheet.component";

@Injectable({
    providedIn: 'root',
})
export class ViewCSheetService {
    private ____OR: OverlayRef | null
    constructor(private overaly: Overlay, private injector: Injector) { }
    Setup(update: Map<string, Map<string, string>>, nn: string[]) {
        const config = new OverlayConfig({
            positionStrategy: this.overaly.position().global().centerHorizontally().centerVertically(),
            hasBackdrop: true,
            panelClass: "SHEET",
            backdropClass: "BackDSHEET",
        })
        this.____OR = this.overaly.create(config)
        const portal = new ComponentPortal(CheatsheetComponent, null,)
        const __ref = this.____OR.attach(portal)
        __ref.instance.NYGplayersNickNames = nn
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