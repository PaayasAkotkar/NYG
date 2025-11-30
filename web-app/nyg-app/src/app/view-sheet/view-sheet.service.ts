import { GlobalPositionStrategy, Overlay, OverlayConfig, OverlayRef } from "@angular/cdk/overlay";
import { ComponentPortal } from "@angular/cdk/portal";
import { Injectable, Injector } from "@angular/core";
import { StatboardComponent } from "../statboard/statboard.component";

@Injectable({
    providedIn: 'root',
})
export class ViewSheetService {
    private ____OR: OverlayRef | null = null
    constructor(private overaly: Overlay,) { }
    Setup(rating: string, name: string, teamname: any) {
        const config = new OverlayConfig({
            hasBackdrop: true,
            panelClass: "SHEET2",
            backdropClass: "BackDSHEET",
        })

        this.____OR = this.overaly.create(config)
        const portal = new ComponentPortal(StatboardComponent, null,)
        const __ref = this.____OR.attach(portal)

        __ref.instance.playerName = name
        __ref.instance.teamName = teamname
        __ref.instance.rating = rating

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