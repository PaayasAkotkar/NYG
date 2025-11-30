import { Component, Input, OnInit } from '@angular/core';

import { ProgressBarComponent } from "../progress-bar/progress-bar.component";

import { IGameOver } from '../../lobby/resources/lobbyResource';

interface GameResultPattern {
  newRating: number // will be +11 or....
  previousRating: number
  record: number
  powersUsed: string[]
  challengesSet: string[]
  gamePoints: number
  tier: string
}

interface SaveC {
  heading: string
  template: string
}

@Component({
  selector: 'app-game-result',
  imports: [ProgressBarComponent],
  templateUrl: './game-result.component.html',
  styleUrls: ['./touchup/animation.scss', './touchup/game-result.scss', './touchup/color.scss']
})
export class GameResultComponent implements OnInit {
  @Input({ required: true }) OCSinfo: IGameOver = {
    ratingDecrementedBy: 0, ratingIncrementedBy: 0, newRating: 0, alert: false
  }
  OCSRating: number = 400
  OCScurrentProgress: number = 1000
  OCSmaxProgress: number = 1500
  OCSminProgress: number = 1500
  OCSratingIncreasedBy: number = 24

  display: string = ""
  grandMaster = '/tier/grand master.webp'
  master = '/tier/master.webp'
  novice = '/tier/novice.webp'
  superMaster = '/tier/super master.webp'
  // 1400, 2000, 2500


  ngOnInit(): void {

    this.OCSRating = this.OCSinfo.newRating

    if (this.OCSRating < 1400) {
      this.OCSmaxProgress = 2000
      this.OCSminProgress = 400
    } else if (this.OCSRating >= 1400 && this.OCSRating < 2000) {
      this.OCSmaxProgress = 2400
      this.OCSminProgress = 1400

    } else if (this.OCSRating >= 2000 && this.OCSRating < 2400) {
      this.OCSmaxProgress = 2400
      this.OCSminProgress = 2000

    } else {
      this.OCSminProgress = 2500
      this.OCSmaxProgress = 8000
    }

    this.OCSratingIncreasedBy = this.OCSinfo.ratingIncrementedBy + this.OCSinfo.ratingDecrementedBy
    this.OCScurrentProgress = (this.OCSRating - this.OCSminProgress) / (this.OCSmaxProgress - this.OCSminProgress) * 100
    this.updateDisplay()
  }
  updateDisplay() {

    switch (true) {
      // to reach master
      case this.OCSRating >= 1400 && this.OCSRating < 2000:
        this.display = this.master
        break
      // to reach grandmaster
      case this.OCSRating >= 2000 && this.OCSRating < 2400:
        this.display = this.grandMaster
        break
      // to reach super grandmaster
      case this.OCSRating >= 2400 && this.OCSRating < 2500:
        this.display = this.superMaster
        break
      default:
        this.display = this.novice
        break
    }
  }
}
