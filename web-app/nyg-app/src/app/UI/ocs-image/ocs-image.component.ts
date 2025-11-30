import { EventEmitter, AfterViewInit, Component, effect, ElementRef, Input, OnInit, Output, signal, ViewChild, WritableSignal } from '@angular/core';
import { ImageService } from './image.service';
import { DomSanitizer, SafeUrl } from '@angular/platform-browser';

@Component({
  selector: 'app-ocs-image',
  imports: [],
  templateUrl: './ocs-image.component.html',
  styleUrls: ['./touchup/ocs-image.scss']
})
export class OcsImageComponent implements OnInit {
  constructor(private image: ImageService, private sen: DomSanitizer) { }

  @Input({ required: true }) NYGsrc: WritableSignal<string> = signal('')
  @Input({ required: true }) NYGenableEvent: boolean = false

  @Output() OCSEVENT = new EventEmitter<File>()
  @Output() OCSClick = new EventEmitter<void>()
  ngOnInit(): void {
  }

  Register(ev: Event) {
    const input = ev.target as HTMLInputElement
    if (input.files) {
      const file = input.files[0]
      this.OCSEVENT.emit(input.files[0])
    }
  }
  click() {
    this.OCSClick.emit()
  }
}
