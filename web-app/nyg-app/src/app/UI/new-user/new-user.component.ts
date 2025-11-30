import { Component, OnInit } from '@angular/core';
import { OcsInputComponent } from "../ocs-input/ocs-input.component";
import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { OcsBtnsComponent } from "../ocs-btns/ocs-btns.component";

import { LoginService } from '../../login-service/login.service';
import { LoginProfile } from '../../login-service/resources/loginResources';

@Component({
  selector: 'app-new-user',
  imports: [OcsInputComponent, ReactiveFormsModule, OcsBtnsComponent],
  templateUrl: './new-user.component.html',
  styleUrl: './new-user.component.scss'
})
export class NewUserComponent implements OnInit {
  loginForm: FormGroup
  constructor(private fb: FormBuilder, private login: LoginService) { }
  ngOnInit(): void {
    this.loginForm = this.fb.group({ name: [''], nickname: [''] })
  }
  Create() {
    var nickname = this.loginForm.get('nickname')?.value
    var name = this.loginForm.get('name')?.value
    this.login.Login(name, nickname)
  }
}
