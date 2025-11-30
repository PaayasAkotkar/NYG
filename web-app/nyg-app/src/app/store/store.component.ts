import { EventEmitter, Component, Input, OnInit, Output } from '@angular/core';

import { OcsCostSheetComponent } from "../UI/upgrade-menu/ocs-cost-sheet/ocs-cost-sheet.component";
import { OcsBtn2Component } from "../setup/ocs-btn-2/ocs-btn-2.component";

import { ButtonTagService } from '../UI/button-tag/button-tag.service';
import mqtt from 'mqtt';
import { BehaviorSubject } from 'rxjs';
import { SessionService } from '../storage/session/session.service';
import { NYGmqttService } from '../mqtt-service/nygmqtt.service';
import { Coin, PurchaseReciept, Spur } from '../mqtt-service/vouchers/mqttVouchers';
export interface StoreAlert {
  coin: string,
  spur: string,
  isStoreClicked: boolean
}
export interface IStore {
  Title: string
  Price: string
}


@Component({
  selector: 'app-store',
  imports: [OcsCostSheetComponent, OcsBtn2Component],
  templateUrl: './store.component.html',
  styleUrl: './store.component.scss'
})
export class StoreComponent implements OnInit {
  constructor(private nygmt: NYGmqttService, private se: SessionService, private cp: ButtonTagService, private sp: ButtonTagService) { }
  @Input({ required: true }) OCSspurs: IStore[]
  @Input({ required: true }) OCScoins: IStore[]
  @Output() OCSspurPurchase: EventEmitter<boolean> = new EventEmitter<boolean>()
  @Output() OCScoinPurchase: EventEmitter<boolean> = new EventEmitter<boolean>()

  ID: string = ""
  ngOnInit(): void {
    this.cp.CoinPurchase().subscribe({
      next: (token) => {
        if (token.purchase) {
          var re = new Date()
          var date = `${re.getDate()}/${re.getMonth() + 1}/${re.getFullYear()}`
          var meridiem = re.getHours() <= 12 ? 'AM' : 'PM'
          var time = `${re.getHours()}:${re.getMinutes()}:${re.getSeconds()} ${meridiem}`
          var r: PurchaseReciept = { purchaseMade: this.lockPrice, time: time, date: date }
          var struct: Coin = { coin: this.lockCoin, id: this.ID, reciept: r }
          this.nygmt.UpdateCoin(struct)
          this.OCScoinPurchase.emit(true)
        }
      }
    })
    this.sp.SpurPurchase().subscribe({
      next: (token) => {
        if (token.purchase) {
          var i = this.se.getItem("session")
          if (i) {
            this.ID = i
            var re = new Date()
            var date = `${re.getDate()}/${re.getMonth() + 1}/${re.getFullYear()}`
            var meridiem = re.getHours() <= 12 ? 'AM' : 'PM'
            var time = `${re.getHours()}:${re.getMinutes()}:${re.getSeconds()} ${meridiem}`
            var r: PurchaseReciept = { purchaseMade: this.lockPrice, time: time, date: date }
            var struct: Spur = { spur: this.lockSpur, id: this.ID, reciept: r }
            this.nygmt.UpdateSpur(struct)
            this.OCSspurPurchase.emit(true)
          }
        }
      }
    })

  }

  Page: string = "SPURS"

  toSpurs() {
    this.Page = "SPURS"
  }
  toCoins() {
    this.Page = "COINS"
  }
  toOffers() {
    this.Page = "OFFERS"
  }
  toThemes() {
    this.Page = "THEMES"
  }

  CoinPurchase(price: string) {
    var i = this.se.getItem("session")
    if (i) {
      this.ID = i
      var clean = parseInt(price.replace(/C/, ""))

      this.lockCoin = price
      var p = (clean / 2) * 10 * 2.5
      this.lockPrice = p.toString()
      this.cp.Setup(p.toString(), "PURCHASE", "CANCEL", true, false)
    }
  }

  lockSpur: string = ""
  lockCoin: string = ""
  lockPrice: string = ""

  SpurPurchase(price: string) {
    var i = this.se.getItem("session")
    if (i) {
      this.ID = i

      var clean = parseInt(price.replace(/S/, ""))
      var p = (clean / 2) * 10 * 2.5

      this.lockSpur = price
      this.lockPrice = p.toString()
      this.sp.Setup(p.toString(), "PURCHASE", "CANCEL", false, false)
    }
  }
}
