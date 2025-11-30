import { Component, Input, ViewEncapsulation } from '@angular/core';
import { TagsComponent } from "../tags/tags.component";

@Component({
  selector: 'app-room-name',
  imports: [TagsComponent],
  templateUrl: './room-name.component.html',
  encapsulation: ViewEncapsulation.Emulated
})
export class RoomNameComponent {
  @Input() NYGroomName: string = "Room"
  roomname: string = "ROOM"
}
