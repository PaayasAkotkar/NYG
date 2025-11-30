import { AfterViewInit, Component, ElementRef, Inject, Injectable, Input, OnChanges, OnInit, PLATFORM_ID, QueryList, Renderer2, SimpleChanges, ViewChild, ViewChildren } from '@angular/core';
import { OnChanceDirective } from './on-chance.directive';
import { CommonModule, isPlatformBrowser } from '@angular/common';
import { OnTeamColorDirective } from "../game-tickets/on-team-color.directive";

@Component({
  selector: 'app-bar',
  imports: [OnChanceDirective, CommonModule, OnTeamColorDirective],
  templateUrl: './bar.component.html',
  styleUrls: ['./touchup/bar.scss', './touchup/color.scss', './touchup/animations.scss'],
})
export class BarComponent implements OnInit, OnChanges, AfterViewInit {
  @Input() PlayersNames: string[] = []
  @ViewChildren('slot') slot!: QueryList<ElementRef<HTMLDivElement>>
  fromColors = ['#bd0000', '#ff9000', '#ffc30e', '#96cd21']
  toColors = ['white']
  @Input({ required: true }) NYGheader: string = "KING"
  @Input({ required: true }) OCScurrentChance: number = 9
  @Input({ required: true }) NYGTeamName: string = ''

  getDynamicColor(index: number, type: 'from' | 'to'): string {
    const colorArray = type === 'from' ? this.fromColors : this.toColors;
    return colorArray[index % colorArray.length];
  }
  constructor(@Inject(PLATFORM_ID) private pID: Object, private renderer: Renderer2) { }
  ngAfterViewInit(): void {

  }
  ngOnChanges(changes: SimpleChanges): void {

  }
  animate() {
    if (this.slot) {

      var to: string = '#45556B'
      var from: string = '#f5f4de'
      var fromPer = '--fromPer'
      var toPer = '--toPer'
      var percentage = `${this.OCScurrentChance}%`
      var percentage2 = `${this.OCScurrentChance / 2}%`

      // min=0,max=8
      // currentChances->min=10,max=0
      // currentChances>max->currentChances-2
      // c->for 10->8, 9->8-1, 8->8-2,7->8-1,6->8-1, 7->8-1, 6->6-1,5->5-1,

      // in-order to remove from left->right
      this.slot.toArray().slice(this.OCScurrentChance).forEach((el: ElementRef<HTMLDivElement>) => {
        const see = el.nativeElement
        see.style.setProperty('--from', from)
        see.style.setProperty('--to', to)
      })

      // switch (this.OCScurrentChance) {
      //   case 3: this.slot.toArray().slice(0, 3).forEach((el: ElementRef<HTMLDivElement>) => {
      //     const see = el.nativeElement
      //     see.style.setProperty('--from', from)
      //     see.style.setProperty('--to', to)
      //   })
      //     break
      //   case 7: this.slot.toArray().slice(0, this.OCScurrentChance).forEach((el: ElementRef<HTMLDivElement>) => {
      //     const see = el.nativeElement
      //     see.style.setProperty('--from', from)
      //     see.style.setProperty('--to', to)
      //   })
      //     break
      //   case 1: this.slot.toArray().slice(0, 7).forEach((el: ElementRef<HTMLDivElement>) => {
      //     const see = el.nativeElement
      //     see.style.setProperty('--from', from)
      //     see.style.setProperty('--to', to)
      //   })
      //     break
      //   case 0:
      //     this.slot.toArray().forEach((el: ElementRef<HTMLDivElement>) => {
      //       const see = el.nativeElement
      //       see.style.setProperty('--from', from)
      //       see.style.setProperty('--to', to)
      //     })
      //     break
      // }
    }
  }
  ngOnInit(): void {
    this.animate()
  }
}
