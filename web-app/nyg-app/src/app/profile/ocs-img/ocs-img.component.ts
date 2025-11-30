import { isPlatformBrowser } from '@angular/common';
import { EventEmitter, AfterViewInit, ChangeDetectorRef, Component, ElementRef, Inject, Input, NgZone, OnInit, Output, PLATFORM_ID, signal, ViewChild, ViewEncapsulation, WritableSignal } from '@angular/core';
import { OcsImageComponent } from "../../UI/ocs-image/ocs-image.component";

@Component({
  selector: 'app-ocs-img',
  imports: [OcsImageComponent],
  templateUrl: './ocs-img.component.html',
  styleUrls: ['./touchup/ocs-color.scss', './touchup/ocs-img.scss'],
  encapsulation: ViewEncapsulation.Emulated
})
export class OcsImgComponent implements OnInit {
  @Input({ required: true }) OCSPicURL: WritableSignal<string> = signal("/default/default-pic.png")
  @Output() Click: EventEmitter<void> = new EventEmitter<void>()
  
  ngOnInit(): void { }
  listen() {
    this.Click.emit()
  }
}
