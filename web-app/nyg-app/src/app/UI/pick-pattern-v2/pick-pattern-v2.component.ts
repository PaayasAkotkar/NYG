import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { HAND } from '../../helpers/hand';

import { OCSFieldFormat } from '../../helpers/patterns';


@Component({
  selector: 'app-pick-pattern-v21',
  imports: [ReactiveFormsModule],
  templateUrl: './pick-pattern-v2.component.html',
})
export class PickPatternV2Component implements OnInit {
  @Input() OCSformGroup: FormGroup
  @Input({ required: true }) OCSheader: string
  @Input({ required: true }) OCScollections: OCSFieldFormat[]
  @Output() OCSgetBody = new EventEmitter<string>()
  @Input({ required: true }) OCSbtnHeader: string
  @Input({ required: true }) OCSmyTeam: string
  @Output() OCSbtn = new EventEmitter<any>()

  TEST2() {
    this.OCSbtn.emit()
  }

  TEST() {
    this.OCSgetBody.emit(this.SelectedItem.front())
  }

  FillControls() {
    this.OCScollections.forEach(item => {
      this.OCSformGroup.addControl<any>(item.OCStemplate, this.FB.control(item.OCSformControl))
    })
  }
  constructor(private FB: FormBuilder) { }

  ngOnInit(): void {
    this.FillControls()
    this.OCSformGroup.valueChanges.subscribe({
      next: (token: any) => {
        this.Update()
      }
    })
  }


  SelectedItem = new HAND.Queue()

  Update() {
    const formValue: any = this.OCSformGroup.value;
    const value = Object.keys(formValue).filter(key => formValue[key] === true).join(' ')
    this.SelectedItem.enqueue(value)
    this.SelectedItem.store = this.SelectedItem.Range().filter(e => !/\s\S/.test(e))

    Object.keys(formValue).filter(key => {
      this.OCSformGroup.get(key)?.valueChanges.subscribe({
        next: (token) => {
          if (token) {

            Object.keys(this.OCSformGroup.controls).forEach(otherKey => {

              // only uncheck the previous selected value of book
              if (otherKey !== key) {
                this.OCSformGroup.get(otherKey)?.setValue(false, { emitEvent: false });
              }
            })
          }
        }
      })
    })

  }
}
