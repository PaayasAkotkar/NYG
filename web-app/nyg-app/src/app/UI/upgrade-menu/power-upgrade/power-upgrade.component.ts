import { Component, ElementRef, Inject, Input, OnChanges, OnInit, PLATFORM_ID, QueryList, SimpleChanges, ViewChild, viewChild, ViewChildren } from '@angular/core';
import { OnUpgradeColorDirective } from './on-upgrade-color.directive';
import { isPlatformBrowser } from '@angular/common';
import { powerAnimation } from './animation/animation';
import { interval, timer } from 'rxjs';
import { RemoveClassDirective } from "../../power-up/power-up-pattern/remove-class.directive";

@Component({
  selector: 'app-power-upgrade',
  imports: [OnUpgradeColorDirective, RemoveClassDirective],
  animations: [powerAnimation],
  templateUrl: './power-upgrade.component.html',
  styleUrls: ['./touchup/power-upgrade.scss', './touchup/color.scss',
    './touchup/animations.scss'
  ],
})
export class PowerUpgradeComponent implements OnChanges, OnInit {
  @Input() OCSpower: string = "NEXUS"
  @Input({ required: true }) OCSdonatedSpurs: number = 45
  OCSlvl: string | number = "S"
  @Input() activate: 'open' | 'close' = 'open'
  @ViewChildren('slot') slot!: QueryList<ElementRef<HTMLDivElement>>
  @Input() RemoveClickAnimation: boolean = false

  ngOnInit(): void {
    this.update()
  }

  animate() {

  }

  update() {
    if (this.slot)
      switch (true) {
        case this.OCSdonatedSpurs <= 40:
          var from = '#CED0D0'
          var to = '#F0ECF0'
          var defaultFrom = '#C2BD00'
          var defaultTo = '#1AD6C2'
          var percentage = `${this.OCSdonatedSpurs}%`
          var percentage2 = `${this.OCSdonatedSpurs / 2}%`
          var fromPer = '--fromPer'
          var toPer = '--toPer'
          // this.slot.forEach((e: ElementRef<HTMLDivElement>) => {
          //   const c = e.nativeElement
          //   c.style.setProperty('--from', from)
          //   c.style.setProperty('--to', to)
          // })

          var only = this.OCSdonatedSpurs / 5
          this.slot.toArray().slice(0, only).forEach((el: ElementRef<HTMLDivElement>) => {
            const see = el.nativeElement
            see.style.setProperty('--from', defaultFrom)
            see.style.setProperty('--to', from)
            see.style.setProperty(fromPer, percentage)
            see.style.setProperty(toPer, percentage2)
          })

          this.slot.toArray().slice(only, this.slot.length).forEach((el: ElementRef<HTMLDivElement>) => {
            const see = el.nativeElement
            see.style.setProperty('--from', from)
            see.style.setProperty('--to', to)
            see.style.setProperty(fromPer, percentage)
            see.style.setProperty(toPer, percentage2)
          })
          break

        case this.OCSdonatedSpurs > 40 && this.OCSdonatedSpurs < 75 && this.OCSdonatedSpurs != 100:
          var from = '#CED0D0'
          var to = '#F0ECF0'
          var defaultFrom = '#C2BD00'
          var defaultTo = '#1AD6C2'
          var percentage = `${this.OCSdonatedSpurs}%`
          var percentage2 = `${this.OCSdonatedSpurs / 2}%`
          var fromPer = '--fromPer'
          var toPer = '--toPer'

          var only = this.OCSdonatedSpurs / 5
          this.slot.toArray().slice(0, only).forEach((el: ElementRef<HTMLDivElement>) => {
            const see = el.nativeElement
            see.style.setProperty('--from', defaultTo)
            see.style.setProperty('--to', to)
            see.style.setProperty(fromPer, percentage)
            see.style.setProperty(toPer, percentage2)
          })

          this.slot.toArray().slice(only, this.slot.length).forEach((el: ElementRef<HTMLDivElement>) => {
            const see = el.nativeElement
            see.style.setProperty('--from', from)
            see.style.setProperty('--to', to)
            see.style.setProperty(fromPer, percentage)
            see.style.setProperty(toPer, percentage2)
          })
          break

        case this.OCSdonatedSpurs >= 75 && this.OCSdonatedSpurs != 100:
          var from = '#CED0D0'
          var to = '#F0ECF0'
          var defaultFrom = '#C2BD00'
          var defaultTo = '#1AD6C2'
          var percentage = `${this.OCSdonatedSpurs}%`
          var percentage2 = `${this.OCSdonatedSpurs}%`
          var fromPer = '--fromPer'
          var toPer = '--toPer'

          var only = this.OCSdonatedSpurs / 5
          this.slot.toArray().slice(0, only).forEach((el: ElementRef<HTMLDivElement>) => {
            const see = el.nativeElement
            see.style.setProperty('--from', defaultTo)
            see.style.setProperty('--to', defaultFrom)
            see.style.setProperty(fromPer, percentage)
            see.style.setProperty(toPer, percentage2)
          })

          this.slot.toArray().slice(only, this.slot.length).forEach((el: ElementRef<HTMLDivElement>) => {
            const see = el.nativeElement
            see.style.setProperty('--from', from)
            see.style.setProperty('--to', to)
            see.style.setProperty(fromPer, percentage)
            see.style.setProperty(toPer, percentage2)
          })
          break

        case this.OCSdonatedSpurs === 100:
          var from = '#CED0D0'
          var to = '#F0ECF0'
          var defaultFrom = '#C2BD00'
          var defaultTo = '#1AD6C2'
          var percentage = `100%`
          var percentage2 = `100%`
          var fromPer = '--fromPer'
          var toPer = '--toPer'

          this.slot.toArray().forEach((el: ElementRef<HTMLDivElement>) => {
            const see = el.nativeElement
            see.style.setProperty('--from', defaultFrom)
            see.style.setProperty('--to', defaultTo)
            see.style.setProperty(fromPer, percentage)
            see.style.setProperty(toPer, percentage2)
          })


      }

  }
  constructor(@Inject(PLATFORM_ID) private pID: Object) { }
  ngOnChanges(changes: SimpleChanges): void {
    if (isPlatformBrowser(this.pID)) {
      if (this.OCSdonatedSpurs < 60) {
        this.OCSlvl = 1
      } else if (this.OCSdonatedSpurs < 99) {
        this.OCSlvl = 2
      } else { this.OCSlvl = 'S' }
    }
  }
}
