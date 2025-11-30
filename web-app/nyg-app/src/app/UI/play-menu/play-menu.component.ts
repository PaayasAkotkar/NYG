import { Component, EventEmitter, Output, ViewEncapsulation } from '@angular/core';
import { Router } from '@angular/router';
import { RoomService } from '../../roomService/room.service';

@Component({
  selector: 'app-play-menu',
  imports: [],
  templateUrl: './play-menu.component.html',
  styleUrls: ['./touchup/play-menu.scss', './touchup/color.scss'],
  encapsulation: ViewEncapsulation.Emulated

})
export class PlayMenuComponent {
  CUSTOM = ['HOST'] // instead of custom
  QUICK = ['SURF'] // instead of quick match
  CREATE = ['MAKE'] // instead of create
  JOIN = ['JOIN']

  @Output() toC = new EventEmitter<boolean>
  @Output() toJ = new EventEmitter<boolean>
  constructor(private router: Router, private room: RoomService) { }

  ngOnInit(): void {
    const __ref = [this.CUSTOM]
    this.CUSTOM = [...this.CUSTOM.toLocaleString()]
    this.QUICK = [...this.QUICK.toLocaleString()]
    this.CREATE = [...this.CREATE.toLocaleString()]
    this.JOIN = [...this.JOIN.toLocaleString()]
  }

  toRoom() {
    this.router.navigateByUrl('/Make')
    this.toC.emit(true)
  }

  toJoin() {
    this.toJ.emit(true)
    this.router.navigateByUrl('/Join')
  }
}
