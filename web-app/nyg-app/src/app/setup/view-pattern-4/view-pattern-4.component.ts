import { Component, Input, signal } from '@angular/core';
import { HomeProfileService } from '../../UI/home-profile/home-profile.service';
import { IHomePageProfile, HomeProfileComponent } from '../../UI/home-profile/home-profile.component';
import { ViewPattern3Component } from "../view-pattern3/view-pattern3.component";


import { OcsBtn2Component } from "../ocs-btn-2/ocs-btn-2.component";

@Component({
  selector: 'app-view-pattern-4',
  imports: [ViewPattern3Component, HomeProfileComponent, OcsBtn2Component],
  templateUrl: './view-pattern-4.component.html',
  styleUrls: ['./touchup/color.scss', './touchup/profile.scss', './touchup/animation.scss'],

})
export class ViewPattern4Component {
  @Input({ required: false }) NYGtwovtwoProfile: IHomePageProfile = { Name: "KING", Rating: 0, Record: "", Spur: "", Coin: "", GamesPlayed: 0, Tier: "", ProfilePic: signal(""), Points: 0 }
  @Input({ required: false }) NYGclashProfile: IHomePageProfile = { Name: "KING", Rating: 0, Record: "", Spur: "", Coin: "", GamesPlayed: 0, Tier: "", ProfilePic: signal(""), Points: 0 }
  @Input({ required: false }) NYGonevoneProfile: IHomePageProfile = { Name: "KING", Rating: 0, Record: "", Spur: "", Coin: "", GamesPlayed: 0, Tier: "", ProfilePic: signal(""), Points: 0 }

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
