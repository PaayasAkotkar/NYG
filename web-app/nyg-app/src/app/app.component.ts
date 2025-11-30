import { Component, DestroyRef, HostListener, inject, OnInit, signal, ViewEncapsulation } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { RoomService } from './roomService/room.service';
import { MenuComponent } from "./menu/menu.component";
import { myAnimation } from './animation';
import { LoadLogoService } from './UI/load-logo/load-logo.service';
import { LoginService } from './login-service/login.service';
import { MiscellaenousService } from './storage/miscellaenous/miscellaenous.service';
import { FreezePowerService } from './UI/power-up/power-freeze/power-freeze.service';
import { LocifyGame, RoomSettings } from './lobby/resources/lobbyResource';
interface Profile {
  name: string
  nickname: string
  img: string
  id: string
}
@Component({
  selector: 'app-root',
  animations: [myAnimation],
  imports: [MenuComponent, RouterOutlet],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
  encapsulation: ViewEncapsulation.Emulated
})
export class AppComponent implements OnInit {
  title = 'NYG';

  HIDE: boolean = true
  Done: boolean = false
  isAlive: boolean = true

  guard = inject(DestroyRef)

  isDone(e: any) {
    this.Done = e
  }

  constructor(private fs: FreezePowerService, private _setup: MiscellaenousService, private room: RoomService, private load: LoadLogoService, private login: LoginService) {
    this.guard.onDestroy(() => { this.isAlive = false })
  }
  reqGuard: boolean = false
  lockGuard: boolean = false
  ngOnInit(): void {
    console.log('NOW YOU GUESS by PAAYAS AKOTKAR')
    if (!this.reqGuard) {
      this._setup.InitStorage()
      // important to that the login is done successfully
      this.room.connect()
      this.login.Init()
      this.login.GetEvents()
      this.reqGuard = true
      this.room.Rooms().subscribe({
        next: (token) => {
          var LGPattern = /LocifyGame:(?=[\s\S\w])/
          var _game = LGPattern.test(token)
          var LRPattern = /LocifyRoom:(?=[\s\S\w])/
          var roomDets = LRPattern.test(token)


          switch (true) {
            case _game:
              var p = 'LocifyGame:'
              var lg: LocifyGame = {
                redScore: 0, blueScore: 0,
                round: 0, set: 0, powers: {},
                session: false, mode: "", roomname: "",
                teamname: "", nicknames: [], stats: {}
              }
              var _lg: LocifyGame = this.afterColonJson(token, p, lg)
              this.lockGuard = _lg.session;
              break
            case roomDets:
              var p = 'LocifyRoom:'
              var lr: RoomSettings = {
                category: "", mode: "",
                book: "", gameTime: -1, decisionTime: -1,
                powers: {}, cap: 0, isEntertainment: false,
                roomname: "", friend: false, private: false
                , code: "", starter: false, setupToss: false, reverse: false,
                client: false, session: false
              }

              var _lr: RoomSettings = this.afterColonJson(token, p, lr)
              this.lockGuard = _lr.session
          }



        }
      })

    }
  }

  public getRouterOutletState(outlet: RouterOutlet) {
    return outlet.isActivated ? outlet.activatedRoute : '';
  }
  afterColonJson(token: string | undefined, parse: string, _interface: any) {
    if (token) {
      token = token.substring(parse.length)
      _interface = JSON.parse(token)
    }
    return _interface
  }
  Load() {
    if (this.isAlive)
      this.load.Setup()
  }

  UnLoad() {
    if (this.isAlive)
      this.load.closeSetup()
  }

  @HostListener('document:keydown', ['$event'])
  HideMenu(event: KeyboardEvent) {
    if (event.key == 'Control') {
      this.HIDE = !this.HIDE
    }
  }

}
