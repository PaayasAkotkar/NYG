import mqtt, { IClientOptions } from "mqtt";
import { Coin, MqttFieds, Spur, Upgrade } from "../vouchers/mqttVouchers";
import { HttpClient } from "@angular/common/http";
import { SessionService } from "../../storage/session/session.service";

export class MqttMethods extends MqttFieds {
    constructor(private http: HttpClient, private se: SessionService) { super() }
    connectSpur(): void {
        const conn = mqtt.connect(this.spurURI).connect()
        conn.on('message', (topic: string, token) => {
            console.log("topic: ", topic, "token: ", token)
        })
    }
    connectCoin(): void {
        const conn = mqtt.connect(this.coinURI).connect()
        conn.on('message', (topic: string, token) => {
            console.log("topic: ", topic, "token: ", token)
        })
    }

    connectUpgrade(): void {
        var i = this.se.getItem('session')
        var clientId
        if (i) {
            clientId = i
        }
        console.log("connection upgrade")
        const conn = mqtt.connect(this.upgradeURI, {
            clientId,
            clean: true,
            protocolVersion: 4,
            connectTimeout: 4000,
            reconnectPeriod: 1000,
        })
        conn.on('connect', () => {
            console.log('connected upgrade 🔥')
            conn.subscribe("upgrade")
        }).subscribe({})

        conn.on('message', (topic: string, token) => {
            console.log("topic: ", topic, "token: ", token)

        })
    }
    UpdateSpur(struct: Spur) {
        var s = this.se.getItem('session')
        if (s) {
            this.spurPurch += s
            this.http.put(this.spurPurch, struct).subscribe()
        }
    }

    UpdateCoin(struct: Coin) {
        var s = this.se.getItem('session')
        if (s) {
            this.coinPurch += s
            this.http.put(this.coinPurch, struct).subscribe()
        }
    }
    
    UpgradePower(struct: Upgrade) {
        var s = this.se.getItem('session')
        if (s) {
            this.makeUpgrade += s
            this.http.patch(this.makeUpgrade, struct).subscribe()
        }
    }

}