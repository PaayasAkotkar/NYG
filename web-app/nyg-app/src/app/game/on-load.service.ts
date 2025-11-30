import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { Injectable, NgZone, WritableSignal } from '@angular/core';
import { concatMap, tap, timer } from 'rxjs';
import { EntranceComponent } from '../UI/entrance/entrance.component';
import { SpotLightComponent } from '../UI/spot-light/spot-light.component';
import { GameMiscelleanous, NamesAsPerTeamsRes } from '../lobby/resources/lobbyResource';

@Injectable({
  providedIn: 'root'
})
export class OnLoadService {

  private ____OR: OverlayRef | null = null
  constructor(private overaly: Overlay, private Zone: NgZone) { }
  Setup(teamPlayers: GameMiscelleanous, GuessSheet: WritableSignal<string>, firstInterval: number = 10000, SecondInterval: number = 300) {
    this.closeSetup() // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerHorizontally().centerVertically(),
      panelClass: "GAMESTART",
      hasBackdrop: true,
      backdropClass: "STARTUPTIMER",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(SpotLightComponent, null,)
    const __ref = this.____OR.attach(portal)

    this.Zone.runOutsideAngular(() => {
      this.Zone.run(() => {
        timer(0).pipe(
          concatMap(() => {
            return timer(firstInterval).pipe(
              tap(() => {
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
                __ref.instance.OCSteamPlayers = teamPlayers
              })
            )
          }),

          concatMap(() => {
            return timer(SecondInterval).pipe(
              tap(() => {
                if (this.____OR?.hasAttached()) {
                  this.closeSetup()
                }
              })
            )
          }),
        ).subscribe()
      })

    })

  }
  closeSetup() {
    if (this.____OR) {
      this.____OR.dispose()
      this.____OR = null
    }
  }
}
