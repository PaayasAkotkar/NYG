import { BehaviorSubject } from "rxjs"
export interface Coin {
    coin: string
    id: string
    reciept: PurchaseReciept
}
export interface Spur {
    spur: string
    id: string
    reciept: PurchaseReciept
}
export interface Upgrade {
    spur: string // donated spur
    coin: string // coins spend
    power: string // must be the key of the power
    id: string
    newLevel: number
}
export interface PurchaseReciept {
    purchaseMade: string
    time: string
    date: string
}


export class MqttFieds {

    protected coinPurch: string = 'http://localhost:7460/coin/purchase/'
    protected spurPurch: string = 'http://localhost:7460/spur/purchase/'
    protected makeUpgrade: string = 'http://localhost:7460/upgrade/power/'

    protected spurURI: string = 'ws://localhost:1884'
    protected coinURI: string = 'ws://localhost:1883'
    protected upgradeURI: string = 'ws://localhost:1885'

    protected recieveCoinToken: BehaviorSubject<Coin> = new BehaviorSubject<Coin>({ coin: "", id: "", reciept: { time: "", purchaseMade: "", date: "" }, })
    protected recieveSpurToken: BehaviorSubject<Spur> = new BehaviorSubject<Spur>({ spur: "", id: "", reciept: { time: "", purchaseMade: "", date: "" }, })
    protected recieveUpgradeToken: BehaviorSubject<Upgrade> = new BehaviorSubject<Upgrade>({
        spur: "", coin: "", power: "", id: "", newLevel: -1
    })

}