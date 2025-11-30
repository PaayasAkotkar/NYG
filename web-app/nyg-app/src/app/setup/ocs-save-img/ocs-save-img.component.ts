import { Component, EventEmitter, Input, OnChanges, OnInit, Output, signal, WritableSignal } from '@angular/core';
import { OcsImageComponent } from "../../UI/ocs-image/ocs-image.component";
import { _F, ImageService } from '../../UI/ocs-image/image.service';
import { Subscription } from 'rxjs';
import { DomSanitizer, SafeUrl } from '@angular/platform-browser';

@Component({
  selector: 'app-ocs-save-img',
  imports: [OcsImageComponent],
  templateUrl: './ocs-save-img.component.html',
  styleUrl: './ocs-save-img.component.scss'
})
export class OcsSaveImgComponent implements OnInit {
  constructor(private image: ImageService, private san: DomSanitizer) { }

  @Input({ required: true }) NYGsrc: WritableSignal<string> = signal('E:\project-nyg-web-app\nyg-app\src\app\UI\ocs-image\default\default-pic.png')
  defaultPic: string = "/default/default-pic.png"
  Token: WritableSignal<any> = signal('')
  Sub: Subscription = new Subscription()
  @Output() ListenFile = new EventEmitter<_F>()
  @Output() FormDataFile = new EventEmitter<File>()

  FileName: File

  Register(token: File) {
    this.FileName = token
    this.image.SetFile(this.FileName)

    this.FormDataFile.emit(token)
  }

  ngOnInit(): void {
    this.defaultPic = this.NYGsrc()
    this.image.Token().subscribe({
      next: (token) => {
        if (token.FileSrc != "") {
          this.NYGsrc.set(token.FileSrc)
          this.ListenFile.emit(token)
        } else {
          this.NYGsrc.set(this.defaultPic)
        }
      }
    })
  }

}
