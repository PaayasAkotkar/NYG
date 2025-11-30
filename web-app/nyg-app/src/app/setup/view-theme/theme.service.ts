import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { Injectable } from '@angular/core';
import { ViewThemeComponent } from './theme.component';
import { BehaviorSubject, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SetThemeService {

  private ____OR: OverlayRef | null = null
  private recieveToken = new BehaviorSubject<string>('')
  constructor(private overaly: Overlay) { }
  Setup() {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerVertically().centerHorizontally(),
      hasBackdrop: true,
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(ViewThemeComponent, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.OCSEVENT.subscribe((token) => {

      this.recieveToken.next(token)
    })
    this.____OR.backdropClick().subscribe(() => {
      this.closeSetup()
    })
    return this.____OR
  }

  GetTheme(): Observable<string> {
    return this.recieveToken.asObservable()
  }
  closeSetup() {
    if (this.____OR) {
      this.____OR.dispose()
      this.____OR = null
    }
  }
}
