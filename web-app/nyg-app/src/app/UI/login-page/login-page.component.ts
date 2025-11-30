import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { LoginService } from '../../login-service/login.service';
import { LoginProfile } from '../../login-service/resources/loginResources';
import { v4 } from 'uuid';

@Component({
  selector: 'app-login-page',
  imports: [ReactiveFormsModule,],
  templateUrl: './login-page.component.html',
  styleUrls: ['./touchup/login.scss', './touchup/color.scss', './touchup/animations.scss'],
})
export class LoginPageComponent implements OnInit {
  _login: FormGroup
  ngOnInit(): void {
    this._login = this.fb.group({
      xx0s: [''], // for username
      yyss89: [''] // for password
    })
  }
  login() {
    var id = v4()
    const AppendData: FormData = new FormData()
    AppendData.append("img", new Blob())

    var x: LoginProfile = {
      name: "kingp",
      nickname: "xcx",
      id: id,
    }
    this.l.Login(x.name, x.nickname)
  }
  constructor(private fb: FormBuilder, private l: LoginService) { }
}
