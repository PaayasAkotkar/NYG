import { Component, Input, OnInit } from '@angular/core';
import { IProfile_, JoinedProfile } from '../../lobby/resources/lobbyResource';
import { OcsDetComponent } from "../../profile/ocs-det/ocs-det.component";




@Component({
  selector: 'app-ocs-miniview',
  imports: [OcsDetComponent],
  templateUrl: './ocs-miniview.component.html',
  styleUrls: ['./touchup/animation.scss', './touchup/color.scss', './touchup/home-profile.scss'],
})
export class OcsMiniviewComponent implements OnInit {
  @Input() NYGProfile: IProfile_
  img: string = ""
  grandMaster = '/tier/grand master.webp'
  master = '/tier/master.webp'
  novice = '/tier/novice.webp'
  superMaster = '/tier/super master.webp'
  ngOnInit(): void {
    switch (this.NYGProfile.tier) {
      case "novice":
        this.img = this.novice
        break
      case "grand master":
        this.img = this.grandMaster
        break
      case "super grand master":
        this.img = this.superMaster
        break
      case "master":
        this.img = this.master
        break
    }
  }
}
