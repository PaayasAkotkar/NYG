import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { Injectable, OnDestroy, WritableSignal } from '@angular/core';
import { EntranceComponent } from './entrance.component';
import { GameMiscelleanous, NamesAsPerTeams, NamesAsPerTeamsRes } from '../../lobby/resources/lobbyResource';

@Injectable({
  providedIn: 'root'
})
export class EntranceService {
  private ____OR: OverlayRef | null = null
  constructor(private overaly: Overlay) { }

  Setup(playersAgainst: GameMiscelleanous, procced: WritableSignal<boolean>) {
    if (this.____OR?.hasAttached()) {
      this.closeSetup()
    }
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerVertically().centerHorizontally(),
      panelClass: "ENTRANCE",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(EntranceComponent, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.OCSteamPlayers = playersAgainst
    __ref.instance.OCSproceed = procced
    return this.____OR
  }
  closeSetup() {
    if (this.____OR) {
      this.____OR.dispose()
      this.____OR = null
    }
  }

}
