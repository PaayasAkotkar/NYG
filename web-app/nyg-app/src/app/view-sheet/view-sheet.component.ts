import { Component, ElementRef, Input, OnInit, viewChild, ViewChild, ViewEncapsulation } from '@angular/core';
import { StatboardComponent } from "../statboard/statboard.component";
import { CdkAccordionModule } from '@angular/cdk/accordion';
import { Overlay, OverlayModule } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { ViewSheetService } from './view-sheet.service';
@Component({
  selector: 'app-view-sheet',
  imports: [CdkAccordionModule, OverlayModule],
  templateUrl: './view-sheet.component.html',
  styleUrls: ['./touchup/view-sheet.scss',],
  encapsulation: ViewEncapsulation.Emulated
})
export class ViewSheetComponent implements OnInit {
  @Input() playerName: string = "KING"
  @Input() teamName: any
  @Input() rating: string = "00"
  @Input() updateStatsSheet: Map<string, Map<string, string>> = new Map([])
  @ViewChild('origin', { static: false }) origin: ElementRef
  round: string[] = ['R1', 'R2']
  set: string[] = ['S1', 'S2', 'S3']
  isOpen: boolean = true
  show: boolean = false

  constructor(private overlay: Overlay, private service: ViewSheetService) { }

  toggle() {
    this.show = !this.show
    if (this.show) {
      this.show = false
      this.service.Setup(
        this.rating,
        this.playerName,
        this.teamName,
      )
    } else {
      this.service.closeSetup()
    }

  }
  ngOnInit(): void { }
}
