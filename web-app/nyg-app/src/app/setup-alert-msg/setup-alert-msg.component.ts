import { Component, Input, OnDestroy, OnInit } from '@angular/core';
import { TagsService } from '../UI/tags/tags.service';

@Component({
  selector: 'app-setup-alert-msg',
  imports: [],
  templateUrl: './setup-alert-msg.component.html',
})
export class SetupAlertMsgComponent implements OnInit, OnDestroy {

  @Input({ required: true }) NYGalertFor: string = ""
  @Input({ required: true }) NYGalertMsg: string = ""
  @Input({ required: true }) NYGMode: string = ""

  constructor(private ts: TagsService) {
  }

  ngOnInit(): void {
    this.ts.Setup(this.NYGalertMsg, this.NYGalertFor, this.NYGMode)
  }

  ngOnDestroy(): void {
    this.ts.closeSetup()
  }
}
