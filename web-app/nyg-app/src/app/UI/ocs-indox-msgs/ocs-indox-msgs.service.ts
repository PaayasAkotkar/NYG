import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { Injectable } from '@angular/core';
import { IMessages, OcsIndoxMsgsComponent } from './ocs-indox-msgs.component';
import { IMessage } from '../../lobby/resources/lobbyResource';

@Injectable({
  providedIn: 'root'
})
export class OcsIndoxMsgsService {

  private ____OR: OverlayRef | null = null
  constructor(private overaly: Overlay) { }
  Setup(msgs: IMessages[]) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerVertically().centerHorizontally(),
      hasBackdrop: true,
      panelClass: "VIEWmsg",
      backdropClass: "INDOXMESSAGES",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(OcsIndoxMsgsComponent, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.OCSmsgs = msgs
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
