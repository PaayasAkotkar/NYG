import { HttpClient } from "@angular/common/http";
import { UpdateLivelyVoucher } from "../voucher/update-livelyVouchers";
import { IDeck, IIMG, IName, INickname, } from "../../patches/patch";
export interface Credits {
    coins: string
    spurs: string
    boardTheme: string
    guessTheme: string
    id: string

}
export class UpdateLivelyResource extends UpdateLivelyVoucher {
    constructor(private id: string, private http: HttpClient,) { super() }

    UpdateDeck(first: string, second: string, isDefault: boolean) {
        const token: IDeck = {
            _first: first, _second: second, isDefault: isDefault, id: this.id
        }
        this.http.put(this.updateDeck, token).subscribe()
    }
    UpdateImage(img: FormData) {
        const token: IIMG = {
            id: this.id, image: img,
        }
        this.http.patch(this.updateImg, token).subscribe()
    }
    UpdateName(name: string) {
        const token: IName = {
            name: name, id: this.id
        }
        this.http.patch(this.updateName, token).subscribe()
    }
    UpdateNickname(nickname: string) {
        const token: INickname = {
            nickname: nickname, id: this.id
        }
        this.http.patch(this.updateNickname, token).subscribe()
    }
    UpdateCoins(coins: Credits) {
        coins.id = this.id
        this.http.patch(this.updateCoins, coins).subscribe()
    }
    UpdateSpurs(spurs: Credits) {
        spurs.id = this.id
        this.http.patch(this.updateSpurs, spurs).subscribe()
    }
    UpdateCredits(credits: Credits) {
        credits.id = this.id
        this.http.patch(this.updateCoins, credits).subscribe()
    }
    UpdateTheme(boardTheme: string, guessTheme: string) {
        var token: Credits = {
            id: this.id,
            boardTheme: boardTheme,
            guessTheme: guessTheme,
            coins: "",
            spurs: ""
        }
        this.http.patch(this.updateTheme, token).subscribe()

    }
}