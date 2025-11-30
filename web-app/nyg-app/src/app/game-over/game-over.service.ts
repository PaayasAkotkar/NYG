import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { Injectable } from '@angular/core';
import { GameOverComponent } from './game-over.component';
import { ComponentPortal } from '@angular/cdk/portal';
import { GameResultComponent } from '../UI/game-result/game-result.component';
import { IGameOver } from '../lobby/resources/lobbyResource';

@Injectable({
  providedIn: 'root'
})
export class GameOverService {

  private ____OR: OverlayRef | null = null
  constructor(private overaly: Overlay) { }
  Setup(gameOverInfo: IGameOver) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerVertically().centerHorizontally(),
      hasBackdrop: true,
      backdropClass: "RATING",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(GameResultComponent, null,)
    const __ref = this.____OR.attach(portal)
    __ref.instance.OCSinfo = gameOverInfo
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
