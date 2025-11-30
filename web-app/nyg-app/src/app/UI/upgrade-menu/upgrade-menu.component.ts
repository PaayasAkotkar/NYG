import { EventEmitter, ChangeDetectorRef, Component, Input, OnInit, signal, Output, OnChanges, SimpleChanges, NgZone, } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { OCSbuttonV1, } from "./upgrade-part1/ocs-button.component";
import { PowerUpgradeComponent } from "./power-upgrade/power-upgrade.component";
import { OcsCounterComponent } from "../ocs-counter/ocs-counter.component";
import { OcsBtn2Component } from "../../setup/ocs-btn-2/ocs-btn-2.component";
import { ButtonTagService } from '../button-tag/button-tag.service';
import { BehaviorSubject, timer } from 'rxjs';
import { AlertService } from '../alert/alert.service';
import { NYGmqttService } from '../../mqtt-service/nygmqtt.service';
import { Upgrade } from '../../mqtt-service/vouchers/mqttVouchers';
import { SessionService } from '../../storage/session/session.service';
export interface UpgradeAlert {
  upgradeClick: boolean
  spur: string
  coin: string
  valid: boolean
}
export interface UpgradePattern {
  power: string
  spur: number
  counter: number
  level: number
  allow: boolean
  display: string
}
export interface PocketCredits {
  spur: string
  coin: string
}

@Component({
  selector: 'app-upgrade-menu',
  imports: [ReactiveFormsModule, OCSbuttonV1, PowerUpgradeComponent, OcsCounterComponent, OcsBtn2Component],
  templateUrl: './upgrade-menu.component.html',
  styleUrl: './upgrade-menu.component.scss'
})
export class UpgradeMenuComponent implements OnInit, OnChanges {
  upgradePowersForm: FormGroup = this.FB.group(
    {
      // will be send in the number
      BetToken: [],
      CovertToken: [],
      RewindToken: [],
      DrawToken: [],
      TagToken: [],
      NexusToken: [],
      FreezeToken: [],
    }
  )
  nav_ = 1
  Pages: number[] = []

  DonateSpur: number = 25
  CurrentSpur: number = 25
  Count = signal<number>(5)
  Eye = new BehaviorSubject<number>(0)
  price = 9
  @Input({ required: true }) OCSlist: UpgradePattern[]
  @Input({ required: true }) NYGcurrentCredits: PocketCredits

  frontRowPowers: UpgradePattern[] = []
  backRowPowers: UpgradePattern[] = []
  lockPower: string = ""

  ngOnChanges(changes: SimpleChanges): void {
    if (changes['OCSlist']) {
      this.OCSlist = this.OCSlist.filter(e => e.power != "")
      this.Arrange()
      this.Pages = [1, 2]
    }
  }
  ID: string = ""
  ngOnInit(): void {
    this.OCSlist.sort((e, i) => e.power.length - i.power.length)

    this.Arrange()

    this.pur.UpgradePurchase().subscribe({
      next: (token) => {
        if (token.made) {
          var currentCurrency = parseInt(this.NYGcurrentCredits.coin.replace("C", ""))
          var currentPurchase = parseInt(this.lockPrice.replaceAll("C", ""))
          var currentSpur = parseInt(this.NYGcurrentCredits.coin)
          var addSpur = this.Count()
          const succeed = currentSpur > addSpur && currentCurrency > currentPurchase

          if (this.lockPower != "") {
            if (succeed) {

              var i = this.se.getItem('session')
              if (i) {
                this.ID = i
                var newLevel = this.Count()
                var struct: Upgrade = { power: this.lockPower.toLowerCase(), id: this.ID, newLevel: newLevel, spur: this.price.toString() + "S", coin: currentPurchase.toString() + "C", }
                // proceed
                this.m.Setup("purchase successfuly", "CREDITS")
                this.mq.UpgradePower(struct)
                timer(400).subscribe(() => {
                  this.m.closeSetup()
                })
                token.made = false
              }
            } else if (currentCurrency < currentPurchase) {
              var need = currentPurchase - currentCurrency
              this.m.Setup("need more " + need.toString(), "coins credits")
              timer(400).subscribe(() => {
                this.m.closeSetup()
              })
              token.made = false
            } else if (addSpur < currentSpur) {
              var need = addSpur - currentSpur
              this.m.Setup("need more " + need.toString(), "spur credits")
              timer(400).subscribe(() => {
                this.m.closeSetup()
              })
              token.made = false
            }
          }
        }
      }
    })

    this.Eye.subscribe({
      next: (token) => {
        this.price = token
      }
    })
    // this.OCSlist = this.OCSlist.sort((e, a) => a.spur - e.spur)

    // this.evenPowers = this.OCSlist.filter((_, len) => len % 2 == 0)
    // this.oddPowers = this.OCSlist.filter((_, len) => len % 2 != 0)
  }

  coinPerc(spurs: number): number {
    var q = (spurs - 1.1) * 99.1 / 9.8
    var j = q * spurs
    return Math.floor(j)
  }

  Nav(e: number): void {
    this.nav_ = parseInt(String(e))
  }
  Add(index: number) {
    this.run.run(() => {
      if (this.frontRowPowers[index].counter < 100 && this.frontRowPowers[index].spur != 100) {
        this.lockPower = this.frontRowPowers[index].power
        this.Count.set(this.frontRowPowers[index].spur)
        this.Count.update(v => v += 5)
        this.frontRowPowers[index].counter = this.Count()
        this.frontRowPowers[index].spur = this.Count()
        this.frontRowPowers[index].display = `${this.frontRowPowers[index].counter} S`
        this.Eye.next(this.Count())
        this.cdr.detectChanges()
      }
    })
  }

  Subtract(index: number) {
    var valid: boolean = false
    for (let i of this.OCSlist) {
      if (i.power == this.frontRowPowers[index].power) {
        if (this.Count() > i.level) {
          valid = true
        }

      }
    }
    if (valid) {
      this.run.run(() => {
        if (this.frontRowPowers[index].counter > 0 && this.frontRowPowers[index].spur <= 100) {
          this.lockPower = this.frontRowPowers[index].power
          this.Count.set(this.frontRowPowers[index].spur)
          this.Count.update(v => v -= 5)
          this.frontRowPowers[index].counter = this.Count()
          this.frontRowPowers[index].spur = this.Count()
          this.frontRowPowers[index].display = `${this.frontRowPowers[index].counter} S`
          this.Eye.next(this.Count())
          this.cdr.detectChanges()
        }
      })

    }

  }

  // Arrange arranges the list into front row and backrow
  Arrange(): void {
    this.frontRowPowers = this.OCSlist.filter(e => e.allow)
    this.backRowPowers = this.OCSlist.filter(e => !e.allow)

    if (this.frontRowPowers.length == 2) {
      const max = 2
      const min = max
      this.frontRowPowers.push(...this.backRowPowers.slice(0, max))
      this.backRowPowers = this.backRowPowers.slice(min, this.backRowPowers.length)
    } else if (this.frontRowPowers.length == 1) {
      const max = 3
      const min = max
      this.frontRowPowers.push(...this.backRowPowers.slice(0, 3))
      this.backRowPowers = this.backRowPowers.slice(min, this.backRowPowers.length)
    } else if (this.frontRowPowers.length == 0) {
      const min = 0
      const max = this.OCSlist.length - 1
      const mid = min + ((max - min) / 2)
      this.frontRowPowers = this.OCSlist.slice(mid, this.OCSlist.length)
      this.backRowPowers = this.OCSlist.slice(min, mid)
    }
  }
  lockPrice: string = ""
  Purchase() {
    var price = this.coinPerc(this.Eye.value).toString()
    this.lockPrice = price + "C"
    this.pur.Setup(price, "UPGRADE", "CANCEL", false, true)
  }

  constructor(private se: SessionService, private mq: NYGmqttService, private run: NgZone, private m: AlertService, private FB: FormBuilder, private cdr: ChangeDetectorRef, private pur: ButtonTagService) { }
}
