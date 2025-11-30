import { AfterViewInit, Component, OnInit, signal, WritableSignal } from '@angular/core';
import { RefSetupBoardComponent } from "../ref-setup-board/ref-setup-board.component";

import { OCSFieldFormat } from '../../helpers/patterns';
import { OcsCardComponent } from './ocs-card/ocs-card.component';

import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { PowerNexusComponent } from "../power-up/power-nexus/power-nexus.component";
import { PowerDrawComponent } from "../power-up/power-draw/power-draw.component";
import { Ref25xGuessTextComponent } from "../guess-ref/guess-text-b/Ref25xGuessText.component";


import { PowerFreezeComponent } from "../power-up/power-freeze/power-freeze.component";



import { BetGuessComponent } from "../power-up/power-bet/bet-picker/bet-guess.component";
import { RecBetItems } from '../../lobby/resources/lobbyResource';
import { RefTagPowerComponent } from "../ref-tag-power/ref-tag-power.component";
import { RefRewindPowerComponent } from "../ref-rewind-power/ref-rewind-power.component";
import { CdkAccordionModule } from '@angular/cdk/accordion';
import { OcsHeaderComponent } from '../../setup/ocs-header/ocs-header.component';
import { OcsHeaderV2Component } from '../ocs-header-v2/ocs-header-v2.component';

@Component({
  selector: 'app-ocs-manual',
  imports: [CdkAccordionModule, ReactiveFormsModule, RefSetupBoardComponent, OcsCardComponent, PowerNexusComponent, PowerDrawComponent, Ref25xGuessTextComponent, PowerFreezeComponent, BetGuessComponent, RefTagPowerComponent, RefRewindPowerComponent],
  templateUrl: './ocs-manual.component.html',
  styleUrls: ['./touchup/ocs-grid.scss', './touchup/manual.scss', './touchup/color.scss',
    './touchup/animation.scss'
  ],
})
export class OcsManualComponent implements OnInit, AfterViewInit {
  NYGclashMode: WritableSignal<string> = signal('node')
  NYGlocifyMode: WritableSignal<string> = signal('locify')
  islNYGClash: WritableSignal<boolean> = signal(false)
  isNYGClash: WritableSignal<boolean> = signal(true)
  forgroup: FormGroup = new FormBuilder().group({})
  _credits: string = `A PAAYAS AKOTKAR CREATION`
  _credition: string = `ONE COLOR STUDIO`


  aboutClash: string =
    `Clash
          A turn based <strong>ROUND</strong> where <strong>4</strong> Players go up against each other
    Each player <strong>FIGHTS</strong> for the <strong>CHANCES</strong>
    And yeah not to forget that if the hand is <strong>HAND</strong> you <strong>OUT</strong>
And most important you can go with your OWN <strong>POWER DECK</strong> 
`
  aboutLocify: string = `
  Locify
        A self hosted room either for 2v2 or 1v1 
For 2v2 
the lobby is divided into red and blue team 
Each player has its turn to play a round 
While the partner can watch its team game by spectating
Team with most amount of points in 6 rounds wins the game
`
  aboutLocifysecond: string = `
Fun fact: powers are divided randomly so creator choose wisely the powers
For 1v1
both the players fight for the points the most amount of points wins the game  
  `
  clashSS: string = './ss/clash.png'
  NYGlocifyList: WritableSignal<OCSFieldFormat[]> = signal(
    [
      { OCStemplate: "2015 new Mid School", OCSformControl: false, OCSToolTip: "" },
      { OCStemplate: "2016 got into NBA", OCSformControl: false, OCSToolTip: "" },
      { OCStemplate: "2017 trip to BANGLORE", OCSformControl: false, OCSToolTip: "" },
      { OCStemplate: "2018 welcome to NAGPUR", OCSformControl: false, OCSToolTip: "" },
      { OCStemplate: "2019 graducated secondary school", OCSformControl: false, OCSToolTip: "" },
      { OCStemplate: "2020 welcome back to PUNE", OCSformControl: false, OCSToolTip: "" },
      { OCStemplate: "2021 reunion with old FRIENDS", OCSformControl: false, OCSToolTip: "" },

    ]
  )
  NYGclashList: WritableSignal<OCSFieldFormat[]> = signal([
    { OCStemplate: "2001 I AM HERE", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2009 welcome to PUNE", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2010 Elementary School", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2011 got into CRICKET", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2012 went to Playground", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2013 too much into EA SPORTs", OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: "2014 trip to KOKAN", OCSformControl: false, OCSToolTip: "" },
  ])
  Bet: string = ` just like a pocker we bet on something that we know the opponent might pick
          if the opponent picks your bet you gains the point `

  Nexus: string = ` a hint that can help you recall the word`

  Covert: string = ` hide the word that player types to make the spell error`

  Freeze: string = `freeze both the boards to gain time for thinking and unlock it as soon as you remember the guess`

  Tag: string = `tag in the partner in place of you in the game `

  Draw: string = `offers the opponent to tag the partners in`

  Rewind: string = `tricks the opponent that results into rewinding the game `
  Series_and_Rules: string = `1. Dictionary to be set by one of the player
           in Locify on 1st round it is decided by toss 
           while in clash it is decided by the "Gods of NYG"
2. Once the Dictionary is set the Challenge is to be set by both the players to each other `


  result_calculation: string = `a. guessed word matches the challenge list
b. the word cannot be repeated each round 
c. for clash the word to be guessed under 10sec and under set time in Locify

`

  penalty: string = `Triggered upon decision timeout or game timeout
  note: for game timeout both the players get's the penalty
  note: max can be 1 on 2nd penalty the opponent is decided a winner
                Guidelines:
                   a.penalty on not setting up the toss side
                   b.penalty on not tossing the coin
                   c.penalty on not setting up the dictionary
                   d.penalty on not setting up the challenge
                   e. 1st penalty the guess time is deduced for current round
  `
  init: OCSFieldFormat[] = []
  ngOnInit(): void {

  }
  ngAfterViewInit(): void {
  }
  nexus: WritableSignal<string> = signal("PA_Y_S")
  bets: WritableSignal<RecBetItems> = signal(
    {
      firstCup: "NYG",
      secondCup: "NEXT",
      thirdCup: "GEN"
    }
  )
  constructor(private fb: FormBuilder) {

    this.init = [
      { OCStemplate: "PENALTY", OCSformControl: "", OCSToolTip: this.penalty },
      { OCStemplate: "RESULTS", OCSformControl: "", OCSToolTip: this.result_calculation },
      { OCStemplate: "RULES", OCSformControl: "", OCSToolTip: this.Series_and_Rules },


    ]
  }
  credits: string = '/logo/nyg-logo-4.webp'
}

