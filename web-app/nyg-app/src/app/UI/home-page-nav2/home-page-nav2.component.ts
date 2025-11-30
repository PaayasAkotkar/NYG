import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';

@Component({
  selector: 'app-home-page-nav2',
  imports: [],
  templateUrl: './home-page-nav2.component.html',
  styleUrls: ['./touchup/color.scss', './touchup/ocs-nme.scss', './touchup/animation.scss'],

})
export class HomePageNav2Component implements OnInit {
  @Input({ required: true }) OCSheader: string
  @Input({ required: true }) OCSsubHeader: string
  @Input({ required: true }) OCSsubHeader2: string
  default: string = 'tier/grand msater.webp'
  @Input({ required: true }) OCSsrc: string = ""
  @Output() NYGListen = new EventEmitter<void>()
  Listen() {
    this.NYGListen.emit()
  }
  ngOnInit(): void {

  }
}
