import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { Injectable } from '@angular/core';
import { ButtonTagComponent, PurhcaseMade } from './button-tag.component';
import { ComponentPortal } from '@angular/cdk/portal';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ButtonTagService {

  private ____OR: OverlayRef | null = null
  private coinPurchase: BehaviorSubject<PurhcaseMade> = new BehaviorSubject<PurhcaseMade>({ made: false, purchase: "" })
  private spurPurchase: BehaviorSubject<PurhcaseMade> = new BehaviorSubject<PurhcaseMade>({ made: false, purchase: "" })
  private upgradePurchase: BehaviorSubject<PurhcaseMade> = new BehaviorSubject<PurhcaseMade>({ made: false, purchase: "" })

  constructor(private overaly: Overlay) { }
  Setup(Title: string, SubTitle: string, SubTitle2: string, isCoin: boolean, isUpgrade: boolean) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerVertically().centerHorizontally(),
      hasBackdrop: true,
      backdropClass: "DBTBack",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(ButtonTagComponent, null,)
    const __ref = this.____OR.attach(portal)

    __ref.instance.NYGtitle = Title
    __ref.instance.NYGsubTitle = SubTitle
    __ref.instance.NYGsubTitle2 = SubTitle2
    __ref.instance.NYGcoin = isCoin
    __ref.instance.NYGupgrade = isUpgrade

    if (isCoin) {
      __ref.instance.NYGCoinPurchase.asObservable().subscribe({
        next: (token) => {
          this.coinPurchase.next(token)
        }
      })
    }

    if (!isCoin) {
      __ref.instance.NYGSpurPurchase.asObservable().subscribe({
        next: (token) => {
          this.spurPurchase.next(token)
        }
      })
    }
    if (this.upgradePurchase) {
      __ref.instance.NYGUpgradePurchase.asObservable().subscribe({
        next: (token) => {
          this.upgradePurchase.next(token)
        }
      })
    }
    this.____OR.backdropClick().subscribe(() => {
      this.closeSetup()
    })

    return this.____OR
  }
  CoinPurchase() {
    return this.coinPurchase.asObservable()
  }
  SpurPurchase() {
    return this.spurPurchase.asObservable()
  }
  UpgradePurchase() {
    return this.upgradePurchase.asObservable()
  }
  closeSetup() {
    if (this.____OR) {
      this.____OR.dispose()
      this.____OR = null
    }
  }
}
