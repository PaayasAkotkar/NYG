import { Component, NgZone, OnChanges, OnInit, signal, SimpleChanges, WritableSignal } from '@angular/core';
import { CreateRoomComponent } from "../UI/create-room/create-room.component";
import { filter, Observable, take, timer } from 'rxjs';
import { NavigationEnd, Router } from '@angular/router';
import { OCSFieldFormat } from '../helpers/patterns';
import { NYGqService } from '../graphql/nygq.service';
import { gql } from 'apollo-angular';

@Component({
  selector: 'app-make-room',
  imports: [CreateRoomComponent],
  templateUrl: './make-room.component.html',
})
export class MakeRoomComponent implements OnInit {

  SportsBooks: WritableSignal<OCSFieldFormat[]> = signal([
  ])
  EntertainmentBooks: WritableSignal<OCSFieldFormat[]> = signal([])

  constructor(private books: NYGqService) {

  }

  ngOnInit(): void {
    this.books._Books().subscribe({
      next: (token) => {
        this.procced.set(true)
        var eb = token.data.updatedBooks.entertainment
        var etemp: OCSFieldFormat[] = []
        for (let i of eb) {
          etemp.push({ OCStemplate: i, OCSformControl: false, OCSToolTip: "" })
        }
        var sb = token.data.updatedBooks.sports
        var stemp: OCSFieldFormat[] = []
        for (let i of sb) {
          stemp.push({ OCStemplate: i, OCSformControl: false, OCSToolTip: "" })
        }
        this.EntertainmentBooks.set(etemp)
        this.SportsBooks.set(stemp)
      }
    })

  }
  procced: WritableSignal<boolean> = signal(false)
}
