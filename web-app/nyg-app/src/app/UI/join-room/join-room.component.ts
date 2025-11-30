import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { RoomService } from '../../roomService/room.service';
import { Router } from '@angular/router';
import { LobbyResource } from '../../lobby/resources/lobbyResource';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { OcsInputComponent } from "../ocs-input/ocs-input.component";
import { filter } from 'rxjs';


@Component({
  selector: 'app-join-room',
  imports: [ReactiveFormsModule, OcsInputComponent],
  templateUrl: './join-room.component.html',
  styleUrls: ['./touchup/container.scss',
    './touchup/join-room.scss', './touchup/color.scss',
    './touchup/animations.scss'
  ],
  encapsulation: ViewEncapsulation.Emulated
})
export class JoinRoomComponent implements OnInit {
  constructor(private room: RoomService, private router: Router) {
  }
  // join room
  joinRoomForm = new FormBuilder().group({
    JoinToken: ['', [Validators.required]],
    // JoinNickNameToken: ['', [Validators.required]],
    JoinCodeToken: ['', []],
  })
  nameToolTip: string = "name of the room"
  nicknameToolTip: string = "name to display assoicated players"
  codeToolTip: string = "code of the room"
  isNYGnameError: boolean = false
  isNYGnicknameError: boolean = false
  NDEFAULT = "the name is not valid remove"
  SDEFAULT = "the filed must not left empty"

  NYGnameError: string = ""
  NYGnicknameError: string = ""
  NYGcategoryError: string
  NYGtypeError: string
  isNYGcategoryError: boolean = false
  isNYGtypeError: boolean = false
  isNYGpowerError: boolean = false
  NYGpowerError: string
  proccedNavigation: boolean = false
  PDEFAULT = "MUST SELECT 4 POWERs"
  getJoinRoomToken = String(this.joinRoomForm.get('JoinToken')?.value)
  getJoinRoomCode = String(this.joinRoomForm.get('JoinCodeToken')?.value)
  // getJNickName: string = String(this.joinRoomForm.get('JoinNickNameToken')?.value)
  Friend: boolean = false
  ngOnInit(): void {

    this.NYGnameError = "the name is not valid  "
    const restriction = /[^a-zA-Z]$/;
    this.joinRoomForm.get('JoinToken')?.valueChanges.pipe(filter(token => token != undefined && typeof token == 'string')).subscribe({
      next: (token) => {
        if (restriction.test(token)) {
          for (let i = 0; i < token.length; i++) {
            if (restriction.test(token[i])) {
              if (!this.NYGnameError.includes(token[i])) {

                this.isNYGnameError = true
                this.NYGnameError += token[i]
              }
            }
          }
          token = token.replace(restriction, "")
          setTimeout(() => {
            this.isNYGnameError = false
            this.NYGnameError = this.NDEFAULT
          }, 2500)
          this.joinRoomForm.get('JoinToken')?.setValue(token, { emitEvent: false })
        }
        else if (token.length > 8) {
          this.isNYGnameError = true
          this.NYGnameError = "REACHED MAX LENGTH 8"
          this.joinRoomForm.get('JoinToken')?.setValue(token.substring(0, 8))
          setTimeout(() => {
            this.isNYGnameError = false
          }, 3500)
        }

      }
    })

  }

  JoinRoom(e: Event) {
    e.preventDefault()
    this.getJoinRoomToken = String(this.joinRoomForm.get('JoinToken')?.value)
    this.getJoinRoomCode = String(this.joinRoomForm.get('JoinCodeToken')?.value)

    if (this.getJoinRoomCode != "") {
      this.Friend = false
      this.room.joinRoom(this.getJoinRoomToken, true,
        this.getJoinRoomCode)
    } else {
      this.Friend = false
      this.room.joinRoom(this.getJoinRoomToken, false, "")
    }
    this.router.navigateByUrl('/Game-Hall')

  }
}
