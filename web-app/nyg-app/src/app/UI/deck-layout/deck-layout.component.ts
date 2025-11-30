import { Component, OnInit, signal, WritableSignal } from '@angular/core';
import { PowerRewindComponent } from "../power-up/power-rewind/power-rewind.component";
import { PowerSetttingsPatternComponent } from "../power-setttings-pattern/power-setttings-pattern.component";
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { PowerUpPatternComponent } from "../power-up/power-up-pattern/power-up-pattern.component";
import { GoBooleanMap, HAND } from '../../helpers/hand';
import { OCSFieldFormat } from '../../helpers/patterns';
import { ScrollBehvaiorDirective } from './scroll-behvaior.directive';


import { OcsBtnsComponent } from "../ocs-btns/ocs-btns.component";
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { OcsXxBtnComponent } from "../ocs-xx-btn/ocs-xx-btn.component";
import { UpdateLivelyService } from '../../update-lively/update-lively.service';
import { IDeck } from '../../patches/patch';
import { LoginService } from '../../login-service/login.service';
import { timer } from 'rxjs';
import { RoomService } from '../../roomService/room.service';
import { NYGqService } from '../../graphql/nygq.service';
import { SessionService } from '../../storage/session/session.service';
import { UpgradePattern } from '../upgrade-menu/upgrade-menu.component';

import { PowerUpgradeComponent } from "../upgrade-menu/power-upgrade/power-upgrade.component";

@Component({
  selector: 'app-deck-layout',
  imports: [ScrollBehvaiorDirective, ReactiveFormsModule, OcsBtnsComponent, CommonModule, OcsXxBtnComponent, PowerUpgradeComponent],
  templateUrl: './deck-layout.component.html',
  styleUrls: ['./touchup/deck.scss', './touchup/color.scss', './touchup/animations.scss']
})
export class DeckLayoutComponent implements OnInit {
  _f: FormGroup
  _f2: FormGroup
  id: string = ""
  OCSN = 12
  Powers: string[] = []
  CurrentPowers: string[] = []
  _switch: boolean = false
  Switching: WritableSignal<string> = signal("default")
  From: WritableSignal<string> = signal("default")
  DeckUpdate: string[] = []
  toPlay() {
    var k1 = Object.keys(this._f2.controls)
    for (var i of k1) {
      if (this._f2.get(i)?.value) {
        this.DeckUpdate.push(i)
      }
    }
    this.DeckUpdate = this.DeckUpdate.map(e => e.toLowerCase())
    this._r.navigateByUrl('/ClashLobby')
    this.Clash()
    timer(10).subscribe(() => {
      this._d.UpdateDeck(this.DeckUpdate[0], this.DeckUpdate[1], false)
    })
  }
  notation: Record<string, number> = {}
  NYGprofile: UpgradePattern[] = []
  recev: IDeck = { _first: "", _second: "", isDefault: false, id: "" }
  defaultDeck: IDeck = { _first: "", _second: "", isDefault: false, id: "" }

  ngOnInit(): void {
    var i = this.se.getItem('session')
    if (i) {
      this.id = i
      this._f = this.FB.group({})
      this._f2 = this.FB.group({})

      this.q._Login(this.id).subscribe({

        next: (token) => {
          if (token.data.login.updatedPowerupUpgrades) {
            var up = token.data.login.updatedPowerupUpgrades
            var u = token.data.login.currentDeck
            this.recev = u
            this.CurrentPowers.push(this.recev._first.toUpperCase())
            this.CurrentPowers.push(this.recev._second.toUpperCase())
            this.NYGprofile = up.upgrades
            this.defaultDeck = this.recev
            for (let i of this.NYGprofile) {
              this.Powers.push(i.power)
              this.notation[i.power] = i.level
            }
            this.Powers.forEach(item => {
              this._f.addControl<string>(item, this.FB.control(false))
              this._f2.addControl<string>(item, this.FB.control(false))

            })
            this.CurrentPowers.forEach(item => {
              this._f2.addControl<string>(item, this.FB.control(true))
              this._f.addControl<string>(item, this.FB.control(false))
            })
            this.Powers = this.Powers.filter(e => !this.CurrentPowers.includes(e))
          }

        }
      })
      this.q.sub_UpdatedPowerupUpdgrades(this.id).subscribe({
        error: (e) => { console.error(e) }
      })
    }
    this.q.sub_UpdatedDeck(this.id).subscribe({
      next: (token) => {
        if (token.data?.currentDeck) {
          this.defaultDeck = token.data.currentDeck
        }
      }
    })
  }

  _Switch(_for: string) {
    this._switch = !this._switch
    this.Switching.set(_for)
  }

  onUse(power: string) {
    var a = this.Powers.indexOf(power)
    var b = this.CurrentPowers.indexOf(this.Switching())
    var temp = this.Powers[a]
    this.Powers[a] = this.CurrentPowers[b]
    this.CurrentPowers[b] = temp
    this._switch = false
    // default it
    this._f2.get(this.Switching())?.setValue(false)
    this._f2.get(power)?.setValue(true)
    this.From.set(power)
  }

  Save() {
    this.DeckUpdate = [] // empty it
    // this will be send form
    this._switch = false
    var g = ""
    var h = ""
    var d = false

    if (this.Switching() != "default" && this.From() != "default") {

      var keys = Object.keys(this._f2.controls)
      var get_: string[] = []
      for (var k of keys) {
        if (this._f2.get(k)?.value) {
          get_.push(k)
          this.DeckUpdate.push(k)
        }
      }
    } else {

      d = true
      var keys = Object.keys(this._f2.controls)
      var get_: string[] = []

      for (var k of keys) {
        if (this._f2.get(k)?.value) {
          get_.push(k)
          this.DeckUpdate.push(k)
        }
      }
    }

    g = get_[0]
    h = get_[1]

    var c: IDeck = {
      _first: g.toLowerCase(),
      _second: h.toLowerCase(),
      id: this.id,
      isDefault: d,
    }

    if (this.defaultDeck != c) {
      // just for testing
      timer(10).subscribe(() => {
        this._d.UpdateDeck(c._first, c._second, c.isDefault)
      })
    }
  }

  powers: {
    covert: boolean, rewind: boolean, freeze: boolean,
    nexus: boolean, bet: boolean
  } = {
      covert: false, nexus: false, rewind: false, freeze: false, bet: false
    }

  Clash() {

    for (let i of this.DeckUpdate) {
      switch (i) {
        case 'nexus':
          this.powers.nexus = true
          break
        case 'covert':
          this.powers.covert = true
          break
        case 'bet':
          this.powers.bet = true
          break
        case 'freeze':
          this.powers.freeze = true
          break
        case 'rewind':
          this.powers.rewind = true
          break
      }
    }
    this.room.Clash(this.powers.covert, this.powers.nexus,
      this.powers.rewind, this.powers.freeze,
      this.powers.bet
    )
  }

  constructor(private se: SessionService, private q: NYGqService, private room: RoomService, private FB: FormBuilder, private _r: Router, private _d: UpdateLivelyService, private l: LoginService) { }

}