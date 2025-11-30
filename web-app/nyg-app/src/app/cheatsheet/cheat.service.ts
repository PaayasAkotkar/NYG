import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { Injectable, WritableSignal } from '@angular/core';
import { CheatsheetComponent } from './cheatsheet.component';
import { Rounds, SheetUpdate } from '../lobby/resources/lobbyResource';

@Injectable({
  providedIn: 'root'
})
export class CheatService {

  private ____OR: OverlayRef | null = null
  constructor(private overaly: Overlay) { }
  Setup(OCSnicknames: string[], OCSupdateSheet: Record<string, Record<number, string>>, OCSredScore: number, OCSblueScore: number) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerVertically().centerHorizontally(),
      hasBackdrop: true,
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(CheatsheetComponent, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.NYGplayersNickNames = OCSnicknames
    __ref.instance.NYGupdateSheet = OCSupdateSheet
    __ref.instance.NYGredScore = OCSredScore
    __ref.instance.NYGblueScore = OCSblueScore

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
