import { Component, Input, OnInit } from '@angular/core';
import { OnProgressDirective } from "../daily-events/on-progress.directive";
import { IDailyEvents } from '../daily-events/daily-events.component';

@Component({
  selector: 'app-progress-bar',
  imports: [OnProgressDirective],
  templateUrl: './progress-bar.component.html',
  styleUrls: ['./touchup/animation.scss', './touchup/color.scss', './touchup/pb.scss'],
})
export class ProgressBarComponent implements OnInit {
  defaultPic: string = "/default/default-pic.png"

  @Input() OCSimgSrc: string = ""
  @Input() OCSProgress: number = 1000
  ngOnInit(): void {
    if (this.OCSimgSrc == "") {
      this.OCSimgSrc = this.defaultPic
    }
  }
}
