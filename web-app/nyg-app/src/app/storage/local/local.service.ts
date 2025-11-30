import { isPlatformBrowser } from '@angular/common';
import { Inject, Injectable, PLATFORM_ID } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class LocalService {

  constructor(@Inject(PLATFORM_ID) private platformId: Object) {
  }
  setItem(key: string, item: string) {
    if (isPlatformBrowser(this.platformId)) {

      localStorage.setItem(key, item)
    }

  }
  getItem(key: string) {
    var item: null | string = ""
    if (isPlatformBrowser(this.platformId)) {

      item = localStorage.getItem(key)

    }
    return item
  }
  getKey(key: string) {
    if (isPlatformBrowser(this.platformId)) {
      localStorage.getItem(key)

    }
  }
  removeItem(key: string) {
    if (isPlatformBrowser(this.platformId)) {
      localStorage.removeItem(key)
    }
  }
  getLength() {
    var length = -1
    if (isPlatformBrowser(this.platformId)) {
      length = localStorage.length
    }
    return length
  }

}
