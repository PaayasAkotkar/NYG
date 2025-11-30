import { Component, Output, EventEmitter } from '@angular/core';
import { OCSbuttonV1 } from "../../UI/upgrade-menu/upgrade-part1/ocs-button.component";
import { Ref11xScoreboardComponent } from '../../UI/scoreboard-ref/scoreboard-a/Ref11xScoreboard.component';
import { Ref25xScoreboardComponent } from '../../UI/scoreboard-ref/scoreboard-b/Ref25xScoreboard.component';
import { Ref11xGuessTextComponent } from "../../UI/guess-ref/guess-text-a/Ref11xGuessText.component";
import { Ref25xGuessTextComponent } from "../../UI/guess-ref/guess-text-b/Ref25xGuessText.component";

import { OcsBtnsComponent } from "../../UI/ocs-btns/ocs-btns.component";
import { UpdateLivelyService } from '../../update-lively/update-lively.service';

@Component({
  selector: 'view-theme',
  imports: [Ref11xGuessTextComponent, Ref25xGuessTextComponent, Ref11xScoreboardComponent, Ref25xGuessTextComponent, OCSbuttonV1, OcsBtnsComponent, Ref25xGuessTextComponent, Ref11xGuessTextComponent, Ref11xScoreboardComponent, Ref25xScoreboardComponent],
  templateUrl: './theme.component.html',
  styleUrl: './theme.component.scss'
})
export class ViewThemeComponent {
  PageNavigation: number = 1
  Navigate(e: number) {
    this.PageNavigation = parseInt(String(e))
  }
  @Output() OCSEVENT = new EventEmitter<string>()
  constructor(private login: UpdateLivelyService) { }
  ThemeName: "text25x" | "text11x" = "text11x"
  boardTheme: "text25x" | "text11x" = "text11x"
  guessTheme: "text25x" | "text11x" = "text11x"

  Check() {
    if (this.PageNavigation == 1) {
      this.boardTheme = "text25x"
      this.guessTheme = 'text25x'
    } else {
      this.boardTheme = "text11x"
      this.guessTheme = "text11x"

    }
    this.OCSEVENT.emit(this.ThemeName)
    this.login.UpdateTheme(this.boardTheme, this.guessTheme)
  }
}
