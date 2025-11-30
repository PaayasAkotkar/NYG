import { Observable } from "rxjs";
import { GetProfilePattern, PostProfilePattern, ProfileFields } from "../vouchers/profileVoucher";
import { Inject, NgZone, PLATFORM_ID } from "@angular/core";
import { isPlatformBrowser } from "@angular/common";
import { HttpClient } from "@angular/common/http";
import { IIMG } from "../../setup/setup.component";

export class ProfileMethods extends ProfileFields {

    constructor(private id: string, private http: HttpClient, private zone: NgZone, @Inject(PLATFORM_ID) protected platformID: Object) {
        super()
    }
    Connect() {
        try {
            if (isPlatformBrowser(this.platformID)) {
                if (this.Conn && this.Conn.readyState == EventSource.CONNECTING) {
                    return
                }
                this.Conn = new EventSource(this.URLs[0])

                this.Conn.onopen = () => { }

                this.Conn.onmessage = (token: MessageEvent) => {
                    this.zone.run(() => {
                        const conv = JSON.parse(token.data)
                        this.receivedTokens.next(conv)
                        this.Conn.close()
                    })
                }

                this.Conn.onerror = (error: Event) => {
                }
            }
        } catch (e) {
            console.error("error 🚦: ", e)
        }
    }

    Get(): Observable<GetProfilePattern> {
        return this.receivedTokens.asObservable()
    }

    Register(data: FormData): Observable<FormData> {
        data.delete("id")
        data.append("id", this.id)
        return this.http.patch<FormData>(this.URLs[0], data)
    }
}