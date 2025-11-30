import { HttpClient } from "@angular/common/http";
import { LoginFields } from "../voucher/loginVouchers";
import { IDeck } from "../../patches/patch";
import { SessionService } from "../../storage/session/session.service";
import { v4 } from "uuid";
import { timer } from "rxjs";
import { NgZone } from "@angular/core";

export interface LoginProfile {
    id: string
    name: string
    nickname: string
}

export interface IIMG {
    id: string
    image: FormData
}

export interface INickname {
    id: string
    nickname: string
}
export interface IName {
    id: string
    name: string
}
export class LoginResources extends LoginFields {
    constructor(private run: NgZone,
        private http: HttpClient, private id: string) { super() }
    Login(name: string, nickname: string) {
        try {

            const token: LoginProfile = { nickname: nickname, name: name, id: this.id }
            this.http.post(this.postURL, token).subscribe()
        } catch (e) {
            console.error("login error")
        }
    }

    // default: incase if the user doesnt login
    Init() {
        try {
            this.run.run(() => {
                var nickname = 'user_' + this.id
                var name = 'user_' + this.id
                const token: LoginProfile = { id: this.id, name: name, nickname: nickname, }
                this.http.post(this.postURL, token).subscribe()
            })

        } catch (e) {
            console.error('init error: ', e)
        }
    }

    GetEvents() {
        this.http.put(this.getEvents, { id: this.id }).subscribe({
            next: (token) => { },
            error: (err) => {
                console.error(err)
            }
        })
    }
}