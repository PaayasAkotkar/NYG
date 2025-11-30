import { Injectable } from '@angular/core';
import { PartyChatMethods } from './resources/resources';
interface PartyForm {
  id: string,
  token: string,
}
@Injectable({
  providedIn: 'root'
})
export class PartyChatService extends PartyChatMethods {

  constructor() {
    super()
  }

}
