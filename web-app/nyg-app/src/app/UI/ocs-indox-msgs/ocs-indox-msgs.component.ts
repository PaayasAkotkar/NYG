import { Component, Input, OnInit } from '@angular/core';
import { OcsIndoxMsgComponent } from "../ocs-indox-msg/ocs-indox-msg.component";
export interface IMessages {
  author: string
  msg: string
  view: boolean // if the message seen
  date: string
}
@Component({
  selector: 'app-ocs-indox-msgs',
  imports: [OcsIndoxMsgComponent],
  templateUrl: './ocs-indox-msgs.component.html',
  styleUrls: ['./touchup/box.scss', './touchup/color.scss',
    './touchup/animation.scss'
  ],
})
export class OcsIndoxMsgsComponent implements OnInit {
  @Input() OCSmsgs: IMessages[] = [
    { author: "king", msg: "wel come to game", view: true, date: "" },
    { author: "king", msg: "wel come to game", view: true, date: "" },
    { author: "king", msg: "wel come to game", view: true, date: "" },
    { author: "king", msg: "wel come to game", view: true, date: "" },
    { author: "king", msg: "wel come to game", view: true, date: "" },
    { author: "king", msg: "wel come to game", view: true, date: "" },

  ]
  ngOnInit(): void {

    var _default: IMessages = { author: "PAAYAS", msg: "welcome to NYG!!! 🤗", view: true, date: "09 Nov 2025" }
    this.OCSmsgs.push(_default)
    this.OCSmsgs = this.OCSmsgs.filter(e => !e.view)
  }
}
