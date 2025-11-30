import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { Injectable, OnDestroy, WritableSignal } from '@angular/core';
import { GuessSheetComponent } from '../guess-sheet/guess-sheet.component';
import { ComponentPortal } from '@angular/cdk/portal';
import { Subject, takeUntil } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ViewGuessSheetService implements OnDestroy {

  private ____OR: OverlayRef | null
  constructor(private overaly: Overlay) { }
  Setup(guessSheetValue: WritableSignal<string>) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().top().left(),
      hasBackdrop: false,
      panelClass: "SHEETGUESS",
      backdropClass: "BackDSHEET",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(GuessSheetComponent, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.GuessSheetValue = guessSheetValue
    this.____OR.backdropClick().pipe(takeUntil(this.stop)).subscribe(() => {
      this.closeSetup()
    })
    return this.____OR
  }
  private stop: Subject<void> = new Subject<void>()
  ngOnDestroy(): void {
    this.closeSetup()
  }

  closeSetup() {
    if (this.____OR) {
      this.____OR.dispose()
      this.____OR = null
    }
  }
}
