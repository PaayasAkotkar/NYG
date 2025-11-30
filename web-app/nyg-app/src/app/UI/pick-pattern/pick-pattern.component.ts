import { Component, effect, EventEmitter, Input, OnChanges, OnInit, Output, signal, SimpleChanges, WritableSignal } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { HAND } from '../../helpers/hand';
import { OcsBtnsComponent } from "../ocs-btns/ocs-btns.component";
import { OCSFieldFormat } from '../../helpers/patterns';
import { OnEnter } from '../../animations/enter-leave';
import { FieldComponent } from '../../field/field/field.component';
import { RoomService } from '../../roomService/room.service';
export interface IMyView {
  position: number
  scrollHappend: boolean
  selectedToken: string | any
}
@Component({
  animations: [OnEnter],
  selector: 'app-pick-pattern',
  imports: [ReactiveFormsModule, FieldComponent, OcsBtnsComponent],
  templateUrl: './pick-pattern.component.html',
  styleUrls: ['./touchup/animation.scss', './touchup/pick-pattern.scss']
})
export class PickPatternComponent implements OnInit {
  @Input({ required: true }) OCSformGroup: FormGroup
  @Input({ required: true }) OCScollections: WritableSignal<OCSFieldFormat[]> = signal([])
  @Input({ required: true }) OCSheader: string = ""
  @Input({ required: true }) OCSbtnHeader: string = ""
  @Input({ required: true }) OCSmyTeam: string = ""
  @Input({ required: true }) OCSproceed: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGMode: string = ''
  @Output() OCSbtnEvent = new EventEmitter()
  @Output() OCSgetSelectedToken = new EventEmitter<any>()
  @Output() OCSposition = new EventEmitter<number>()
  @Output() OCSview = new EventEmitter<IMyView>()

  SelectedItem = new HAND.Queue()
  collectionKeys: string[] = []

  Click() {
    this.OCSbtnEvent.emit()
  }

  FillControls() {
    this.OCScollections().forEach(item => {
      this.OCSformGroup.addControl<any>(item.OCStemplate, this.FB.control(item.OCSformControl))
    })
  }

  ngOnInit(): void {
    this.FillControls()

    this.OCSformGroup.valueChanges.subscribe({
      next: () => {
        this.Update()
      }
    })
  }
  getPosition(pos: number) {
    this.OCSposition.emit(pos)
    this.OCSview.emit({ position: pos, scrollHappend: true, selectedToken: this.collectionKeys.pop() })
    console.log("positon: ", pos)
  }


  Token(token: string) {
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
      if (old && this.OCSformGroup.get(newly)?.value == true) {
        this.collectionKeys = this.collectionKeys.filter(e => e != old)
        this.collectionKeys.forEach(e => {
          if (this.OCSformGroup.get(e)?.value == true) {
            this.OCSformGroup.get(old)?.setValue(false, { emitEvent: false })
          }
        })
      }
    }
    this.OCSgetSelectedToken.emit(this.collectionKeys.pop())
  }

  Update() {
    const formValue: any = this.OCSformGroup.value;
    const value = Object.keys(formValue).filter(key => formValue[key] === true).join(' ')
    this.SelectedItem.enqueue(value)

    Object.keys(formValue).filter(key => {
      this.OCSformGroup.get(key)?.valueChanges.subscribe({
        next: (token) => {
          if (token) {

            Object.keys(this.OCSformGroup.controls).forEach(otherKey => {

              // only uncheck the previous selected value of book
              if (otherKey !== key) {
                this.OCSformGroup.get(otherKey)?.setValue(false, { emitEvent: false });
              }

            });
          }
        }
      })
    })

  }

  constructor(private FB: FormBuilder, private room: RoomService) {
    effect(() => {
      if (this.OCSproceed()) {
        this.FillControls()
      }
    })
  }

}
