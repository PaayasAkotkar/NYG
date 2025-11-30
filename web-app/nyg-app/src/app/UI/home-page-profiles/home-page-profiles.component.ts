import { Component, Input, signal } from '@angular/core';
import { ViewPattern3Component } from "../../setup/view-pattern3/view-pattern3.component";
import { HomeProfileService } from '../home-profile/home-profile.service';
import { IHomePageProfile, HomeProfileComponent } from '../home-profile/home-profile.component';

import { OcsBtnsComponent } from "../ocs-btns/ocs-btns.component";

@Component({
  selector: 'app-home-page-profiles',
  imports: [ViewPattern3Component, HomeProfileComponent, OcsBtnsComponent],
  templateUrl: './home-page-profiles.component.html',
  styleUrls: ['./touchup/color.scss', './touchup/profile.scss', './touchup/animation.scss'],

})
export class HomePageProfilesComponent {
  @Input({ required: false }) NYGtwovtwoProfile: IHomePageProfile = { Name: "KING", Rating: 0, Record: "", Spur: "", Coin: "", GamesPlayed: 0, Tier: "", ProfilePic: signal(""), Points: 0 }
  @Input({ required: false }) NYGclashProfile: IHomePageProfile = { Name: "KING", Rating: 0, Record: "", Spur: "", Coin: "", GamesPlayed: 0, Tier: "", ProfilePic: signal(""), Points: 0 }
  @Input({ required: false }) NYGonevoneProfile: IHomePageProfile = { Name: "KiNG", Rating: 0, Record: "", Spur: "", Coin: "", GamesPlayed: 0, Tier: "", ProfilePic: signal(""), Points: 0 }

  constructor(private prof: HomeProfileService) { }
  viewClash: boolean = false
  viewLocify: boolean = false
  toClash() {
    this.viewClash = !this.viewClash
    this.viewLocify = false
  }
  toLocify() {
    this.viewLocify = !this.viewLocify
    this.viewClash = false
  }

}

