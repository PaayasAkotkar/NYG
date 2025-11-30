import { ConnectedPosition, Overlay, OverlayConfig, OverlayModule, OverlayRef } from '@angular/cdk/overlay';
import { Component, ElementRef, Input, ViewChild, ViewEncapsulation } from '@angular/core';
import { MatCard, MatCardModule } from '@angular/material/card';
import { ViewSheetService } from '../view-sheet/view-sheet.service';
import { LobbyResource } from '../lobby/resources/lobbyResource';
import { RoomService } from '../roomService/room.service';

interface displayPro {
  name: string,
  team: string
}

@Component({
  selector: 'app-statboard',
  imports: [MatCardModule, OverlayModule],
  templateUrl: './statboard.component.html',
  styleUrls: ['./touchup/statboard.scss',
    './touchup/sheet.scss', './touchup/rating.scss'
  ],
  encapsulation: ViewEncapsulation.Emulated
})
export class StatboardComponent {
  Set = ['S1', 'S2', 'S3']
  Round = ['R1', 'R2']
  constructor(private service: ViewSheetService, private room: RoomService) {

  }
  // close() {
  //   this.service.closeSetup()
  // }

  @Input() playerName: any = "NAF" // not able to find
  @Input() teamName: any = "NYG" // now you guess
  @Input() rating: any = "99" // rating
  @Input() _update: string = ""
}
