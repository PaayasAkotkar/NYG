import { Component, NgZone, OnInit, ViewEncapsulation } from '@angular/core';
import { Overlay, OverlayConfig, OverlayModule } from '@angular/cdk/overlay';
import { ComponentPortal, PortalModule } from '@angular/cdk/portal';
import { Route, Router, RouterModule, RouterOutlet } from '@angular/router';
import { distinct, Observable, Subscription } from 'rxjs';
import { LobbyResource } from '../lobby/resources/lobbyResource';
import { RoomService } from '../roomService/room.service';
import { MatCardModule } from '@angular/material/card';
import { ViewMenuComponent } from '../view-menu/view-menu.component';
import { MenuService } from '../view-menu/view-menu.service';
import { CommonModule } from '@angular/common';
@Component({
  selector: 'app-menu',
  imports: [MatCardModule, RouterModule, CommonModule],
  templateUrl: './menu.component.html',
  styleUrls: ['./touchup/menu.scss', './touchup/color.scss', './touchup/head.scss',
    './touchup/animations.scss'
  ],
  encapsulation: ViewEncapsulation.Emulated
})
export class MenuComponent extends LobbyResource implements OnInit {
  MenuNav = ['H', 'P', 'S'] // s is for store
  show: boolean = false
  VISIT = ['youtube', 'instagram', 'reddit', 'facebook', 'x']

  NYGheaders = [
    {
      template: 'PLAY NOW', template2: 'HOME', class: 'PN-class',

    }, { template: 'QUICK MATCH', template2: 'CONNECT', class: 'QM-class', },
    { template: 'MY ROOM', template2: 'CREATE', class: 'MR-class', },
    { template: 'SETUP', template2: 'OPTIONS', class: 'S-class', },

  ]

  constructor(private __run: NgZone, private router: Router, private _R: RoomService, private service: MenuService) {
    super(_R, __run)
  }
  ngOnInit(): void {
  }
  backHome: boolean = false
  Play() {
    this.router.navigateByUrl('/Home')
    this.backHome = !this.backHome

    this.service.closeSetup()
  }

  Home() {
    this.router.navigateByUrl('/Home')

    this.service.closeSetup()
  }
  _Join() {
    this.router.navigateByUrl('/Join')
    this.service.closeSetup()
  }
  _Make() {
    this.router.navigateByUrl('/Make')
    this.service.closeSetup()
  }
  _Settings() {
    this.router.navigate(['/Settings', "setup"])
    this.service.closeSetup()
  }
}
