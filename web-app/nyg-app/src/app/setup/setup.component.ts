import { Component, ElementRef, Input, OnInit, signal, ViewChild, WritableSignal } from '@angular/core';
import { FormBuilder, ReactiveFormsModule } from '@angular/forms';
import { OcsDetComponent } from "../profile/ocs-det/ocs-det.component";
import { OCSbuttonV1 } from "../UI/upgrade-menu/upgrade-part1/ocs-button.component";
import { ViewPattern3Component } from "./view-pattern3/view-pattern3.component";
import { SetScoreboardService } from './view-pattern/set-scoreboard.service';
import { SetThemeService } from './view-theme/theme.service';
import { OcsBtn2Component } from "./ocs-btn-2/ocs-btn-2.component";
import { OcsInputServiceV2 } from './ocs-input-2/ocs-input-service-v2.service';

import { IStore, StoreAlert, StoreComponent } from '../store/store.component';
import { ProfileComponent } from '../profile/profile.component';
import { OcsSaveImgComponent } from './ocs-save-img/ocs-save-img.component';
import { PocketCredits, UpgradeAlert, UpgradeMenuComponent, UpgradePattern } from '../UI/upgrade-menu/upgrade-menu.component';
import { ActivatedRoute, Route } from '@angular/router';
import { ProfileService } from '../profile-service/profile.service';
import { _F } from '../UI/ocs-image/image.service';
import { NYGqService } from '../graphql/nygq.service';
import { gql } from 'apollo-angular';
import { SessionService } from '../storage/session/session.service';
import { HomePageLabelComponent } from '../UI/home-page-label/home-page-label.component';
import { HomeProfileComponent, IHomePageProfile } from '../UI/home-profile/home-profile.component';
import { HomeProfileService } from '../UI/home-profile/home-profile.service';
import { HAND } from '../helpers/hand';
import { concatMap, filter, isEmpty, Subscription, tap, timer } from 'rxjs';
import { NYGmqttService } from '../mqtt-service/nygmqtt.service';
import { Upgrade } from '../mqtt-service/vouchers/mqttVouchers';

import { ViewPattern4Component } from "./view-pattern-4/view-pattern-4.component";
import { OcsManualService } from '../UI/ocs-manual/ocs-manual.service';
export interface IIMG {
  id: string
  imageURL: string
  imageName: string
  imageType: string
}
export interface IProfile {
  Name: WritableSignal<string>
  Nickname: WritableSignal<string>
  ProfilePic: WritableSignal<string>
  GamesPlayed: WritableSignal<number>
  Points: WritableSignal<number>
  Rating: WritableSignal<number>
  Tier: WritableSignal<string>
  Record: WritableSignal<string>
  Coin: string
  Spur: string
}
@Component({
  selector: 'app-setup',
  imports: [HomePageLabelComponent, UpgradeMenuComponent, OcsSaveImgComponent, ProfileComponent, StoreComponent, ReactiveFormsModule, OcsDetComponent, OCSbuttonV1, ViewPattern3Component, OcsBtn2Component, ViewPattern4Component],
  templateUrl: './setup.component.html',
  styleUrls: ['./touchup/setup.component.scss', './touchup/animations.scss']
})
export class SetupComponent implements OnInit {
  constructor(private vv: OcsManualService, private prof: HomeProfileService, private se: SessionService, private q: NYGqService, private ps: ProfileService, private router: ActivatedRoute, private nn: OcsInputServiceV2, private FB: FormBuilder, private sco: SetScoreboardService, private gu: SetThemeService, private ins: OcsInputServiceV2) { }

  NYGProfile: IProfile = {
    GamesPlayed: signal(0), Points: signal(0),
    ProfilePic: signal(""), Tier: signal(""), Record: signal(""),
    Coin: "", Spur: "", Name: signal(""), Nickname: signal(""), Rating: signal(0)
  }
  NYGcredits: PocketCredits = { coin: "1C", spur: "1S" }
  OCSpicturePath: WritableSignal<string> = signal("")
  OCSgamesPlayed: WritableSignal<number> = signal(121)
  OCSrating: WritableSignal<number> = signal(400)
  OCSpoints: WritableSignal<number> = signal(12)
  // OCSrecords: WritableSignal<string> = signal('.2s')
  OCStier: WritableSignal<string> = signal('NOVICE')
  coinDoIt: boolean = true
  spurDoIt: boolean = true

  sm = ['YOUTUBE', 'X', 'INSTAGRAM', 'FACEBOOK', 'REDDIT']

  PageNavigation_: string = "profile"

  ThemeName: WritableSignal<string> = signal('')
  DisplayName: WritableSignal<string> = signal('KING')
  DisplayNickName: WritableSignal<string> = signal('KING')
  ID: string = ""
  OCSupgradeList: UpgradePattern[] = []
  OCSPlayer2v2Profile: IHomePageProfile = { Name: "", Rating: 0, Record: "", Spur: "", Coin: "", GamesPlayed: 0, Tier: "", ProfilePic: signal(""), Points: 0 }
  OCSPlayer1v1Profile: IHomePageProfile = { Name: "", Rating: 0, Record: "", Spur: "", Coin: "", GamesPlayed: 0, Tier: "", ProfilePic: signal(""), Points: 0 }
  OCSPlayerClashProfile: IHomePageProfile = { Name: "", Rating: 0, Record: "", Spur: "", Coin: "", GamesPlayed: 0, Tier: "", ProfilePic: signal(""), Points: 0 }

  view2v2() {
    this.prof.Setup(this.OCSPlayer2v2Profile)
  }

  view1v1() {
    this.prof.Setup(this.OCSPlayer1v1Profile)
  }

  viewClash() {
    this.prof.Setup(this.OCSPlayerClashProfile)
  }
  viewManual() {
    this.vv.Setup()
  }
  PageNavigate(page: string) {
    this.PageNavigation_ = page
  }
  coinPurcahse(made: boolean) {

    if (made)
      timer(1).pipe(
        concatMap(() => timer(100).pipe(tap(() => {
          this.coinDoIt = false
        }))),
        concatMap(() => timer(1400).pipe(tap(() => {
          this.coinDoIt = true
        }))),
      ).subscribe()

  }
  spurPurchase(made: boolean) {
    if (made)
      timer(1).pipe(
        concatMap(() => timer(100).pipe(tap(() => {
          this.spurDoIt = false
        }))),
        concatMap(() => timer(1400).pipe(tap(() => {
          this.spurDoIt = true
        }))),
      ).subscribe()
  }
  // CreateUpgradePattern this was a old way of doing it now it is shift to map
  CreateUpgradePattern(p: {
    covert: number, nexus: number, tag: number, draw: number,
    rewind: number, freeze: number, bet: number
  }, l: UpgradePattern[]) {
    let keys = Object.keys(p).filter(k => k != "__typename") as (keyof typeof p)[]

    for (let key of keys) {
      if (p[key] as number == 100) {
        l.push({ power: key, spur: p[key] as number, counter: 0, display: "", allow: false, level: 0, })
      } else {
        l.push({ power: key, spur: p[key] as number, counter: 0, display: "", allow: true, level: 0, })
      }
    }
  }

  OCScoin: IStore[] = [{ Title: "COIN", Price: "400C" }]
  OCSspur: IStore[] = [{ Title: "SPUR", Price: "400S" }]
  wasAUpgradeClick: boolean = false


  ngOnInit(): void {

    var i = this.se.getItem("session")
    if (i) {
      this.ID = i

      const lo = this.q._Login(this.ID)
      lo.subscribe({
        next: (token) => {
          if (token.data?.login && token.data.login.updatedDisplay && token.data.login.updatedPowerupUpgrades && token.data.login.updatedPlayerCredits && token.data.login.updatedStore) {
            // shallow copy for mutation
            var x = token.data.login.updatedPowerupUpgrades.upgrades
            var a = x.map(item => ({ ...item }))
            this.OCSupgradeList = a.filter(e => e.power != ' ')
            this.OCSupgradeList.sort((e, i) => e.power.length - i.power.length)

            this.OCScoin = token.data.login.updatedStore.coin
            this.OCSspur = token.data.login.updatedStore.spur
            var us = token.data.login.updatedPlayerCredits
            this.NYGProfile.Coin = us.coins
            this.NYGProfile.Spur = us.spurs
            var ud = token.data.login

            this.NYGProfile.ProfilePic = signal(ud.updatedDisplay.image.imgURL)
            this.NYGProfile.Name = signal(ud.updatedDisplay.name)
            this.NYGProfile.Nickname = signal(ud.updatedDisplay.nickname)

            this.OCSPlayer1v1Profile.Name = ud.updatedDisplay.name
            this.OCSPlayer2v2Profile.Name = ud.updatedDisplay.name
            this.OCSPlayerClashProfile.Name = ud.updatedDisplay.name

            this.OCSPlayer2v2Profile.GamesPlayed = ud.updatedDisplay.twovtwoProfile.gamesPlayed
            this.OCSPlayer1v1Profile.GamesPlayed = ud.updatedDisplay.onevoneProfile.gamesPlayed
            this.OCSPlayerClashProfile.GamesPlayed = ud.updatedDisplay.clashProfile.gamesPlayed

            this.OCSPlayer2v2Profile.Rating = ud.updatedDisplay.twovtwoProfile.rating
            this.OCSPlayer1v1Profile.Rating = ud.updatedDisplay.onevoneProfile.rating
            this.OCSPlayerClashProfile.Rating = ud.updatedDisplay.clashProfile.rating

            this.OCSPlayer2v2Profile.Points = ud.updatedDisplay.twovtwoProfile.points
            this.OCSPlayer1v1Profile.Points = ud.updatedDisplay.onevoneProfile.points
            this.OCSPlayerClashProfile.Points = ud.updatedDisplay.clashProfile.points

            this.OCSPlayer2v2Profile.Tier = ud.updatedDisplay.twovtwoProfile.tier
            this.OCSPlayer1v1Profile.Tier = ud.updatedDisplay.onevoneProfile.tier
            this.OCSPlayerClashProfile.Tier = ud.updatedDisplay.clashProfile.tier

            this.OCSpicturePath.set(ud.updatedDisplay.image.imgURL)
            this.OCSPlayer1v1Profile.ProfilePic.set(ud.updatedDisplay.image.imgURL)
            this.OCSPlayer2v2Profile.ProfilePic.set(ud.updatedDisplay.image.imgURL)
            this.OCSPlayerClashProfile.ProfilePic.set(ud.updatedDisplay.image.imgURL)

            this.DisplayName.set(ud.updatedDisplay.name)
            this.DisplayNickName.set(ud.updatedDisplay.nickname)
            timer(1).pipe(
              concatMap(() => timer(100).pipe(tap(() => {
                this.spurDoIt = false
              }))),
              concatMap(() => timer(1400).pipe(tap(() => {
                this.spurDoIt = true
              }))),
            ).subscribe()
            timer(1).pipe(
              concatMap(() => timer(100).pipe(tap(() => {
                this.coinDoIt = false
              }))),
              concatMap(() => timer(1400).pipe(tap(() => {
                this.coinDoIt = true
              }))),
            ).subscribe()
            this.NYGcredits.coin = this.NYGProfile.Coin
            this.NYGcredits.spur = this.NYGProfile.Spur
          }

        },
        error: (err) => console.error("error: ", err)
      })
      const wupuq = this.q.sub_UpdatedPowerupUpdgrades(this.ID)
      wupuq.subscribe({
        next: (token) => {
          if (token.data?.updatedPowerupUpgrades) {
            // shallow copy for mutation
            var x = token.data.updatedPowerupUpgrades.upgrades
            var a = x.map(item => ({ ...item }))
            this.OCSupgradeList = a.filter(e => e.power != ' ')
            this.OCSupgradeList.sort((e, i) => e.power.length - i.power.length)
          }
        }
      })

      const wusq = this.q.sub_UpdatedStore(this.ID)
      wusq.subscribe({
        next: (token) => {
          if (token.data?.updatedStore) {
            this.OCScoin = token.data.updatedStore.coin
            this.OCSspur = token.data.updatedStore.spur
          }
        }
      })

      const wupcq = this.q.sub_UpdatedPlayerCredits(this.ID)
      wupcq.subscribe({
        next: (token) => {
          if (token.data?.updatedPlayerCredits) {
            var us = token.data.updatedPlayerCredits
            this.NYGProfile.Coin = us.coins
            this.NYGProfile.Spur = us.spurs
            this.NYGcredits.coin = this.NYGProfile.Coin
            this.NYGcredits.spur = this.NYGProfile.Spur
          }
        }
      })

      const wq = this.q.sub_UpdatedDisplay(this.ID)

      wq.subscribe({
        next: (token) => {
          if (token.data?.updatedDisplay) {
            var ud = token.data.updatedDisplay
            this.NYGProfile.ProfilePic = signal(ud.image)
            this.NYGProfile.Name = signal(ud.name)
            this.NYGProfile.Nickname = signal(ud.nickname)

            this.OCSPlayer1v1Profile.Name = token.data.updatedDisplay.name
            this.OCSPlayer2v2Profile.Name = token.data.updatedDisplay.name
            this.OCSPlayerClashProfile.Name = token.data.updatedDisplay.name

            this.OCSPlayer2v2Profile.GamesPlayed = token.data.updatedDisplay.twovtwoProfile.gamesPlayed
            this.OCSPlayer1v1Profile.GamesPlayed = token.data.updatedDisplay.onevoneProfile.gamesPlayed
            this.OCSPlayerClashProfile.GamesPlayed = token.data.updatedDisplay.clashProfile.gamesPlayed

            this.OCSPlayer2v2Profile.Rating = token.data.updatedDisplay.twovtwoProfile.rating
            this.OCSPlayer1v1Profile.Rating = token.data.updatedDisplay.onevoneProfile.rating
            this.OCSPlayerClashProfile.Rating = token.data.updatedDisplay.clashProfile.rating

            this.OCSPlayer2v2Profile.Points = token.data.updatedDisplay.twovtwoProfile.points
            this.OCSPlayer1v1Profile.Points = token.data.updatedDisplay.onevoneProfile.points
            this.OCSPlayerClashProfile.Points = token.data.updatedDisplay.clashProfile.points

            this.OCSPlayer2v2Profile.Tier = token.data.updatedDisplay.twovtwoProfile.tier
            this.OCSPlayer1v1Profile.Tier = token.data.updatedDisplay.onevoneProfile.tier
            this.OCSPlayerClashProfile.Tier = token.data.updatedDisplay.clashProfile.tier

            this.OCSpicturePath.set(token.data.updatedDisplay.image)
            this.OCSPlayer1v1Profile.ProfilePic.set(token.data.updatedDisplay.image)
            this.OCSPlayer2v2Profile.ProfilePic.set(token.data.updatedDisplay.image)
            this.OCSPlayerClashProfile.ProfilePic.set(token.data.updatedDisplay.image)

            this.DisplayName.set(token.data.updatedDisplay.name)
            this.DisplayNickName.set(token.data.updatedDisplay.nickname)

          }
        }

      })
    }

    this.router.paramMap.subscribe(
      {
        next: (token) => {
          var j = token.get('page')
          if (j) {
            this.PageNavigation_ = j; // The '+' converts the string to a number
          }
        }
      }
    )

    this.nn.GetToken().subscribe({
      next: (token) => {
        if (token != "") {
          this.DisplayNickName.set(token)
        }
      }
    })
  }
  sub!: Subscription
  setTheme() {
    this.gu.Setup()
  }

  setName() {
    this.ins.Setup("NAME")
  }

  setNickName() {
    this.nn.Setup("NICKNAME")
  }

  setPicture() {
    this.ins.Setup("PICTURE",)
  }

  ImageFile_(token: _F) {
    this.imgFile = token.File
    this.imgURL = token.PostBase64
    this.imgType = token.FileType
    this.imgName = token.FileName
  }
  imgName: string = ""
  imgFile: File = new File([], "")
  imgURL: string
  imgType: string = ""

  Save() {
    const append = new FormData()
    append.append("img", this.imgFile, this.imgName)
    append.append("name", this.DisplayName())
    this.ps.Register(append).subscribe()
  }

  changeProfileSettings = this.FB.group(
    {
      ChangeNameToken: [''],
      ChangeDPToken: [''],
      ChangeScoreboardToken: [false],
      ChangeSoundToken: [false],
    }
  )
}
