import { isPlatformBrowser } from '@angular/common';
import { Inject, Injectable, PLATFORM_ID } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class SessionService {
  constructor(@Inject(PLATFORM_ID) private platformId: Object) { }

  setItem(key: string, item: string) {
    if (isPlatformBrowser(this.platformId)) {

      sessionStorage.setItem(key, item)

    }
  }
  getItem(key: string) {

    var item
    if (isPlatformBrowser(this.platformId)) {
      item = sessionStorage.getItem(key)
    }
    return item
  }
  getKey(key: string) {
    if (isPlatformBrowser(this.platformId)) {

      sessionStorage.getItem(key)

    }
  }
  removeItem(key: string) {
    if (isPlatformBrowser(this.platformId)) {

      sessionStorage.removeItem(key)

    }
  }
  getLength() {
    var length = -1
    if (isPlatformBrowser(this.platformId)) {
      length = sessionStorage.length
    }
    return length
  }
}
