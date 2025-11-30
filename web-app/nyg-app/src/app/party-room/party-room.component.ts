import { ChangeDetectorRef, Component, HostListener, Input, NgZone, OnChanges, OnDestroy, OnInit, signal, SimpleChanges, ViewEncapsulation, WritableSignal } from '@angular/core';
import { IProfile_, JoinedProfile, RoomSettings, RoomSettingsParcel } from '../lobby/resources/lobbyResource';
import { RoomService } from '../roomService/room.service';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { PartyChatService } from '../partychatSerivce/party-chat.service';
import { BehaviorSubject, delay, distinctUntilChanged, of, repeat, single, Subject, Subscription, switchMap, takeUntil, } from 'rxjs';
import { CountComponent } from "../UI/count/count.component";
import { FriendCodeService } from '../UI/friend-code/friend-code.service';
import { RoomNameComponent } from "../UI/room-name/room-name.component";
import { FriendCodeComponent } from "../UI/friend-code/friend-code.component";
import { RoomSettingsService } from '../UI/room-settings/room-settings.service';
import { LobbyBoxComponent } from "../UI/lobby-box/lobby-box.component";
import { TagsComponent } from "../UI/tags/tags.component";
import { DragDropModule } from '@angular/cdk/drag-drop';
import { LongPollingService } from '../long-polling/long-polling.service';
import { OCSFieldFormat } from '../helpers/patterns';
import { NYGqService, PutMessage } from '../graphql/nygq.service';
import { gql } from 'apollo-angular';
import { SessionService } from '../storage/session/session.service';
import { v4 } from 'uuid';
import { Router } from '@angular/router';

// incomplete

interface FormRegisters {
  nickName: string
  id: string
  roomID: string
  token: string
}

@Component({
  selector: 'app-party-room',
  imports: [DragDropModule, ReactiveFormsModule, CommonModule, CountComponent, RoomNameComponent, FriendCodeComponent, LobbyBoxComponent, TagsComponent],
  templateUrl: './party-room.component.html',
  styleUrls: ['./touchup/animation.scss', './touchup/color.scss', './touchup/party-room.scss', './touchup/box.scss'],
  encapsulation: ViewEncapsulation.Emulated
})
export class PartyRoomComponent implements OnInit, OnDestroy {
  constructor(private router: Router, private se: SessionService, private q: NYGqService, private options: RoomSettingsService, private room: RoomService, private fc: FriendCodeService, private pc: PartyChatService, private cdr: ChangeDetectorRef,
    private zone: NgZone) {
  }

  Code: string = ""
  NYGmode: string = ""
  NYGnickName: string = ""
  @Input({ required: true }) NYGmyNickname: string = ""
  sportsBookCollection: OCSFieldFormat[]

  @Input({ required: false }) NYGentertainmentCollection: OCSFieldFormat[]
  @Input({ required: true }) NYGcollection: WritableSignal<OCSFieldFormat[]>
  @Input({ required: false }) NYGactiveConn: WritableSignal<number>
  @Input({ required: true }) NYGroomSettings: RoomSettings
  @Input({ required: true }) NYGplayersProfile: JoinedProfile[] = []
  Sub = new Subject<void>()

  getPartyChatToken: string = ""
  texts: string[] = []
  latest: WritableSignal<string[]> = signal([])

  // chatting
  partyChatForm = new FormGroup(
    { PartyChatToken: new FormControl(' '), }
  )

  ngOnInit(): void {
    this.q._Books().pipe(takeUntil(this.Sub)).subscribe({
      next: (token) => {
        var eb = token.data.updatedBooks.entertainment
        var etemp: OCSFieldFormat[] = []
        for (let i of eb) {
          etemp.push({ OCStemplate: i, OCSformControl: false, OCSToolTip: "" })
        }
        var sb = token.data.updatedBooks.sports
        var stemp: OCSFieldFormat[] = []
        for (let i of sb) {
          stemp.push({ OCStemplate: i, OCSformControl: false, OCSToolTip: "" })
        }
        this.NYGentertainmentCollection = etemp
        this.sportsBookCollection = stemp

      }
    })
    var f: boolean = this.NYGroomSettings.friend || this.NYGroomSettings.private
    var c = this.NYGroomSettings.code
    var m = this.NYGroomSettings.mode

    this.NYGnickName = this.NYGmyNickname

    if (f) {
      this.Code = c
    }
    this.NYGmode = m
    this.q.sub_LatestMsg(this.NYGroomSettings.roomname).pipe(takeUntil(this.Sub)).subscribe({
      next: (token) => {
        this.zone.run(() => {
          if (token.data?.latest) {
            var ltest = token.data.latest
            this.texts = [...this.texts, ltest.msg]
            this.latest.set(this.texts)
            this.NYGnickName = ltest.name
            this.cdr.detectChanges()
          }
        })
      }
    })
    this.cdr.detectChanges()
  }

  ngOnDestroy(): void {
    this.Sub.next()
    this.Sub.complete()
  }

  Options() {
    var ie = signal(this.NYGroomSettings.isEntertainment)
    var pow = this.NYGroomSettings.powers
    this.NYGroomSettings.cap = 2
    var max = 4
    if (ie() == true) {
      this.NYGcollection.set(this.NYGentertainmentCollection)
    } else {
      this.NYGcollection.set(this.sportsBookCollection)
    }
    this.options.Setup(
      this.NYGcollection,
      ie, pow, max, this.NYGroomSettings
    )
  }

  Leave() {
    this.room.Leave(this.NYGroomSettings.roomname)
    this.router.navigateByUrl('/Home')
  }

  Chat(e: Event) {
    if (e.cancelable && !e.defaultPrevented) {
      e.preventDefault()
    }
    var i = this.se.getItem('session')
    if (i) {
      this.ID = i

      var chat = this.partyChatForm.get('PartyChatToken')?.value
      if (chat) {
        this.getPartyChatToken = chat
        var m: PutMessage = { msg: this.getPartyChatToken, id: this.ID, roomname: this.NYGroomSettings.roomname, name: this.NYGnickName }
        // if we do it via graphql than it will be two same tokens from one name meaning repeated
        this.q.publishMessage(m).pipe(takeUntil(this.Sub)).subscribe({
          error: (err) => {
            console.error(err)
          }
        })
        this.cdr.detectChanges()
        this.partyChatForm.get('PartyChatToken')?.setValue('')
      }
    }
  }

  ID: string = ""
  @HostListener('document:keydown', ['$event'])
  HideMenu(event: KeyboardEvent) {
    if (event.key == 'Enter') {
      this.Chat(event)
    }
  }
}
