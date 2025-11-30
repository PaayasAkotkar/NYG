import { query } from '@angular/animations';
import { isPlatformBrowser } from '@angular/common';
import { Inject, Injectable, NgZone, OnInit, PLATFORM_ID } from '@angular/core';
import { ApolloClient, DocumentNode, FetchResult, Observable, throwServerError } from '@apollo/client/core';
import { Apollo, ApolloBase, gql, QueryRef } from 'apollo-angular';

import { BehaviorSubject, Subscription } from 'rxjs';
import { IStore } from '../store/store.component';
import { UpgradePattern } from '../UI/upgrade-menu/upgrade-menu.component';
import { IDailyEvents } from '../UI/daily-events/daily-events.component';
import { IDeck } from '../patches/patch';

export interface PutMessage {
  id: string
  msg: string
  roomname: string
  name: string
}

@Injectable({
  providedIn: 'root'
})
export class NYGqService implements OnInit {
  private recieveToken: BehaviorSubject<any> = new BehaviorSubject<any>("")
  private query: BehaviorSubject<any> = new BehaviorSubject<any>("")
  constructor(@Inject(PLATFORM_ID) private pid: Object, private z: NgZone, private _init: Apollo, private client: Apollo) { }
  public _refClient: Apollo
  publishMessage(message: PutMessage) {
    const q = gql`
mutation Put($input:PUT!){
  POST(client:$input){
    status
    name 
    msg
  }
}
    `
    const m = this._init.use('init_room').mutate<{
      POST: {
        status: number
        name: string
        msg: string
      }
    }>({
      mutation: q, variables: {
        input: message
      }
    })

    return m
  }
  ngOnInit(): void { }

  sub_LatestMsg(roomname: string) {
    const q = gql`
    subscription LatestMessage($room:String!){
      latest(room:$room){
           id
    msg
    roomname
name
      }
    }
    `
    const lm = this._init.use('chat_room').subscribe<{
      latest: {
        id: string
        msg: string
        roomname: string
        name: string
      }
    }>({ query: q, variables: { room: roomname } })
    return lm
  }

  sub_UpdatedDeck(id: string) {
    const q = gql`
 subscription CurrentDeck($id:ID!)         {
currentDeck(id:$id){
   id
    _first
    _second
    isDefault}
}`

    const t = this._init.subscribe<{
      currentDeck: IDeck
    }>({ query: q, variables: { id: id } })
    return t
  }
  _Books() {
    const q = gql`
    query{
     updatedBooks{
 entertainment
sports
      }
    }
`
    const b = this._init.use('login').watchQuery<
      {
        updatedBooks: {
          entertainment: string[],
          sports: string[]
        }
      }>({ query: q }).valueChanges
    return b
  }
  _Login(ID: string) {
    const l = gql`
        query Login($id: ID!){
      login(id: $id){
        msg
        updatedPlayerCredits(id: $id){
          coins
          spurs
          theme
        }
        updatedEvents(id: $id){
        locify{
            eventName
            reward
            completed
            toComplete
            progress
            isClash
            description
          }
        clash{
            eventName
            reward
            completed
            toComplete
            progress
            isClash
            description
          }
        }
        updatedStore(id: $id){
              spur{
            Title
            Price
          }
              coin{
            Title
            Price
          }
        }
        
        updatedDisplay(id: $id){
          image{
            imgURL
            imgName
          }
          clashProfile{
            rating,
              points,
              gamesPlayed,
              tier,
          }

          onevoneProfile{
            rating,
              points,
              gamesPlayed,
              tier,
          }

          twovtwoProfile{
            rating,
              points,
              gamesPlayed,
              tier,
          }

          name
          nickname
        }

        updatedPowerupUpgrades(id: $id){
          upgrades{
            power
            spur
            counter
            level
            allow
            display
          }
        }
        
currentDeck(id:$id){
   id
    _first
    _second
    isDefault
}

      }
    }
    `

    const wl = this._init.use('login').watchQuery<
      {
        login: {
          msg: string,
          updatedPlayerCredits: {
            coins: string,
            spurs: string,
          },
          updatedStore: {
            spur: IStore[],
            coin: IStore[],
          },
          updatedDisplay: {
            image: { imgURL: string, imgName: string }, name: string, nickname: string,
            onevoneProfile: { name: string, nickname: string, rating: number, points: number, gamesPlayed: number, tier: string },
            twovtwoProfile: { name: string, nickname: string, rating: number, points: number, gamesPlayed: number, tier: string },
            clashProfile: { name: string, nickname: string, rating: number, points: number, gamesPlayed: number, tier: string },
          },
          updatedPowerupUpgrades: {
            upgrades: UpgradePattern[]
          },
          updatedEvents: {
            clash: IDailyEvents[],
            locify: IDailyEvents[]
          },
          currentDeck: IDeck
        }

      }>
      ({ query: l, variables: { id: ID } }).valueChanges

    return wl
  }
  sub_UpdatedPowerupUpdgrades(ID: string) {
    const upuq = gql`
        subscription UpdatedPowerupUpgrades($id: ID!){
      updatedPowerupUpgrades(id: $id){
        upgrades{
          power
          spur
          counter
          level
          allow
          display
        }
      }
    }
    `
    const wupuq = this._init.use('setup_settings').subscribe<
      {
        updatedPowerupUpgrades: {
          upgrades: {
            power: string,
            spur: number,
            counter: number,
            level: number,
            allow: boolean,
            display: string
          }[]
        }
      }>
      ({ query: upuq, variables: { id: ID } })

    return wupuq
  }

  sub_UpdatedPlayerCredits(ID: string) {
    const upcq = gql`
        subscription UpdatedPlayerCredits($id: ID!){
      updatedPlayerCredits(id: $id){
        coins
        spurs
      }
    }
    `
    const wupcq = this._init.use('setup_settings').subscribe<
      {
        updatedPlayerCredits: {
          coins: string,
          spurs: string,
        }
      }
    >({ query: upcq, variables: { id: ID } })
    return wupcq
  }

  sub_UpdatedStore(ID: string) {
    const usq = gql`

      subscription UpdatedStore($id: ID!){
      updatedStore(id: $id){
              spur{
          Title
          Price
        }
              coin{
          Title
          Price
        }

      }
    }
    `
    const wusq = this._init.use('setup_settings').subscribe<{
      updatedStore: {

        spur: IStore[],
        coin: IStore[],
      }
    }>({ query: usq, variables: { id: ID } })

    return wusq
  }

  ClientSubscribe(): Apollo {
    this._refClient = this._init
    return this._refClient
  }
  sub_UpdatedEvents(id: string) {
    const q = gql`
 subscription UpdatedEvents($id: ID!){
      updatedEvents(id: $id){

  locify{
          eventName
          reward
          completed
          toComplete
          progress
          isClash
          description
        }

        clash{
          eventName
          reward
          completed
          toComplete
          progress
          isClash
          description
        }

      }
    }
    `
    const wue = this._init.use('setup_settings').subscribe<{

      updatedEvents: {
        clash: IDailyEvents[],
        locify: IDailyEvents[]
      },

    }>({ query: q, variables: { id: id } })

    return wue
  }
  Client(): Apollo {
    this._refClient = this.client
    return this._refClient
  }
  sub_UpdatedDisplay(ID: string) {
    const q = gql`

subscription UpdatedDisplay($id: ID!){
      updatedDisplay(id: $id){
          # image
          clashProfile{
          rating,
            points,
            gamesPlayed,
            tier,
          }

          onevoneProfile{
          rating,
            points,
            gamesPlayed,
            tier,
          }

          twovtwoProfile{
          rating,
            points,
            gamesPlayed,
            tier,
          }

        name
        nickname
      }
    }
    `
    const wq = this._init.use('setup_settings').subscribe<
      {
        updatedDisplay: {
          image: string, name: string, nickname: string,
          onevoneProfile: { name: string, nickname: string, rating: number, points: number, gamesPlayed: number, tier: string },
          twovtwoProfile: { name: string, nickname: string, rating: number, points: number, gamesPlayed: number, tier: string },
          clashProfile: { name: string, nickname: string, rating: number, points: number, gamesPlayed: number, tier: string },
        },
      }>
      ({ query: q, variables: { id: ID } })
    return wq
  }

  Init<T>(query: DocumentNode, _use: string, var_: any = {}) {
    if (isPlatformBrowser(this.pid)) {
      return this._init.use(_use).watchQuery<T>(
        {
          query: query,
          variables: var_,
          notifyOnNetworkStatusChange: true
        }
      )
    } else {
      return null
    }
  }

  Watch(query: DocumentNode,) {
    if (isPlatformBrowser(this.pid)) {
      this.client.watchQuery({
        query: query,
      }).valueChanges.subscribe({
        next: (token) => {
          this.recieveToken.next(token.data)
        }
      })
    }
  }

  // DescQuery describes the query just like mysql desc table mytable;
  DescQuery(_typename: string, _client: string) {
    const query = gql`
query GetTypeFields($name: String!)     {
      __type(name: $name) {
        name
        kind
      fields {
          name
          type {
            name
            kind
          ofType{
              kind
              name
            }
          }
        }
      }
    }
    `
    if (isPlatformBrowser(this.pid)) {
      this.client.use(_client).watchQuery({
        query: query,
        variables: { name: _typename },
        notifyOnNetworkStatusChange: true,
      }).valueChanges.subscribe({
        next: (token) => {
          this.query.next(token.data)
        }
      })
    }
    return this.query.asObservable()
  }

  Fetch() {
    return this.recieveToken.asObservable()
  }

}
