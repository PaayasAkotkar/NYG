import { Component, Input, OnInit } from '@angular/core';

import { BarComponent } from "../UI/bar/bar.component";
import { IGameOver } from '../lobby/resources/lobbyResource';
import { GameOverService } from './game-over.service';
@Component({
  selector: 'app-game-over',
  imports: [],
  templateUrl: './game-over.component.html',
})
export class GameOverComponent implements OnInit {
  @Input({ required: false }) OCSgameOverInfo: IGameOver
  ngOnInit(): void {
    this.ga.Setup(this.OCSgameOverInfo)
  }
  constructor(private ga: GameOverService) { }
}
