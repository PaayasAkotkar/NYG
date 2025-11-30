import { Component, Input, Output, EventEmitter } from '@angular/core';
import { OcsBtnsComponent } from "../ocs-btns/ocs-btns.component";
import { ButtonTagService } from './button-tag.service';
export interface PurhcaseMade {
  made: boolean
  purchase: string
}
@Component({
  selector: 'app-button-tag',
  imports: [OcsBtnsComponent],
  templateUrl: './button-tag.component.html',
  styleUrls: ['./touchup/animations.scss', './touchup/button-tag.scss', './touchup/color.scss'],

})
export class ButtonTagComponent {
  @Input() NYGtitle: string = "500R"
  @Input() NYGsubTitle: string = "PURCHASE"
  @Input() NYGsubTitle2: string = "CANCEL"
  @Input() NYGcoin: boolean = false
  @Input() NYGupgrade: boolean = false
  @Output() NYGListenSubTitle: EventEmitter<void> = new EventEmitter<void>()
  @Output() NYGListenSubTitle2: EventEmitter<void> = new EventEmitter<void>()
  @Output() NYGCoinPurchase: EventEmitter<PurhcaseMade> = new EventEmitter<PurhcaseMade>()
  @Output() NYGSpurPurchase: EventEmitter<PurhcaseMade> = new EventEmitter<PurhcaseMade>()
  @Output() NYGUpgradePurchase: EventEmitter<PurhcaseMade> = new EventEmitter<PurhcaseMade>()

  // note: this is created only for the store purchase
  Purchase() {
    this.NYGListenSubTitle.emit()

    var made = true
    if (this.NYGcoin) {
      this.NYGCoinPurchase.emit({ made: made, purchase: this.NYGtitle })
      made = false
    } else if (!this.NYGcoin) {
      this.NYGSpurPurchase.emit({ made: made, purchase: this.NYGtitle })
    }
    if (this.NYGupgrade) {
      this.NYGUpgradePurchase.emit({ made: made, purchase: this.NYGtitle })
    }
    this.can.closeSetup()
  }

  Cancel() {
    this.NYGListenSubTitle2.emit()
    this.can.closeSetup()
  }

  constructor(private can: ButtonTagService) { }
}
