import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { Injectable } from '@angular/core';
import { OcsInput2Component } from './ocs-input-2.component';
import { FormControl, FormGroup } from '@angular/forms';
import { BehaviorSubject, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class OcsInputServiceV2 {
  private ____OR: OverlayRef | null = null
  private recieveToken = new BehaviorSubject<string>('')
  constructor(private overaly: Overlay) { }
  Setup(OCSheader: string) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerVertically().centerHorizontally(),
      hasBackdrop: true,
      panelClass: "VIEW",
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(OcsInput2Component, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.NYGheader = OCSheader
    this.____OR.backdropClick().subscribe(() => {
      this.closeSetup()
    })
    __ref.instance.NYGEVENT.subscribe((token) => {
      this.recieveToken.next(token)

    })
    return this.____OR
  }

  GetToken(): Observable<string> {
    return this.recieveToken.asObservable()
  }

  closeSetup() {
    if (this.____OR) {
      this.____OR.dispose()
      this.____OR = null
    }
  }
}
