import { isPlatformBrowser } from '@angular/common';
import { Inject, Injectable, OnInit, PLATFORM_ID, signal, WritableSignal } from '@angular/core';
import { DomSanitizer, SafeUrl } from '@angular/platform-browser';
import { resolve } from 'path';
import { BehaviorSubject, Observable } from 'rxjs';
export interface _F {
  FileName: string
  Base64: SafeUrl
  PostBase64: string
  FileSrc: string
  File: File
  FileType: string
}
@Injectable({
  providedIn: 'root'
})
export class ImageService implements OnInit {
  public imageSrc: WritableSignal<string> = signal("")
  private recieveToken: BehaviorSubject<_F> = new BehaviorSubject<_F>({
    FileName: "", Base64: {}, FileSrc: "", File: new File([], ""),
    PostBase64: "", FileType: "",
  })

  constructor(@Inject(PLATFORM_ID) private pid: Object,
    private by: DomSanitizer) { }

  ngOnInit(): void {
    this.Token().subscribe({
      next: (token) => {
        this.imageSrc.set(token.FileSrc)
      }
    })
  }
  // ConvertBase64 returns the safeurl
  // note: later you have to prefix it data:.....
  ConvertBase64(src: ArrayBuffer): SafeUrl {
    const bytes = new Uint8Array(src)
    let binary = ''
    bytes.forEach(b => binary += String.fromCharCode(b))
    return binary
  }

  // SetFile creates file url for given file 
  // note: it doesnt monitors
  SetFile(fileName: File): void {
    var read = new FileReader()
    read.onload = (e) => {
      if (isPlatformBrowser(this.pid)) {
        var _r = e.target?.result as ArrayBuffer
        const bytes = new Uint8Array(_r)
        let binary = ''
        bytes.forEach((b) => {
          binary += String.fromCharCode(b)
        })
        // data:fileName.type;base64,+safeurl
        // this is how the encoding must be if you are going to set the img
        const base64 = "data:" + fileName.type + ";base64," + btoa(binary)
        var a = this.by.bypassSecurityTrustUrl(base64)
        var _f = e.target?.result as string
        var p = _f.split(",")[1]
        var f: _F = {
          File: fileName,
          Base64: a,
          FileSrc: _f,
          FileName: fileName.name,
          PostBase64: p,
          FileType: fileName.type
        }

        this.recieveToken.next(f)
      }
    }
    read.readAsDataURL(fileName)
    read.onabort = () => {
      console.error("file abort")
    }
    read.onerror = (e) => {
      console.log('file error', e)
    }
    read.onloadend = () => {
      console.log("file loaded")
    }
    read.onprogress = () => {
      console.log("file progressing")
    }
  }

  // Token monitors on file 
  Token(): Observable<_F> {
    return this.recieveToken.asObservable()
  }

}
