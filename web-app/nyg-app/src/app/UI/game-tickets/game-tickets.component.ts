import { Component, OnInit } from '@angular/core';
import { OnTeamColorDirective } from './on-team-color.directive';

// on hold
interface GameTickets {
  PlayerName: string
  Team: string
  MVP: number // points
  Position: string
  isTeamRed: boolean
}
@Component({
  selector: 'app-game-tickets',
  imports: [OnTeamColorDirective],
  templateUrl: './game-tickets.component.html',
  styleUrls: ['./touchup/game-tickets.scss', './touchup/color.scss']
})
export class GameTicketsComponent implements OnInit {
  tickets: GameTickets[] = []
  D: string[] = []
  ngOnInit(): void {
    this.D = ['NAME', 'TEAM', 'VOTES']
    this.tickets = [
      {
        PlayerName: 'KORTNEY',
        Team: 'RED',
        MVP: 11,
        Position: 'MVP',
        isTeamRed: true,
      },
      {
        PlayerName: 'JJ',
        Team: 'BLUE',
        MVP: 4,
        Position: 'RUNNER-UP',
        isTeamRed: true,

      },

    ]
  }
}
