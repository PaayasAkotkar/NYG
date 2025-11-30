import { Component, Input, ViewEncapsulation } from '@angular/core';
import { TagsComponent } from "../tags/tags.component";

@Component({
  selector: 'app-friend-code',
  imports: [TagsComponent],
  templateUrl: './friend-code.component.html',
  styleUrl: './friend-code.component.scss',
  encapsulation: ViewEncapsulation.Emulated,
})
export class FriendCodeComponent {
  @Input() NYGfriendCode: string = "NYG"
  CODE: string = "CODE"
}
