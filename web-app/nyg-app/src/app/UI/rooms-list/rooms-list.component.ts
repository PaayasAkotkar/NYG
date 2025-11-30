import { ChangeDetectorRef, Component, Input, OnChanges, OnInit, signal, SimpleChanges, WritableSignal } from '@angular/core';
import { RoomService } from '../../roomService/room.service';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { HAND } from '../../helpers/hand';
import { RoomListAgent, RoomsPattern } from './agent/room-list';

import { Subscription } from 'rxjs';
import { Router } from '@angular/router';

@Component({
  selector: 'app-rooms-list',
  imports: [ReactiveFormsModule],
  templateUrl: './rooms-list.component.html',
  styleUrls: ['./touchup/room-list.scss', './touchup/animation.scss', './touchup/color.scss',
    './touchup/join.scss'
  ]
})
export class RoomsListComponent implements OnInit {

  @Input({ required: false }) OCSRooms: WritableSignal<RoomsPattern[]> = signal([
  ])
  TemplateX: RoomsPattern[] = []
  JoinRoomForm: FormGroup
  collectionKeys: string[] = []
  @Input({ required: false }) proceed: boolean = false
  watch = new Subscription()

  constructor(private router: Router, private cdr: ChangeDetectorRef, private _room: RoomService, private FB: FormBuilder) { }

  ngOnInit(): void {
    this.JoinRoomForm = this.FB.group({})
    this._room.connect()
    this.watch = this._room.Rooms().subscribe({
      next: (token) => {
        const RoomList = /RoomLists:(?=[\s\S\w])/
        const getRooms = RoomList.test(token)
        switch (true) {
          case getRooms:
            var parse = "RoomLists:"
            var forate: RoomsPattern = { RoomName: "", Category: "", Type: "", Book: "" }
            var _get = this.afterColonJson(token, parse, forate)
            this.OCSRooms.set(_get)
            this.TemplateX.push(_get)
            this.FillControls()
            this.cdr.detectChanges()
            break
        }
      }
    })
  }
  afterColonJson(token: string, parse: string, _interface: any) {
    token = token.substring(parse.length)
    _interface = JSON.parse(token)
    return _interface
  }

  Join(e: Event) {
    e.preventDefault()
    var room = this.collectionKeys[0].toLowerCase()
    this._room.joinRoom(room, false, "")
    this.router.navigate(['/Game-Hall'])

  }

  CollectionToken(token: string) {
    this.collectionKeys.push(token)
    const len = this.collectionKeys.length

    if (len > 1) {
      const newly = this.collectionKeys[len - 1]

      var temp

      if (newly) {
        temp = this.collectionKeys[this.collectionKeys.indexOf(newly) - 1]
      }

      const old = temp

      // this condition prevents the newly and old value to get false  
      if (old && this.JoinRoomForm.get(newly)?.value == true) {
        this.collectionKeys = this.collectionKeys.filter(e => e != old)
        this.collectionKeys.forEach(e => {
          if (this.JoinRoomForm.get(e)?.value == true) {
            this.JoinRoomForm.get(old)?.setValue(false, { emitEvent: false })
          }
        })
      }
    }
  }

  FillControls(): void {
    this.OCSRooms().forEach(item => {
      this.JoinRoomForm.addControl<any>(item.RoomName, this.FB.control(false))
    })
  }
}
