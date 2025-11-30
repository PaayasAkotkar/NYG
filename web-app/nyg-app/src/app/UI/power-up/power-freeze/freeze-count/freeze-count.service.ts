import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { Injectable, signal, WritableSignal } from '@angular/core';
import { FreezeCountComponent } from './freeze-count.component';
import { Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class FreezeCountService {
  private ____OR: OverlayRef | null
  constructor(private overaly: Overlay,) { }
  Setup(freezeAccess: number, stopSignal: WritableSignal<boolean>, clash: WritableSignal<boolean>, RoomName: string) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerVertically(),
      hasBackdrop: false,
      panelClass: "FREEZECOUNT",
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(FreezeCountComponent, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.NYGStopSignal = stopSignal
    __ref.instance.NYGroomName = RoomName
    __ref.instance.NYGclash = clash
    __ref.instance.freezeFor = freezeAccess
    __ref.instance.Watcher.subscribe((token) => {
      this.current.next(token)

    })
    return this.____OR
  }
  current: Subject<boolean> = new Subject<boolean>()
  Listen() {
    return this.current.asObservable()
  }
  closeSetup() {
    if (this.____OR) {
      this.____OR.dispose()
      this.____OR = null
    }
  }
}
