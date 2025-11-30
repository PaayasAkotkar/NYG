import { Component, input, Input, OnInit, signal, WritableSignal } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { RoomService } from '../../../roomService/room.service';

import { OCSFieldFormat } from '../../../helpers/patterns';
import { OcsBtnsComponent } from "../../ocs-btns/ocs-btns.component";
// todo make sure work on click events
// here the first click value is awalys the last click even if you click on another value
@Component({
  selector: 'app-power-draw',
  imports: [ReactiveFormsModule, OcsBtnsComponent],
  templateUrl: './power-draw.component.html',
  styleUrls: ['./touchup/draw.scss', './touchup/color.scss', './touchup/animations.scss']
})
export class PowerDrawComponent implements OnInit {
  constructor(private room: RoomService) { }

  @Input({ required: false }) NYGmyteam: string = ""
  @Input({ required: false }) NYGmyRoom: string = ""
  collectionKeys: string[] = []
  NYGdrawOfferForm = new FormBuilder().group({
    ACCEPT: [false],
    DECLINE: [false]
  })

  NYGcollection: OCSFieldFormat[] = [
    { OCStemplate: 'ACCEPT', OCSformControl: false, OCSToolTip: "" },
    { OCStemplate: 'DECLINE', OCSformControl: false, OCSToolTip: "" },
  ]
  decision: WritableSignal<boolean> = signal(false)

  ngOnInit(): void {
    this.NYGdrawOfferForm.valueChanges.subscribe({
      error: (err) => console.log(err)
    })

  }

  Update(token: string) {
    this.collectionKeys.push(token)
    const len = this.collectionKeys.length

    if (len > 1) {
      const newly = this.collectionKeys[len - 1]

      var temp
      if (newly) {
        temp = this.collectionKeys[this.collectionKeys.indexOf(newly) - 1]

      }
      const old = temp

      // this condition prevents the newly and old value to get false  
      if (old && this.NYGdrawOfferForm.get(newly)?.value == true) {
        this.collectionKeys = this.collectionKeys.filter(e => e != old)
        this.collectionKeys.forEach(e => {
          if (this.NYGdrawOfferForm.get(e)?.value == true) {
            this.NYGdrawOfferForm.get(old)?.setValue(false, { emitEvent: false })
          }
        })
      }
    }
  }
  NYGref = input(false)
  Confirm() {
    if (!this.NYGref) {
      if (this.collectionKeys.pop() == 'ACCEPT')
        this.decision.set(true)
      else {
        this.decision.set(false)
      }
      this.room.SendDrawOffer(this.decision(), this.NYGmyRoom)
    }
  }

}