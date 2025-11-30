import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable, OnDestroy } from '@angular/core';
import { BehaviorSubject, catchError, finalize, from, interval, Observable, of, repeat, retry, Subject, Subscriber, Subscription, switchMap, takeUntil, tap, throwError, timeout, timer } from 'rxjs';
import { fromFetch } from 'rxjs/fetch';
import { v4 } from 'uuid';

interface Registeration {
  roomID: string
  nickName: string
  id: string
  token: string
}
@Injectable({
  providedIn: 'root'
})
export class LongPollingService implements OnDestroy {
  constructor(private http: HttpClient) { }
  private id = v4()
  private stopPolling = new Subject<void>()
  private recieveToken: Subject<string> = new Subject<string>()
  private URLs: string[] = [
    `http://localhost:9200/post-chat/${this.id}`,
    `http://localhost:9200/get-chat/poll/${this.id}`
  ]
  Sub: Subscription = new Subscription()
  private signUp: Registeration = { id: this.id, token: "", roomID: "", nickName: "" }
  ngOnDestroy(): void {
    this.stopPolling.next()
    this.stopPolling.complete()
  }
  Stop() {
    this.stopPolling.next();
    this.stopPolling.complete();
    this.Sub.unsubscribe();

    console.log(`[PollService] Polling stopped for ID ${this.id}.`);

  }

  Start() {

    return timer(0, 500).pipe(
      switchMap(() => this.Awake()),

      catchError((error) => {
        console.error(`polling error 👾 ${this.id}, stopping polling:`, error)
        this.Stop()
        return throwError(() => error)
      }),
      repeat(),
      finalize(() => {
        console.log(` finalize 🥉: ${this.id}`);
        this.Stop()
      }),
      takeUntil(this.stopPolling),

    )
  }

  Fetch(): Observable<string> {
    return this.recieveToken.asObservable()
  }

  Awake() {

    this.URLs[1] = `http://localhost:9200/get-chat/poll/${this.id}`;

    return this.http.get(this.URLs[1], { responseType: 'text' }).pipe(
      takeUntil(this.stopPolling),

      timeout(10000), // important match the client timeout

      tap((token) => {
        console.log("from url: ", this.URLs[1], "token:", token)
      }),
      catchError((err) => {
        console.log("from url: ", this.URLs[1])
        return throwError(() => err?.message || 'NETWORK ERROR');

      }),
    )
  }

  Register(roomID: string, nickName: string) {
    const sendToken: Registeration = {
      id: this.id, token: "", roomID: roomID,
      nickName: nickName
    }

    this.signUp = sendToken
    console.log("ulrs: ", this.URLs, "📫: ", sendToken)
    return this.http.post<Registeration>(this.URLs[0], sendToken).pipe(
      timeout(5000),
      catchError(err => {
        takeUntil(this.stopPolling)
        console.error("❌ registration request failed:", err);
        return throwError(() => err);
      })
    )
  }

  Publish(token: string) {
    var sendToken: Registeration = {
      id: this.signUp.id, token: token, roomID: this.signUp.roomID,
      nickName: this.signUp.nickName
    }

    return this.http.post<Registeration>(this.URLs[0], sendToken,).pipe(
    )
  }
}
