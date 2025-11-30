import { AfterViewInit, Component, ElementRef, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { Subject, takeUntil, timer } from 'rxjs';

@Component({
  selector: 'app-load-logo',
  imports: [],
  templateUrl: './load-logo.component.html',
  styleUrls: ['./touchup/load.scss', './touchup/animations.scss', './touchup/color.scss']
})
export class LoadLogoComponent implements OnInit {
  @ViewChild('ocs_logo') LoadLogo!: ElementRef<HTMLDivElement>
  src = "/logo/nyg-logo.webp"
  srcLeft = "/logo/left-head.webp"
  srcRight = "/logo/right-head.webp"

  ngOnInit(): void {

  }
}
