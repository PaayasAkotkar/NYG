import { Overlay, OverlayConfig, OverlayRef } from '@angular/cdk/overlay';
import { Injectable } from '@angular/core';
import { TagsComponent } from './tags.component';
import { ComponentPortal } from '@angular/cdk/portal';
import { Subscription, take, timer } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class TagsService {
  private ____OR: OverlayRef | null
  private channel = new Subscription()
  constructor(private overaly: Overlay) { }
  Setup(subTitle: string, title: string, mode: string, forTime: number = 3000) {
    this.closeSetup(); // stop rendering
    const config = new OverlayConfig({
      positionStrategy: this.overaly.position().global().centerHorizontally().centerVertically(),
      hasBackdrop: true,
      backdropClass: "TAGG",
    })
    this.____OR = this.overaly.create(config)
    const portal = new ComponentPortal(TagsComponent, null,)
    const __ref = this.____OR.attach(portal)

    __ref.instance.NYGsubTitle = subTitle
    __ref.instance.NYGtitle = title
    __ref.instance.NYGMode = mode
    // new: we'll handle it in the request
    // only invoke for 3sec => 3000
    // this.channel = timer(forTime).pipe(take(1)).subscribe(
    //   {
    //     next: () => {
    //       this.closeSetup()
    //     }
    //   }
    // )

    return this.____OR
  }
  closeSetup() {
    if (this.____OR) {
      this.____OR.dispose()
      this.____OR = null
    }
  }
}
