import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { Injectable, WritableSignal } from '@angular/core';
import { ComponentPortal } from '@angular/cdk/portal';
import { RoomSettingsComponent } from './room-settings.component';
import { OCSFieldFormat } from '../../helpers/patterns';
import { RoomSettings } from '../../lobby/resources/lobbyResource';
@Injectable({
  providedIn: 'root'
})
export class RoomSettingsService {

  private ____OR: OverlayRef | null
  constructor(private overaly: Overlay) { }
  Setup(OCScollection
    : WritableSignal<OCSFieldFormat[]>,
    isEntertainment: WritableSignal<boolean>,
    power: Record<string, boolean>,
    Maxlength: number, roomSettings: RoomSettings
  ) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerHorizontally().centerVertically(),
      hasBackdrop: true,
      panelClass: "ROOMSETTINGS",
      backdropClass: "ROOMSETTINGSBACKDROP",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(RoomSettingsComponent, null,)
    const __ref = this.____OR.attach(portal)

    __ref.instance.Entertainment = isEntertainment
    // __ref.instance.NYGsportsCollection = sportsCollections
    // __ref.instance.NYGentertainmentCollection = entertainmentCollections
    __ref.instance.NYGcollection = OCScollection
    __ref.instance.Powers = power
    __ref.instance.NYGroomSettings = roomSettings
    __ref.instance.Maxlength = Maxlength
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
