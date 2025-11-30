import { ChangeDetectorRef, Component, Input, OnInit, signal, WritableSignal } from '@angular/core';
import { RoomService } from '../roomService/room.service';
import { concatMap, filter, tap, timer } from 'rxjs';
import { HomePageLabelComponent } from "../UI/home-page-label/home-page-label.component";
import { HomePageNavComponent } from "../UI/home-page-nav/home-page-nav.component";
import { IDailyEvents } from "../UI/daily-events/daily-events.component";
import { IHomePageProfile } from "../UI/home-profile/home-profile.component";
import { OcsImgComponent } from "../profile/ocs-img/ocs-img.component";
import { HomePageNav2Component } from "../UI/home-page-nav2/home-page-nav2.component";
import { DailyEventsService } from '../setup-daily-events/daily-events.service';
import { Router } from '@angular/router';
import { HomeProfileService } from '../UI/home-profile/home-profile.service';
import { LocifyComponent } from "../UI/locify/locify.component";
import { ClashComponent } from "../UI/clash/clash.component";
import { gql } from 'apollo-angular';
import { NYGqService } from '../graphql/nygq.service';
import { SessionService } from '../storage/session/session.service';
import { HAND } from '../helpers/hand';
import { UpdateLivelyService } from '../update-lively/update-lively.service';
import { Credits } from '../update-lively/resource/update-livelyResource';
import { HomePagePorfilesService } from '../UI/home-page-profiles/home-page-porfiles.service';
import { OcsIndoxBoxComponent } from "../UI/ocs-indox-box/ocs-indox-box.component";
import { IMessages, OcsIndoxMsgsComponent } from '../UI/ocs-indox-msgs/ocs-indox-msgs.component';
import { OcsIndoxMsgsService } from '../UI/ocs-indox-msgs/ocs-indox-msgs.service';
import { HttpClient } from '@angular/common/http';
@Component({
  selector: 'app-home-page',
  imports: [HomePageLabelComponent, HomePageNavComponent, OcsImgComponent, HomePageNav2Component, LocifyComponent, ClashComponent, OcsIndoxBoxComponent],
  templateUrl: './home-page.component.html',
  styleUrl: './home-page.component.scss'
})
export class HomePageComponent implements OnInit {

  constructor(private co: HttpClient, private indoxms: OcsIndoxMsgsService, private profs: HomePagePorfilesService, private cdr: ChangeDetectorRef, private login: UpdateLivelyService, private se: SessionService, private q: NYGqService, private prof: HomeProfileService, private _route: Router, private _room: RoomService, private _cdr: ChangeDetectorRef, private ev: DailyEventsService) { }

  @Input() NYGProfile: IHomePageProfile = {
    Name: "",
    Rating: 400, Points: 0, Record: "0",
    GamesPlayed: 0,
    Tier: "novice",
    ProfilePic: signal(""),
    Coin: "0C",
    Spur: "0S",
  }

  @Input() OCSImagePath: WritableSignal<string> = signal("")

  eventSeen: boolean = false
  storeSeen: boolean = false

  NYGclashEvents: IDailyEvents[] = []
  NYGlocifyEvents: IDailyEvents[] = []

  @Input({ required: false }) OCSprofilePicSrc: WritableSignal<string> = signal("")
  @Input() OCSindoxMsg: string = "new"
  @Input() isIndoxMsg: boolean = true

  updateCredits: Credits = { coins: "", boardTheme: "", guessTheme: "", spurs: "", id: "" }

  id: string
  coinDoIt: boolean = true
  spurDoIt: boolean = true
  eventDoIt: boolean = true
  OCSPlayerClashProfile: IHomePageProfile = { Name: "", Rating: 0, Record: "", Spur: "", Coin: "", GamesPlayed: 0, Tier: "", ProfilePic: signal(""), Points: 0 }
  OCSPlayer1v1Profile: IHomePageProfile = { Name: "", Rating: 0, Record: "", Spur: "", Coin: "", GamesPlayed: 0, Tier: "", ProfilePic: signal(""), Points: 0 }
  OCSPlayer2v2Profile: IHomePageProfile = { Name: "", Rating: 0, Record: "", Spur: "", Coin: "", GamesPlayed: 0, Tier: "", ProfilePic: signal(""), Points: 0 }
  OCSmesages: IMessages[] = []
  openIndoxBox() {
    this.indoxms.Setup(this.OCSmesages)
  }
  ngOnInit(): void {
    const url = 'http://localhost:1500/nyg-welcome'
    this.co.get<IMessages>(url).subscribe({
      next: (token) => {
        if (token) {
          this.OCSmesages.push(token)
        }
      }
    })
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
    this.ev.getClashRewards().pipe(filter(v => v.reward != "default")).subscribe({
      next: (token) => {
        this.ev.closeSetup()
        if (token.isCoin) {
          if (token.reward != "default") {
            timer(0).pipe(
              concatMap(() => timer(100).pipe(tap(() => {
                this.coinDoIt = false
                this.NYGProfile.Coin = HAND.Calculate(this.NYGProfile.Coin, Number.parseInt(token.reward), /C/)
                this.updateCredits.coins = this.NYGProfile.Coin
                this.login.UpdateCoins(this.updateCredits)
              }))),

              concatMap(() => timer(1400).pipe(tap(() => {
                this.coinDoIt = true
              }))),
            ).subscribe()

          }
        } else if (!token.isCoin) {
          if (token.reward != "default") {
            timer(0).pipe(
              concatMap(() => timer(200).pipe(tap(() => {
                this.spurDoIt = false
                this.NYGProfile.Spur = HAND.Calculate(this.NYGProfile.Spur, Number.parseInt(token.reward), /S/)
                this.updateCredits.spurs = this.NYGProfile.Spur
                this.login.UpdateSpurs(this.updateCredits)
              }))),

              concatMap(() => timer(1400).pipe(tap(() => {
                this.spurDoIt = true
              }))),

            ).subscribe({})
          }
        }
      }
    })

    this.ev.getLocifyRewards().pipe(filter(v => v.reward != "default")).subscribe({
      next: (token) => {
        this.ev.closeSetup()
        if (token.isCoin) {
          if (token.reward != "default") {
            timer(0).pipe(
              concatMap(() => timer(100).pipe(tap(() => {
                this.coinDoIt = false
                this.NYGProfile.Coin = HAND.Calculate(this.NYGProfile.Coin, Number.parseInt(token.reward), /C/)
                this.updateCredits.coins = this.NYGProfile.Coin
                this.login.UpdateCoins(this.updateCredits)
              }))),

              concatMap(() => timer(1400).pipe(tap(() => {
                this.coinDoIt = true
              }))),
            ).subscribe()

          }
        } else if (!token.isCoin) {
          if (token.reward != "default") {
            timer(0).pipe(
              concatMap(() => timer(200).pipe(tap(() => {
                this.spurDoIt = false
                this.NYGProfile.Spur = HAND.Calculate(this.NYGProfile.Spur, Number.parseInt(token.reward), /S/)
                this.updateCredits.spurs = this.NYGProfile.Spur
                this.login.UpdateSpurs(this.updateCredits)
              }))),

              concatMap(() => timer(1400).pipe(tap(() => {
                this.spurDoIt = true
              }))),

            ).subscribe({})

          }
        }
      }
    })

    var i = this.se.getItem("session")

    if (i) {
      this.id = i
      this.q._Login(this.id).subscribe({
        next: (token) => {
          if (token.data.login.updatedEvents) {
            var g = token.data.login.updatedEvents.clash
            var j = token.data.login.updatedEvents.locify

            this.NYGclashEvents = g.map(item => ({ ...item }))
            this.NYGlocifyEvents = j.map(item => ({ ...item }))
            timer(1).pipe(
              concatMap(() => timer(100).pipe(tap(() => {
                this.eventDoIt = false
              }))),
              concatMap(() => timer(1400).pipe(tap(() => {
                this.eventDoIt = true
              }))),
            ).subscribe()
          }
          if (token.data.login.updatedDisplay) {
            this.NYGProfile.Name = token.data.login.updatedDisplay.name
            // this.NYGProfile.Rating = token.data.login.updatedDisplay.nygProfile.rating
            // this.NYGProfile.GamesPlayed = token.data.login.updatedDisplay.nygProfile.gamesPlayed
            // this.NYGProfile.Tier = token.data.login.updatedDisplay.nygProfile.tier
            // this.NYGProfile.Points = token.data.login.updatedDisplay..points
            var ud = token.data.login

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

            this.OCSPlayer1v1Profile.ProfilePic.set(ud.updatedDisplay.image.imgURL)
            this.OCSPlayer2v2Profile.ProfilePic.set(ud.updatedDisplay.image.imgURL)
            this.OCSPlayerClashProfile.ProfilePic.set(ud.updatedDisplay.image.imgURL)


            this.OCSprofilePicSrc.set(token.data.login.updatedDisplay.image.imgURL)
            this.NYGProfile.ProfilePic.set(token.data.login.updatedDisplay.image.imgURL)
            this.NYGProfile.Spur = token.data.login.updatedPlayerCredits.spurs
            this.NYGProfile.Coin = token.data.login.updatedPlayerCredits.coins
            switch (true) {

              case this.NYGProfile.Rating >= 1000 && this.NYGProfile.Rating < 2000:
                this.OCSImagePath.set("tier/master.webp")
                break

              case this.NYGProfile.Rating >= 2000 && this.NYGProfile.Rating < 2500:
                this.OCSImagePath.set("tier/grand master.webp")
                break

              case this.NYGProfile.Rating >= 2500:
                this.OCSImagePath.set("tier/super master.webp")
                break

              default:
                this.OCSImagePath.set("tier/novice.webp")
            }

            this.cdr.detectChanges()
          }
        }
      })
      this.q.sub_UpdatedEvents(this.id).subscribe({
        next: (token) => {
          if (token.data?.updatedEvents) {
            var g = token.data.updatedEvents.clash
            var j = token.data.updatedEvents.locify

            this.NYGclashEvents = g.map(item => ({ ...item }))
            this.NYGlocifyEvents = j.map(item => ({ ...item }))
            timer(1).pipe(
              concatMap(() => timer(100).pipe(tap(() => {
                this.eventDoIt = false
              }))),
              concatMap(() => timer(1400).pipe(tap(() => {
                this.eventDoIt = true
              }))),
            ).subscribe()
          }
        }
      })
      const wupcq = this.q.sub_UpdatedPlayerCredits(this.id)
      wupcq.subscribe({
        next: (token) => {
          if (token.data?.updatedPlayerCredits) {
            var us = token.data.updatedPlayerCredits
            this.NYGProfile.Coin = us.coins
            this.NYGProfile.Spur = us.spurs
          }
        }
      })

      const wq = this.q.sub_UpdatedDisplay(this.id)

      wq.subscribe({
        next: (token) => {
          if (token.data?.updatedDisplay) {
            var ud = token.data.updatedDisplay
            this.NYGProfile.ProfilePic = signal(ud.image)
            this.NYGProfile.Name = ud.name
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

            this.OCSImagePath.set(token.data.updatedDisplay.image)
            this.OCSPlayer1v1Profile.ProfilePic.set(token.data.updatedDisplay.image)
            this.OCSPlayer2v2Profile.ProfilePic.set(token.data.updatedDisplay.image)
            this.OCSPlayerClashProfile.ProfilePic.set(token.data.updatedDisplay.image)

          }
        }

      })
    }
  }

  isLogin: boolean = true

  SeeProfile() {
    if (!this.isLogin) {
      this._route.navigateByUrl('/SignUp')
    } else {
      this.profs.Setup(this.OCSPlayerClashProfile, this.OCSPlayer1v1Profile, this.OCSPlayer2v2Profile)

    }
  }

  SeeDailyEvents() {
    this.eventSeen = this.ev.Setup(this.NYGclashEvents, this.NYGlocifyEvents)
  }

  toStore() {
    this.storeSeen = true
    this._route.navigate(['/Settings', "store"])
  }

  toUpgrade() {
    this._route.navigate(['/Settings', "upgrade"])
  }

  toLocify() {
    this._route.navigateByUrl('/Locify')
  }

  toClash() {
    this._route.navigateByUrl('/Clash')
  }
}