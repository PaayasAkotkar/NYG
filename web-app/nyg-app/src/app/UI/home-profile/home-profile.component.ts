import { Component, CSP_NONCE, Input, OnInit, signal, WritableSignal } from '@angular/core';
import { OcsDetComponent } from "../../profile/ocs-det/ocs-det.component";

import { OcsImgComponent } from "../../profile/ocs-img/ocs-img.component";
import { HomeProfileHeaderComponent } from "./home-profile-header/home-profile-header.component";
import { HomeProfileService } from './home-profile.service';
export interface IHomePageProfile {
  Name: string
  Rating: number
  Record: string
  Points: number
  GamesPlayed: number
  Tier: string
  ProfilePic: WritableSignal<string>
  Coin: string
  Spur: string
}
@Component({
  selector: 'app-home-profile',
  imports: [OcsDetComponent, OcsImgComponent, HomeProfileHeaderComponent],
  templateUrl: './home-profile.component.html',
  styleUrls: ['./touchup/animation.scss', './touchup/color.scss', './touchup/home-profile.scss'],

})
export class HomeProfileComponent implements OnInit {
  @Input({ required: true }) NYGProfile: IHomePageProfile = { Name: "", Rating: 0, Record: "", GamesPlayed: 0, Tier: "", ProfilePic: signal(""), Coin: "", Spur: "", Points: 0 }
  @Input() NYGshowIMG: boolean = true
  ngOnInit(): void { }
  constructor(private j: HomeProfileService) { }
  close() {
    this.j.closeSetup()
  }
}
