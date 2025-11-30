import { isPlatformBrowser } from '@angular/common';
import { Inject, Injectable, PLATFORM_ID } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class CookieService {

  constructor(@Inject(PLATFORM_ID) private pID: Object) { }

  getCookie(token: string) {
    try {
      if (isPlatformBrowser(this.pID)) {
        var name = token + '='
        const cookies = document.cookie.split(';')
        for (var c of cookies) {
          if (c.startsWith(name)) {
            return decodeURIComponent(c.substring(name.length))
          }
        }
      }
    } catch (e) {
      console.error('cookie error: ', e)
    }
    return null
  }

}
