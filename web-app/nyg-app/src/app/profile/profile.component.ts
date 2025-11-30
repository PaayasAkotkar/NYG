import { Component, input, Input, OnInit, signal, ViewEncapsulation, WritableSignal } from '@angular/core';
import { OcsImgComponent } from "./ocs-img/ocs-img.component";

import { OcsNmeComponent } from "./ocs-nme/ocs-nme.component";

@Component({
  selector: 'app-profile',
  imports: [OcsImgComponent, OcsNmeComponent],
  templateUrl: './profile.component.html',
  styleUrls: ['./touchup/profile.scss', './touchup/color.scss'],
  encapsulation: ViewEncapsulation.Emulated
})
export class ProfileComponent implements OnInit {
  @Input({ required: true }) OCSname: WritableSignal<string>
  // @Input({ required: true }) OCSgamesPlayed: WritableSignal<number>
  // @Input({ required: true }) OCSrating: WritableSignal<number>
  // @Input({ required: true }) OCSpoints: WritableSignal<number>
  // @Input({ required: true }) OCSrecord: WritableSignal<string>
  // @Input({ required: true }) OCStier: WritableSignal<string>
  @Input({ required: true }) OCSImageURL: WritableSignal<string> = signal('/default/default-profile.jpg')
  OCSImagePath: WritableSignal<string> = signal("")
  ngOnInit(): void {
    switch (true) {

      // case this.OCSrating() >= 1000 && this.OCSrating() < 2000:
      //   this.OCSImagePath.set("tier/master.webp")
      //   break

      // case this.OCSrating() >= 2000 && this.OCSrating() < 2500:
      //   this.OCSImagePath.set("tier/grand master.webp")
      //   break

      // case this.OCSrating() >= 2500:
      //   this.OCSImagePath.set("tier/super master.webp")
      //   break

      default:
        this.OCSImagePath.set("tier/novice.webp")

    }
  }
}
